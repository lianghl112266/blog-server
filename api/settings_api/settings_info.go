package settingsApi

import (
	utils "blogServer/utils/resp"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoView(c *gin.Context) {
	utils.FailWithCode(utils.SETTINGSERROR, c)
}
