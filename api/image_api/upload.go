package imageApi

import (
	"blogServer/global"
	"blogServer/models"
	"blogServer/utils"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// 定义白名单集合
var whiteSet = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".gif":  {},
	".bmp":  {},
	".tiff": {},
	".psd":  {},
	".raw":  {},
	".ico":  {},
	".webp": {},
	".svg":  {},
}

func (Api) ImageUpload(c *gin.Context) {
	from, err := c.MultipartForm()
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		global.Log.Errorf(err.Error())
		return
	}

	files, ok := from.File["images"]
	if !ok {
		utils.FailWithMessage("file does not exist", c)
		global.Log.Errorf("file does not exist")
		return
	}

	type Resp struct {
		Entry   string
		Success bool
		Msg     string
	}

	ans := []Resp{}

	basePath := global.Config.Upload.Path
	limitSize := global.Config.Upload.Size

	if _, err := os.Stat(basePath); err != nil {
		os.MkdirAll(basePath, 0755)
	}

	for _, f := range files {
		if float32(f.Size)/1024/1024 > limitSize {
			ans = append(ans, Resp{
				Entry:   f.Filename,
				Success: false,
				Msg:     "upload file transcend limit(6.3MB)",
			})
			global.Log.Warnf("upload %s failed, reason: upload file transcend limit(6.3MB)", f.Filename)
			continue
		}

		idx := strings.LastIndexByte(f.Filename, '.')
		if _, ok := whiteSet[strings.ToLower(f.Filename[idx:])]; !ok {
			ans = append(ans, Resp{
				Entry:   f.Filename,
				Success: false,
				Msg:     "upload file is not a picture",
			})
			global.Log.Warnf(" upload file %s is not a picture", f.Filename)
			continue
		}

		ff, _ := f.Open()
		byteData, _ := io.ReadAll(ff)
		hash := utils.Md5(byteData)
		var bannerModel models.BannerModel
		if err = global.DB.Take(&bannerModel, "hash = ?", hash).Error; err == nil {
			ans = append(ans, Resp{
				Entry:   bannerModel.Path,
				Success: true,
				Msg:     "upload file already exists",
			})
			global.Log.Warnf(" upload file %s already exists", f.Filename)
			continue
		}

		filePath := basePath + f.Filename
		c.SaveUploadedFile(f, filePath)
		ans = append(ans, Resp{
			Entry:   filePath,
			Success: true,
			Msg:     "",
		})
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: hash,
			Name: f.Filename,
		})
		global.Log.Infof("upload %s successful", f.Filename)
	}
	utils.OkWithData(ans, c)
}
