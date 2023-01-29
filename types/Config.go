package types

type Config struct {
	System struct {
		Debug bool   `yaml:"debug"`
		Redis string `yaml:"redis"`
	} `yaml:"system"`
	Misskey struct {
		Instance    string `yaml:"instance"`
		Application struct {
			Secret string `yaml:"secret"`
		} `yaml:"application"`
	} `yaml:"misskey"`
	Hydra struct {
		AdminUrl string `yaml:"admin_url"`
	} `yaml:"hydra"`
	Time struct {
		RequestValid    int64 `yaml:"request_valid"`
		LoginRemember   int64 `yaml:"login_remember"`
		ConsentRemember int64 `yaml:"consent_remember"`
		UserinfoCache   int64 `yaml:"userinfo_cache"`
	} `yaml:"time"`
}
