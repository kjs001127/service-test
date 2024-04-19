package model

type AppInstallHooks struct {
	AppID                 string  `json:"appId"`
	InstallFunctionName   *string `json:"installFunctionName"`
	UninstallFunctionName *string `json:"uninstallFunctionName"`
}
