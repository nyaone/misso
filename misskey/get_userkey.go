package misskey

import (
	"misso/config"
	"misso/types"
)

type AuthSessionUserkey_Request struct {
	AppSecret string `json:"appSecret"`
	Token     string `json:"token"`
}

type AuthSessionUserkey_Response struct {
	AccessToken string                `json:"accessToken"`
	User        types.MisskeyUserBase `json:"user"`
}

func GetUserkey(token string) (*AuthSessionUserkey_Response, error) {

	return PostAPIRequest[AuthSessionUserkey_Response]("auth/session/userkey", &AuthSessionUserkey_Request{
		AppSecret: config.Config.Misskey.Application.Secret,
		Token:     token,
	})

}
