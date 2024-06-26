package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/command/model"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type ActivationSvc interface {
	Toggle(ctx context.Context, req ToggleCommandRequest) error
	ToggleByKey(ctx context.Context, req ToggleRequest) error
	Check(ctx context.Context, key model.CommandKey, channelID string) (bool, error)
}

type ActivationSvcImpl struct {
	activationRepo ActivationRepository
	cmdRepo        CommandRepository
}

func NewActivationSvc(activationRepo ActivationRepository, cmdRepo CommandRepository) *ActivationSvcImpl {
	return &ActivationSvcImpl{activationRepo: activationRepo, cmdRepo: cmdRepo}
}

type ToggleCommandRequest struct {
	Command   *model.Command `json:"command"`
	ChannelID string         `json:"channelId"`
	Enabled   bool           `json:"enabled"`
}

type ToggleRequest struct {
	Command   model.CommandKey `json:"command"`
	ChannelID string           `json:"channelId"`
	Enabled   bool             `json:"enabled"`
}

func (s *ActivationSvcImpl) Toggle(ctx context.Context, req ToggleCommandRequest) error {
	return s.activationRepo.Save(ctx, &model.Activation{
		ActivationID: model.ActivationID{
			CommandID: req.Command.ID,
			ChannelID: req.ChannelID,
		},
		Enabled: req.Enabled,
	})
}

func (s *ActivationSvcImpl) ToggleByKey(ctx context.Context, req ToggleRequest) error {
	cmd, err := s.cmdRepo.Fetch(ctx, req.Command)
	if err != nil {
		return err
	}

	return s.Toggle(ctx, ToggleCommandRequest{
		Command:   cmd,
		ChannelID: req.ChannelID,
		Enabled:   req.Enabled,
	})
}

func (s *ActivationSvcImpl) Check(ctx context.Context, key model.CommandKey, channelID string) (bool, error) {
	cmd, err := s.cmdRepo.Fetch(ctx, key)
	if err != nil {
		return false, err
	}

	activation, err := s.activationRepo.Fetch(ctx, model.ActivationID{
		CommandID: cmd.ID,
		ChannelID: channelID,
	})
	if apierr.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return activation.Enabled, nil
}

type InstalledCommandQuerySvc struct {
	activationRepo ActivationRepository
	cmdRepo        CommandRepository
}

func NewInstalledCommandQuerySvc(activationRepo ActivationRepository, cmdRepo CommandRepository) *InstalledCommandQuerySvc {
	return &InstalledCommandQuerySvc{activationRepo: activationRepo, cmdRepo: cmdRepo}
}

type CommandWithActivation struct {
	*model.Command
	Enabled bool `json:"enabled"`
}

func (s *InstalledCommandQuerySvc) FetchAllWithActivation(ctx context.Context, installID appmodel.InstallationID) ([]*CommandWithActivation, error) {
	cmds, err := s.cmdRepo.FetchAllByAppID(ctx, installID.AppID)
	if err != nil {
		return nil, err
	}
	activations, err := s.activationRepo.FetchByChannelIDAndCmdIDs(ctx, installID.ChannelID, idsOfCmds(cmds))
	activationMap := activations.ToMap()

	ret := make([]*CommandWithActivation, 0, len(cmds))
	for _, cmd := range cmds {
		if activation, exists := activationMap[cmd.ID]; exists {
			ret = append(ret, &CommandWithActivation{
				Command: cmd,
				Enabled: activation.Enabled,
			})
		} else {
			ret = append(ret, &CommandWithActivation{
				Command: cmd,
				Enabled: false,
			})
		}
	}

	return ret, nil
}

func idsOfCmds(cmds []*model.Command) []string {
	ret := make([]string, 0, len(cmds))
	for _, cmd := range cmds {
		ret = append(ret, cmd.ID)
	}
	return ret
}

type PreInstallHandler struct {
	activationRepo ActivationRepository
	cmdRepo        CommandRepository
}

func NewPreInstallHandler(activationRepo ActivationRepository, cmdRepo CommandRepository) *PreInstallHandler {
	return &PreInstallHandler{activationRepo: activationRepo, cmdRepo: cmdRepo}
}

func (h *PreInstallHandler) OnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	cmds, err := h.cmdRepo.FetchAllByAppID(ctx, app.ID)
	if err != nil {
		return err
	}

	for _, cmd := range cmds {
		err := h.activationRepo.Save(ctx, &model.Activation{
			ActivationID: model.ActivationID{
				CommandID: cmd.ID,
				ChannelID: channelID,
			},
			Enabled: cmd.EnabledByDefault,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *PreInstallHandler) OnUnInstall(ctx context.Context, app *appmodel.App, channelID string) error {
	cmds, err := h.cmdRepo.FetchAllByAppID(ctx, app.ID)
	if err != nil {
		return err
	}

	return h.activationRepo.DeleteAllBy(ctx, channelID, idsOfCmds(cmds))
}
