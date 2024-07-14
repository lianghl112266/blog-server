package models

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash"`                // 图片的hash值，用于判断重复图片
	Name string `gorm:"size:38" json:"name"` // 图片名称
}

// func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
// 	if b.ImageType == ctype.Local {
// 		// 本地图片，删除，还要删除本地的存储
// 		err = os.Remove(b.Path)
// 		if err != nil {
// 			global.Log.Error(err)
// 			return err
// 		}
// 	}
// 	return nil
// }
