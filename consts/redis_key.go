package consts

const (
	REDIS_KEY_LOGIN_SESSION      = "misso:login:%s"    // Username, token as value
	REDIS_KEY_CONSENT_CSRF       = "misso:consent:%s"  // Random string, consent challenge as value
	REDIS_KEY_SHARE_ACCESS_TOKEN = "misso:share:at:%s" // Subject, access token as value
	REDIS_KEY_SHARE_USER_INFO    = "misso:share:ui:%s" // Subject, user info as value
)
