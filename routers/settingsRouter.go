package routers

import (
	"blogServer/api"

	"github.com/gin-gonic/gin"
)

func settingsRouter() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		Api := api.ApiGroupApp.SettingsApi
		rgPublicSettings := rgPublic.Group("settings")
		{
			rgPublicSettings.GET("", Api.SettingsInfoView)
			rgPublicSettings.POST("", Api.SettingsInfoUpdate)
		}

	})
}
