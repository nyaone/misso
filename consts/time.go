package consts

import "time"

const (
	TIME_REQUEST_VALID = 1 * time.Hour

	TIME_LOGIN_REMEMBER = 10 * time.Minute
	TIME_CONSENT_REMEMBER = 0 // Forever

	TIME_USERINFO_CACHE = 10 * time.Minute
)
