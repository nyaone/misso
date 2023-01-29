package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"misso/consts"
	"misso/global"
	"misso/types"
)

func SaveUserinfo(subject string, userinfo *types.MisskeyUser) error {
	userinfoBytes, err := json.Marshal(userinfo)
	if err != nil {
		global.Logger.Errorf("Failed to parse accept context with error: %v", err)
		return err
	}
	sessUserInfoKey := fmt.Sprintf(consts.REDIS_KEY_USER_INFO, subject)
	err = global.Redis.Set(context.Background(), sessUserInfoKey, userinfoBytes, consts.TIME_USERINFO_CACHE).Err()
	if err != nil {
		global.Logger.Errorf("Failed to save session user info into redis with error: %v", err)
		return err
	}

	return nil
}
