package consts

const (
	REDIS_KEY_LOGIN_SESSION = "misso:login:%s" // Username, token as value
	REDIS_KEY_CONSENT_CSRF  = "misso:consent:%s"
	REDIS_KEY_SHARE_CONTEXT = "misso:share:%s" // Subject, context as value
)
