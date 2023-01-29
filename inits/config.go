package inits

import (
	"gopkg.in/yaml.v3"
	"misso/config"
	"misso/consts"
	"os"
)

func Config() error {
	// Read config file
	configFilePosition, exist := os.LookupEnv("CONFIG_FILE_PATH")
	if !exist {
		configFilePosition = "config.yml"
	}

	configFileBytes, err := os.ReadFile(configFilePosition)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFileBytes, &config.Config)
	if err != nil {
		return err
	}

	// Validate time
	if config.Config.Time.RequestValid <= 0 {
		config.Config.Time.RequestValid = consts.TIME_DEFAULT_REQUEST_VALID
	}
	if config.Config.Time.LoginRemember < 0 {
		// 0 means don't remember (in extreme account switch situations)
		config.Config.Time.LoginRemember = consts.TIME_DEFAULT_LOGIN_REMEMBER
	}
	if config.Config.Time.ConsentRemember < 0 {
		// 0 means remember forever (default behavior)
		config.Config.Time.ConsentRemember = consts.TIME_DEFAULT_CONSENT_REMEMBER
	}
	if config.Config.Time.UserinfoCache <= 0 {
		config.Config.Time.UserinfoCache = consts.TIME_DEFAULT_USERINFO_CACHE
	}

	return nil
}
