package config

type AppConfig struct {
    Port     string `yaml:"port"`
    APIKey   string `yaml:"api_key"`
    Login    string `yaml:"login"`
    Password string `yaml:"password"`
}



type DatabasesConfig []struct {
		Name     string
        Type     string
		path string
		shedule string
    }
	
func MustConfig(config *Config) err error {

}