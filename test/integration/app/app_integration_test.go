package app_test

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	. "github.com/channel-io/ch-app-store/test"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	appID     = "1"
	channelID = "1"
)

type AppIntegrationTestSuite struct {
	testApp *TestApp

	appRepository             svc.AppRepository
	appInstallationRepository svc.AppInstallationRepository
}

var appIntegrationTestSuite AppIntegrationTestSuite

var _ = BeforeSuite(func() {
	appIntegrationTestSuite.testApp = NewTestApp(
		Populate(&appIntegrationTestSuite.appRepository),
		Populate(&appIntegrationTestSuite.appInstallationRepository),
	)
})

var _ = AfterSuite(func() {
	appIntegrationTestSuite.testApp.Stop()
	appIntegrationTestSuite.testApp.TruncateAll()
})

var _ = Describe("AppRepository Save", func() {
	Context("when app is saved", func() {
		It("should save app", func() {
			app := app.App{
				ID:        appID,
				IsPrivate: false,
			}

			ctx := context.Background()

			res, err := appIntegrationTestSuite.appRepository.Save(ctx, &app)

			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))

			res, err = appIntegrationTestSuite.appRepository.FindApp(ctx, appID)

			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})

var _ = Describe("AppRepository FindApp", func() {
	Context("when app is found", func() {
		It("should find app", func() {
			app := app.App{
				ID:        appID,
				IsPrivate: false,
			}

			ctx := context.Background()

			_, _ = appIntegrationTestSuite.appRepository.Save(ctx, &app)
			res, err := appIntegrationTestSuite.appRepository.FindApp(ctx, appID)

			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})

var _ = Describe("AppInstallation Save", func() {
	Context("when app channel is saved", func() {
		It("should save app channel", func() {
			appChannel := &app.AppInstallation{
				ChannelID: channelID,
				AppID:     appID,
			}

			app := app.App{
				ID:        appID,
				IsPrivate: false,
			}

			ctx := context.Background()

			_, _ = appIntegrationTestSuite.appRepository.Save(ctx, &app)

			err := appIntegrationTestSuite.appInstallationRepository.Save(context.Background(), appChannel)

			Expect(err).To(BeNil())
		})
	})
})
