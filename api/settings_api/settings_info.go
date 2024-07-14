package settingsApi

import "github.com/gin-gonic/gin"

func (Api) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "xxx",
	})
}
