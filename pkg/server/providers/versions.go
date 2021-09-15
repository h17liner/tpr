package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h17liner/tpr/pkg/storage"
)

//https://www.terraform.io/docs/internals/provider-registry-protocol.html#list-available-versions

type versionsResponse struct {
	Version []*providerVersion `json:"versions"`
}

type providerVersion struct {
	Version   string              `json:"version"`
	Protocols []string            `json:"protocols"` // recommended but not necessery
	Platforms []*providerPlatform `json:"platforms"`
}

type providerPlatform struct {
	Os   string `json:"os"`
	Arch string `json:"arch"`
}

func (p *ProviderStorage) GetVersions(c *fiber.Ctx) error {
	namespace := c.Params("namespace")
	name := c.Params("type")

	providerTree, err := p.Storage.Tree(namespace, name)
	if err != nil {
		return err
	}

	version := []*providerVersion{}
	for k, v := range providerTree {
		platforms := []*providerPlatform{}
		for _, v := range v {
			if v.Type != storage.BINARY {
				continue
			}
			platforms = append(platforms, &providerPlatform{Os: v.Os, Arch: v.Arch})
		}

		version = append(version, &providerVersion{
			Version:   k,
			Platforms: platforms,
			Protocols: protocols,
		})
	}

	return c.JSON(&versionsResponse{Version: version})
}
