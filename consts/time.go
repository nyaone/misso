package consts

import "time"

const (
	TIME_LOGIN_REQUEST_VALID = 10 * time.Minute
	TIME_LOGIN_SESSION_VALID = 7 * 24 * time.Hour

	TIME_CONSENT_REQUEST_VALID = 1 * time.Hour
	TIME_CONSENT_SESSION_VALID = 30 * 24 * time.Hour
)
