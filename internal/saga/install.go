package saga

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

type InstallSaga struct {
	appChInstallSvc *appChannel.InstallSvc
	appChCfgSvc     *appChannel.ConfigSvc

	appCfgSvc     *app.ConfigSvc
	appInstallSvc *app.InstallSvc
}

func NewInstallSaga(
	appChInstallSvc *appChannel.InstallSvc,
	appChCfgSvc *appChannel.ConfigSvc,
	appCfgSvc *app.ConfigSvc,
	appInstallSvc *app.InstallSvc,
) *InstallSaga {
	return &InstallSaga{
		appChInstallSvc: appChInstallSvc,
		appChCfgSvc:     appChCfgSvc,
		appCfgSvc:       appCfgSvc,
		appInstallSvc:   appInstallSvc,
	}
}

func (s *InstallSaga) Uninstall(ctx context.Context, identifier appChannel.AppChannelIdentifier) error {
	if err := s.appChInstallSvc.Uninstall(ctx, identifier); err != nil {
		return err
	}

	if err := s.appInstallSvc.NotifyUnInstall(
		ctx,
		app.InstallInfo{
			AppID:     identifier.AppID,
			ChannelID: identifier.ChannelID,
		},
	); err != nil {
		return err
	}

	return nil
}

func (s *InstallSaga) Install(ctx context.Context, identifier appChannel.AppChannelIdentifier) (*appChannel.AppChannel, error) {
	install := app.InstallInfo{AppID: identifier.AppID, ChannelID: identifier.ChannelID}

	res, err := s.appInstallSvc.CheckInstallable(ctx, install)
	if err != nil {
		return nil, err
	} else if !res.Result {
		return nil, err
	}

	defaultCfg, err := s.appCfgSvc.DefaultConfigOf(ctx, install)
	if err != nil {
		return nil, err
	}

	created, err := s.appChInstallSvc.Install(ctx, identifier, appChannel.ConfigMap(defaultCfg))
	if err != nil {
		return nil, err
	}

	if err := s.appInstallSvc.NotifyUnInstall(ctx, install); err != nil {
		return nil, err
	}

	return created, nil
}

func (s *InstallSaga) SetConfig(
	ctx context.Context,
	identifier appChannel.AppChannelIdentifier,
	newConfig appChannel.ConfigMap,
) (app.CheckReturn, error) {
	install := app.InstallInfo{AppID: identifier.AppID, ChannelID: identifier.ChannelID}

	ret, err := s.appCfgSvc.CheckConfig(ctx, install, app.ConfigMap(newConfig))
	if err != nil {
		return ret, err
	}
	if !ret.Success {
		return ret, nil
	}

	if err := s.appChCfgSvc.SetConfig(ctx, identifier, newConfig); err != nil {
		return app.CheckReturn{}, err
	}

	return ret, nil
}
