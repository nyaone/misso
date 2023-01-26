package misskey

import "misso/types"

type I_Request struct {
	I string `json:"i"`
}

type I_Response = types.MisskeyUser

func GetUserinfo(accessToken string) (*I_Response, error) {
	return PostAPIRequest[I_Response]("/api/i", &I_Request{
		I: accessToken,
	})
}
