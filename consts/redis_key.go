package consts

const (
	REDIS_KEY_LOGIN_SESSION     = "misso:login:%s"      // Username, token as value
	REDIS_KEY_CONSENT_CSRF      = "misso:consent:%s"    // Random string, consent challenge as value
	REDIS_KEY_USER_ACCESS_TOKEN = "misso:user:token:%s" // Subject, access token as value
	REDIS_KEY_USER_INFO         = "misso:user:info:%s"  // Subject, user info as value
)
