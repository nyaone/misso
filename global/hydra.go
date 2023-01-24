package global

import client "github.com/ory/hydra-client-go/v2"

type hydraClients struct {
	Admin *client.APIClient
}

var Hydra hydraClients
