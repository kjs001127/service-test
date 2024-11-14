package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type InstallHandler struct {
	agreementRepo ChannelRoleAgreementRepo
}

func NewInstallHandler(agreementRepo ChannelRoleAgreementRepo) *InstallHandler {
	return &InstallHandler{agreementRepo: agreementRepo}
}

func (i InstallHandler) OnInstall(ctx context.Context, app *model.App, channelID string) error {
	return nil
}

func (i InstallHandler) OnUnInstall(ctx context.Context, app *model.App, channelID string) error {
	return i.agreementRepo.DeleteAllByInstallID(ctx, model.InstallationID{
		AppID:     app.ID,
		ChannelID: channelID,
	})
}
