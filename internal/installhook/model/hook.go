package model

type AppInstallHooks struct {
	AppID                 string  `json:"appId"`
	InstallFunctionName   *string `json:"installFunctionName,omitempty"`
	UninstallFunctionName *string `json:"uninstallFunctionName,omitempty"`
}
