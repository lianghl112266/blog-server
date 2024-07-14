package api

import settingsApi "blogServer/api/settings_api"

type ApiGroup struct {
	SettingsApi settingsApi.Api
}

var ApiGroupApp = new(ApiGroup)
