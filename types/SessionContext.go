package types

type SessionContext struct {
	MisskeyToken string      `json:"token"`
	User         MisskeyUser `json:"user"`
}
