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
}
