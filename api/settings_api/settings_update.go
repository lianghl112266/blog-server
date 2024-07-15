package settingsApi

import (
	"blogServer/config"
	"blogServer/core"
	"blogServer/global"
	utils "blogServer/utils/resp"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoUpdate(c *gin.Context) {
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
		var qq config.QQ
		err = c.ShouldBindBodyWithJSON(&qq)
		if err != nil {
			global.Log.Errorf(err.Error())
			utils.FailWithCode(utils.ARGUMENTERROR, c)
			return
		}
		global.Config.QQ = qq
		err = core.SetYaml()
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		global.Log.Infof("yaml set successful")
		utils.OkWithMessage("successful", c)
	case "qiniu":
		var qi_niu config.QiNiu
		err = c.ShouldBindBodyWithJSON(&qi_niu)
		if err != nil {
			global.Log.Errorf(err.Error())
			utils.FailWithCode(utils.ARGUMENTERROR, c)
			return
		}
		global.Config.QiNiu = qi_niu
		err = core.SetYaml()
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		global.Log.Infof("yaml set successful")
		utils.OkWithMessage("successful", c)
	case "jwt":
		var jwt config.Jwt
		err = c.ShouldBindBodyWithJSON(&jwt)
		if err != nil {
			global.Log.Errorf(err.Error())
			utils.FailWithCode(utils.ARGUMENTERROR, c)
			return
		}
		global.Config.Jwt = jwt
		err = core.SetYaml()
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		global.Log.Infof("yaml set successful")
		utils.OkWithMessage("successful", c)
	case "site":
		var siteInfo config.SiteInfo
		err = c.ShouldBindBodyWithJSON(&siteInfo)
		if err != nil {
			global.Log.Errorf(err.Error())
			utils.FailWithCode(utils.ARGUMENTERROR, c)
			return
		}
		global.Config.SiteInfo = siteInfo
		err = core.SetYaml()
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		global.Log.Infof("yaml set successful")
		utils.OkWithMessage("successful", c)
	default:
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		global.Log.Errorf("request argument error")
	}
}
