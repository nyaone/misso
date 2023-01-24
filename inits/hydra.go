package inits

import (
	client "github.com/ory/hydra-client-go/v2"
	"misso/config"
	"misso/global"
)

func Hydra() error {
	adminApiConfig := client.NewConfiguration()
	adminApiConfig.Servers = []client.ServerConfiguration{
		{
			URL: config.Config.Hydra.AdminUrl,
		},
	}
	global.Hydra.Admin = client.NewAPIClient(adminApiConfig)

	return nil
}
