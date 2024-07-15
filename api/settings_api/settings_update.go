package settingsApi

import (
	"blogServer/core"
	"blogServer/global"
	"blogServer/utils"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoUpdate(c *gin.Context) {
	type SettingsInfo struct {
		Name string `uri:"name"`
	}

	mp := map[string]any{
		"qq":    &global.Config.QQ,
		"email": &global.Config.Email,
		"jwt":   &global.Config.Jwt,
		"site":  &global.Config.SiteInfo,
		"qiniu": &global.Config.QiNiu,
	}

	settingsInfo := SettingsInfo{}
	err := c.ShouldBindUri(&settingsInfo)
	if err != nil {
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		global.Log.Errorf(err.Error())
		return
	}

	if _, ok := mp[settingsInfo.Name]; !ok {
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		return
	}

	err = c.ShouldBindBodyWithJSON(mp[settingsInfo.Name])
	if err != nil {
		global.Log.Errorf(err.Error())
		utils.FailWithCode(utils.ARGUMENTERROR, c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		global.Log.Errorf(err.Error())
		return
	}
	global.Log.Infof("yaml set successful")
	utils.OkWithMessage("successful", c)

}
