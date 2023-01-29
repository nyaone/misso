package types

type MisskeyUserBase struct {
	Username string `json:"username"`
	// Ignore other fields
}

type MisskeyUser = map[string]any // Just raw json map
