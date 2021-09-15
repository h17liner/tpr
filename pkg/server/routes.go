package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h17liner/tpr/pkg/server/providers"
	"github.com/h17liner/tpr/pkg/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
)

const (
	providersPrefix = "/v1/providers/"
	modulesPrefix   = "/v1/modules/"
)

func (app *FiberApp) routes() *FiberApp {

	app.Get("/.well-known/terraform.json", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"modules.v1":   modulesPrefix,
			"providers.v1": providersPrefix,
		})
	})

	apiProvider := app.Group(providersPrefix)

	var providerStorage *providers.ProviderStorage
	switch viper.GetString("storage.type") {
	case "fs":
		u, err := url.Parse(viper.GetString("storage.url"))
		if err != nil {
			log.Fatal("Failed with url in storage ", err)
		}

		providerStorage = providers.NewProviderStorage(
			storage.NewFs(viper.GetString("storage.path"), u),
		)

		app.Static("/", viper.GetString("storage.path"))
	case "s3":
		log.Fatal("Not implemented")
	case "artifactory":
		log.Fatal("Not implemented")
	default:
		log.Fatal("unknown storage.type.")
	}

	// TF RFC
	//https://www.terraform.io/docs/internals/provider-registry-protocol.html#list-available-versions
	apiProvider.Get("/:namespace/:type/versions", providerStorage.GetVersions)
	//https://www.terraform.io/docs/internals/provider-registry-protocol.html#find-a-provider-package
	apiProvider.Get("/:namespace/:type/:version/download/:os/:arch", providerStorage.FindProvider)

	// Managing providers
	// upload provider.
	apiProvider.Post("/:namespace/:type/:version/upload/:os/:arch", notImplemented)
	// delete provider.
	apiProvider.Delete("/:namespace/:type/:version/:os/:arch", notImplemented)

	return app
}

func notImplemented(c *fiber.Ctx) error {
	return c.JSON("not implemented")
}
