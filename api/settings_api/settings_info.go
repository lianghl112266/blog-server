package settingsApi

import (
	"blogServer/global"
	utils "blogServer/utils/resp"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoView(c *gin.Context) {
	type SettingsInfo struct {
		Name string `uri:"name"`
	}

	settingsInfo := SettingsInfo{}
	err := c.ShouldBindUri(&settingsInfo)
	if err != nil {
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		global.Log.Errorf("request argument error")
		return
	}
	switch settingsInfo.Name {
	case "qq":
		utils.OkWithData(global.Config.QQ, c)
	case "qiniu":
		utils.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		utils.OkWithData(global.Config.Jwt, c)
	case "site":
		utils.OkWithData(global.Config.SiteInfo, c)
	case "email":
		utils.OkWithData(global.Config.Email, c)
	default:
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		global.Log.Errorf("request argument error")
	}
}
