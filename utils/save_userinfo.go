package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"misso/config"
	"misso/consts"
	"misso/global"
	"misso/types"
	"time"
)

func SaveUserinfo(subject string, userinfo *types.MisskeyUser) error {
	userinfoBytes, err := json.Marshal(userinfo)
	if err != nil {
		global.Logger.Errorf("Failed to parse accept context with error: %v", err)
		return err
	}
	sessUserInfoKey := fmt.Sprintf(consts.REDIS_KEY_USER_INFO, subject)
	err = global.Redis.Set(context.Background(), sessUserInfoKey, userinfoBytes, time.Duration(config.Config.Time.UserinfoCache)*time.Second).Err()
	if err != nil {
		global.Logger.Errorf("Failed to save session user info into redis with error: %v", err)
		return err
	}

	return nil
}
