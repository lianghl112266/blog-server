package settingsApi

import (
	"blogServer/global"
	utils "blogServer/utils/resp"

	"github.com/gin-gonic/gin"
)

func (Api) SettingsInfoView(c *gin.Context) {
	utils.OkWithData(global.Config.SiteInfo, c)
}
