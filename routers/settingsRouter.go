package routers

import (
	"blogServer/api"

	"github.com/gin-gonic/gin"
)

func settingsRouter(r *gin.RouterGroup) {
	api := api.ApiGroupApp.SettingsApi
	r.GET("", api.SettingsInfoView)
}
