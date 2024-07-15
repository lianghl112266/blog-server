package settingsApi

import (
	"blogServer/config"
	"blogServer/core"
	"blogServer/global"
	utils "blogServer/utils/resp"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoUpdate(c *gin.Context) {
	var siteInfo config.SiteInfo
	err := c.ShouldBindBodyWithJSON(&siteInfo)
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
}
