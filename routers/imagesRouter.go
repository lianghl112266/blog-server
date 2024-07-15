package routers

import (
	"blogServer/api"

	"github.com/gin-gonic/gin"
)

func imageRouter() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		Api := api.ApiGroupApp.ImageApi

		rgPublicImages := rgPublic.Group("images")
		{
			rgPublicImages.POST("", Api.ImageUpload)
			rgPublicImages.GET("", Api.ImageListView)
		}
	})
}
