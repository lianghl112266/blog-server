package imageApi

import (
	"blogServer/global"
	"blogServer/models"
	"blogServer/service/common"
	"blogServer/utils"

	"github.com/gin-gonic/gin"
)

func (Api) ImageListView(c *gin.Context) {

	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		global.Log.Errorf("request argument error")
		return
	}

	list, cnt, err := common.ComList(models.BannerModel{}, common.Option{PageInfo: page, Debug: false})
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		global.Log.Errorf(err.Error())
		return
	}
	utils.OkWithData(gin.H{"count": cnt, "list": list}, c)
	global.Log.Infof("successful  get image list")
}
