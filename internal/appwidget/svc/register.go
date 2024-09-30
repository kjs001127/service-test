package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/appwidget/model"
	"github.com/channel-io/ch-app-store/internal/util"
	"github.com/channel-io/ch-app-store/lib/db/tx"
	"github.com/channel-io/ch-app-store/lib/log"
)

type RegisterSvc interface {
	Register(ctx context.Context, req *AppWidgetRegisterRequest) error
}

type RegisterSvcImpl struct {
	repo      AppWidgetRepository
	publisher EventPublisher
	logger    log.ContextAwareLogger
}

func NewRegisterSvc(repo AppWidgetRepository, publisher EventPublisher, logger log.ContextAwareLogger) *RegisterSvcImpl {
	return &RegisterSvcImpl{repo: repo, publisher: publisher, logger: logger}
}

func (s *RegisterSvcImpl) DeregisterAll(ctx context.Context, appID string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		if err := s.repo.DeleteAllByAppID(ctx, appID); err != nil {
			return err
		}
		return nil
	}, tx.XLock(namespaceAppWidget, appID))
}

func (s *RegisterSvcImpl) Register(ctx context.Context, req *AppWidgetRegisterRequest) error {
	res, err := s.registerWithTx(ctx, req)
	if err != nil {
		return err
	}

	if len(res.deleted) > 0 {
		go s.publishDeletedEvent(res.deleted)
	}

	return nil
}

func (s *RegisterSvcImpl) registerWithTx(ctx context.Context, req *AppWidgetRegisterRequest) (*UpdateResult, error) {
	return tx.DoReturn(ctx, func(ctx context.Context) (*UpdateResult, error) {
		if err := s.validateRequest(req.AppID, req.AppWidgets); err != nil {
			return nil, errors.WithStack(err)
		}

		oldbies, err := s.repo.FetchAllByAppIDs(ctx, []string{req.AppID})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		res := newUpdateResult(s.repo)
		updater := util.DeltaUpdater[*model.AppWidget, UpdateKey]{
			IDOf:     res.updateKey,
			DoInsert: res.insertResource,
			DoUpdate: res.updateResource,
			DoDelete: res.deleteResource,
		}

		return res, updater.Update(ctx, oldbies, req.AppWidgets)
	}, tx.XLock(namespaceAppWidget, req.AppID))
}

func (s *RegisterSvcImpl) validateRequest(appID string, appWidgets []*model.AppWidget) error {
	if len(appWidgets) > 10 {
		return apierr.BadRequest(fmt.Errorf("you can only register up to 10 appWidgets"))
	}

	for _, appWidget := range appWidgets {
		if len(appWidget.AppID) <= 0 {
			appWidget.AppID = appID
		} else if appWidget.AppID != appID {
			return apierr.BadRequest(fmt.Errorf("request AppID: %s doesn't match AppID of cmd: %s", appID, appWidget.AppID))
		}

		if err := appWidget.Validate(); err != nil {
			return apierr.BadRequest(err)
		}
	}

	return nil
}

func (s *RegisterSvcImpl) publishDeletedEvent(deleted []*model.AppWidget) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.publisher.OnDeleted(ctx, deleted); err != nil {
		s.logger.Errorw(ctx, "widget deleted event publish fail", "err", err, "widgets", deleted)
	}
}

type AppWidgetRegisterRequest struct {
	AppID      string             `json:"appId"`
	AppWidgets []*model.AppWidget `json:"appWidgets"`
}

type UpdateKey struct {
	Scope model.Scope
	Name  string
}

type UpdateResult struct {
	deleted  []*model.AppWidget
	updated  []*model.AppWidget
	inserted []*model.AppWidget

	repo AppWidgetRepository
}

func newUpdateResult(repo AppWidgetRepository) *UpdateResult {
	return &UpdateResult{
		updated:  make([]*model.AppWidget, 0),
		deleted:  make([]*model.AppWidget, 0),
		inserted: make([]*model.AppWidget, 0),
		repo:     repo,
	}
}
func (s *UpdateResult) updateKey(resource *model.AppWidget) UpdateKey {
	return UpdateKey{Scope: resource.Scope, Name: resource.Name}
}

func (s *UpdateResult) insertResource(ctx context.Context, newbie *model.AppWidget) error {
	newbie.ID = uid.New().Hex()
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save widget fail. widget: %v, cause: %w", newbie, err)
	}
	s.inserted = append(s.inserted, newbie)
	return nil
}

func (s *UpdateResult) updateResource(ctx context.Context, oldbie *model.AppWidget, newbie *model.AppWidget) error {
	newbie.ID = oldbie.ID
	if _, err := s.repo.Save(ctx, newbie); err != nil {
		return fmt.Errorf("save widget fail. widget: %v, cause: %w", newbie, err)
	}
	s.updated = append(s.updated, newbie)
	return nil
}

func (s *UpdateResult) deleteResource(ctx context.Context, oldbie *model.AppWidget) error {
	if err := s.repo.Delete(ctx, oldbie.ID); err != nil {
		return fmt.Errorf("delete widget fail. widget: %v, cause: %w", oldbie, err)
	}
	s.deleted = append(s.deleted, oldbie)
	return nil
}
