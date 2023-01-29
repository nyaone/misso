package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"misso/consts"
	"misso/global"
	"misso/misskey"
	"misso/types"
)

func GetUserinfo(subject string) (*types.MisskeyUser, error) {
	// Check cache key
	global.Logger.Debugf("Checking userinfo cache...")
	userinfoCacheKey := fmt.Sprintf(consts.REDIS_KEY_USER_INFO, subject)
	exist, err := global.Redis.Exists(context.Background(), userinfoCacheKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to check userinfo exist status with error: %v", err)
	} else if exist > 0 {
		// Exist, get from cache
		userinfoCacheBytes, err := global.Redis.Get(context.Background(), userinfoCacheKey).Bytes()
		if err != nil {
			global.Logger.Errorf("Userinfo cache exists, but failed to retrieve with error: %v", err)
		} else {
			// Parse into user
			var userinfo types.MisskeyUser
			err = json.Unmarshal(userinfoCacheBytes, &userinfo)
			if err != nil {
				global.Logger.Errorf("Failed to parse userinfo cache into json with error: %v", err)
			} else {
				// Works!
				global.Logger.Debugf("Get cached userinfo successfully.")
				return &userinfo, nil
			}
		}
	}

	// Fallback to get info directly, we need user's access token.
	global.Logger.Debugf("No cached userinfo found (or valid), trying to get latest response.")
	accessTokenCacheKey := fmt.Sprintf(consts.REDIS_KEY_USER_ACCESS_TOKEN, subject)
	accessToken, err := global.Redis.Get(context.Background(), accessTokenCacheKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to get user access token with error: %v", err)
		return nil, err
	}

	// Get userinfo with access token
	userinfo, err := misskey.GetUserinfo(accessToken)
	if err != nil {
		global.Logger.Errorf("Failed to get user info with saved access token with error: %v", err)
		return nil, err
	}

	// Append subject as email to userinfo
	(*userinfo)["email"] = subject

	// Save userinfo into redis
	_ = SaveUserinfo(subject, userinfo) // Ignore errors

	return userinfo, nil

}
