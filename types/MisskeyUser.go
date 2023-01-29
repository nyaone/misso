package types

type MisskeyUserBase struct {
	Username string `json:"username"`
	// Ignore other fields
}

type MisskeyUser = map[string]interface{} // Just raw json map
