package misskey

import "misso/config"

type AuthSessionGenerate_Request struct {
	AppSecret string `json:"appSecret"`
}

type AuthSessionGenerate_Response struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

func GenerateAuthSession() (*AuthSessionGenerate_Response, error) {

	return PostAPIRequest[AuthSessionGenerate_Response]("auth/session/generate", &AuthSessionGenerate_Request{
		AppSecret: config.Config.Misskey.Application.Secret,
	})

}
