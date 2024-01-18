package saga

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

type InstallSaga struct {
	appChInstallSvc *appChannel.InstallSvc
	appChCfgSvc     *appChannel.ConfigSvc

	appCfgSvc    *app.ConfigSvc
	appNotifySvc *app.NotifySvc
	appQuerySvc  *app.QuerySvc
}

func NewInstallSaga(
	appChInstallSvc *appChannel.InstallSvc,
	appChCfgSvc *appChannel.ConfigSvc,
	appCfgSvc *app.ConfigSvc,
	notifySvc *app.NotifySvc,
	appQuerySvc *app.QuerySvc,
) *InstallSaga {
	return &InstallSaga{
		appChInstallSvc: appChInstallSvc,
		appChCfgSvc:     appChCfgSvc,
		appCfgSvc:       appCfgSvc,
		appNotifySvc:    notifySvc,
		appQuerySvc:     appQuerySvc,
	}
}

func (s *InstallSaga) Uninstall(ctx context.Context, identifier appChannel.AppChannelIdentifier) error {
	if err := s.appChInstallSvc.Uninstall(ctx, identifier); err != nil {
		return err
	}

	install := app.InstallInfo{AppId: identifier.AppID, ChannelId: identifier.ChannelID}
	if err := s.appNotifySvc.NotifyUnInstall(ctx, install); err != nil {
		return err
	}

	return nil
}

func (s *InstallSaga) Install(ctx context.Context, identifier appChannel.AppChannelIdentifier) (*appChannel.AppChannel, error) {
	install := app.InstallInfo{AppId: identifier.AppID, ChannelId: identifier.ChannelID}

	res, err := s.appQuerySvc.CheckInstallable(ctx, install)
	if err != nil {
		return nil, err
	} else if !res.Result {
		return nil, err
	}

	defaultCfg := s.appCfgSvc.DefaultConfigOf(ctx, install)

	created, err := s.appChInstallSvc.Install(ctx, identifier, toMap(defaultCfg))
	if err != nil {
		return nil, err
	}

	if err := s.appNotifySvc.NotifyInstall(ctx, install); err != nil {
		return nil, err
	}

	return created, nil
}

func (s *InstallSaga) SetConfig(
	ctx context.Context,
	identifier appChannel.AppChannelIdentifier,
	newConfig appChannel.ConfigMap,
) error {
	install := app.InstallInfo{AppId: identifier.AppID, ChannelId: identifier.ChannelID}

	if err := s.appCfgSvc.ValidateConfigs(ctx, install, toConfigInput(newConfig)); err != nil {
		return err
	}

	if err := s.appChCfgSvc.SetConfig(ctx, identifier, newConfig); err != nil {
		return err
	}

	return nil
}

func toConfigInput(input map[string]string) []*app.ConfigValue {
	var ret []*app.ConfigValue
	for key, val := range input {
		ret = append(ret, &app.ConfigValue{Key: key, Value: val})
	}
	return ret
}

func toMap(configs []*app.ConfigValue) map[string]string {
	ret := make(map[string]string)
	for _, config := range configs {
		ret[config.Key] = config.Value
	}
	return ret
}
