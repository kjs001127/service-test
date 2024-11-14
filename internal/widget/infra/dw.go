package infra

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/widget/model"
)

type DropwizardEventPublisher struct {
	baseURL string
	resty   *resty.Client
}

func NewDropwizardEventPublisher(baseURL string, resty *resty.Client) *DropwizardEventPublisher {
	return &DropwizardEventPublisher{baseURL: baseURL, resty: resty}
}

func (d *DropwizardEventPublisher) OnDeleted(ctx context.Context, appWidgets []*model.AppWidget) error {
	req := d.resty.R()
	req.SetQueryParamsFromValues(url.Values{
		"appWidgetIds": idsOf(appWidgets),
	})
	req.SetContext(ctx)
	res, err := req.Delete(d.baseURL + "/api/admin/app-widget-installations")
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		return fmt.Errorf("appWidget delete request to dw of widgetIds: %v failed with status %s", idsOf(appWidgets), res.Status())
	}

	return nil
}

func idsOf(widgets []*model.AppWidget) []string {
	ret := make([]string, 0, len(widgets))
	for _, w := range widgets {
		ret = append(ret, w.ID)
	}
	return ret
}

func (d *DropwizardEventPublisher) OnUnInstall(ctx context.Context, install appmodel.InstallationID) error {
	req := d.resty.R()
	req.SetQueryParams(map[string]string{
		"appId": install.AppID,
	})
	req.SetContext(ctx)
	res, err := req.Delete(
		fmt.Sprintf(d.baseURL+"/api/admin/channels/%s/app-widget-installations", install.ChannelID),
	)
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		return fmt.Errorf("channel appWidget delete request to dw of installationID: %v failed with status %s", install, res.Status())
	}

	return nil
}
