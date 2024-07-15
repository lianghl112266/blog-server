package api

import (
	imageApi "blogServer/api/image_api"
	settingsApi "blogServer/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settingsApi.Api
	ImageApi    imageApi.Api
}

var ApiGroupApp = new(ApiGroup)
