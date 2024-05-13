package svc_test

import (
	"context"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	. "github.com/channel-io/ch-app-store/test/integration"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"testing"
)

type AppIntegrationTestSuite struct {
	testHelper *TestHelper

	appLifecycleSvc svc.AppLifecycleSvc
	appQuerySvc     svc.AppQuerySvc
	appRepository   svc.AppRepository
}

var suite AppIntegrationTestSuite

var _ = BeforeSuite(func() {
	suite.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&suite.appLifecycleSvc),
		fx.Populate(&suite.appQuerySvc),
		fx.Populate(&suite.appRepository),
	)
	suite.testHelper.WithPreparedTables("apps")
})

var _ = AfterEach(func() {
	suite.testHelper.CleanTables("apps")
})

var _ = AfterSuite(func() {
	suite.testHelper.Stop()
})

var _ = Describe("App create", func() {
	var app *appmodel.App
	var err error

	Context("when creating an app", func() {
		It("should create an app", func() {
			ctx := context.Background()

			app, err = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(app).NotTo(BeNil())
			Expect(app.ID).NotTo(BeEmpty())
			Expect(app.Title).To(Equal("test app"))
		})
	})
})

var _ = Describe("App read", func() {
	Context("when app exists", func() {
		var app *appmodel.App

		BeforeEach(func() {
			ctx := context.Background()

			app, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})
		})

		It("should read an app", func() {
			ctx := context.Background()

			ret, err := suite.appQuerySvc.Read(ctx, app.ID)

			Expect(err).NotTo(HaveOccurred())
			Expect(ret).NotTo(BeNil())
			Expect(ret.ID).To(Equal(app.ID))
			Expect(ret.Title).To(Equal("test app"))
		})
	})

	Context("when app not exists", func() {
		It("should return an error", func() {
			ctx := context.Background()

			ret, err := suite.appQuerySvc.Read(ctx, "not-exist")

			Expect(err).To(HaveOccurred())
			Expect(ret).To(BeNil())
		})
	})
})

var _ = Describe("App update", func() {
	Context("when app exists", func() {
		var app *appmodel.App

		BeforeEach(func() {
			ctx := context.Background()

			app, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})
		})

		It("should update an app", func() {
			ctx := context.Background()

			ret, err := suite.appLifecycleSvc.Update(ctx, &appmodel.App{
				ID:    app.ID,
				Title: "updated test app",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(ret).NotTo(BeNil())
			Expect(ret.ID).To(Equal(app.ID))
			Expect(ret.Title).To(Equal("updated test app"))
		})
	})

	Context("when app not exists", func() {
		It("should return an error", func() {
			ctx := context.Background()

			ret, err := suite.appLifecycleSvc.Update(ctx, &appmodel.App{
				ID:    "not-exist",
				Title: "updated test app",
			})

			Expect(err).To(HaveOccurred())
			Expect(ret).To(BeNil())
		})
	})
})

var _ = Describe("App delete", func() {
	Context("when app exists", func() {
		var app *appmodel.App

		BeforeEach(func() {
			ctx := context.Background()

			app, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})
		})

		It("should delete an app", func() {
			ctx := context.Background()

			err := suite.appLifecycleSvc.Delete(ctx, app.ID)

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when app not exists", func() {
		It("should return an error", func() {
			ctx := context.Background()

			err := suite.appLifecycleSvc.Delete(ctx, "not-exist")

			Expect(err).To(HaveOccurred())
		})
	})
})

var _ = Describe("Read public Apps", func() {
	Context("when public apps exist", func() {
		var app *appmodel.App

		BeforeEach(func() {
			ctx := context.Background()

			app, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})
		})

		It("should read apps", func() {
			ctx := context.Background()

			apps, err := suite.appQuerySvc.ReadPublicApps(ctx, "0", 500)

			Expect(err).NotTo(HaveOccurred())
			Expect(apps).NotTo(BeNil())
			Expect(apps).To(HaveLen(1))
			Expect(apps[0].ID).To(Equal(app.ID))
			Expect(apps[0].Title).To(Equal("test app"))
		})
	})

	Context("when public apps not exist", func() {

		BeforeEach(func() {
			ctx := context.Background()

			_, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title:     "test app",
				IsPrivate: true,
			})
		})

		It("should return empty", func() {
			ctx := context.Background()

			apps, err := suite.appQuerySvc.ReadPublicApps(ctx, "0", 500)

			Expect(err).NotTo(HaveOccurred())
			Expect(apps).To(HaveLen(0))
		})
	})
})

var _ = Describe("Read all by appIDs", func() {
	Context("when apps exist", func() {
		var app *appmodel.App

		BeforeEach(func() {
			ctx := context.Background()

			app, _ = suite.appLifecycleSvc.Create(ctx, &appmodel.App{
				Title: "test app",
			})
		})

		It("should read apps", func() {
			ctx := context.Background()

			apps, err := suite.appQuerySvc.ReadAllByAppIDs(ctx, []string{app.ID})

			Expect(err).NotTo(HaveOccurred())
			Expect(apps).NotTo(BeNil())
			Expect(apps).To(HaveLen(1))
			Expect(apps[0].ID).To(Equal(app.ID))
			Expect(apps[0].Title).To(Equal("test app"))
		})
	})

	Context("when apps not exist", func() {
		It("should return empty", func() {
			ctx := context.Background()

			apps, err := suite.appQuerySvc.ReadAllByAppIDs(ctx, []string{"not-exist"})

			Expect(err).NotTo(HaveOccurred())
			Expect(apps).To(HaveLen(0))
		})
	})
})

func TestAppIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Integration Test Suite")
}
