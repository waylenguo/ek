package docker

import (
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/ek/pkg/aws"
	"github.com/spf13/viper"
)

type Config struct {
	Repositories []*Repository `yaml:"repositories"`
}

type Repository struct {
	Host string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password,omitempty"`
	Type string `yaml:"type,omitempty"`
}

func GetAuthMapping(configPath string) map[string]string {
	authMapping := make(map[string]string)
	config := loadConfig(configPath)
	for _, repository := range config.Repositories {
		if repository.Type != "" && repository.Type == "AWS"{
			password := aws.NewEcr().GetLoginPassword()
			repository.Password = *password
		}
		auth := types.AuthConfig{
			Username: repository.Username,
			Password: repository.Password,
		}
		encodedJSON, err := json.Marshal(auth)
		if err != nil {
			panic("auth encoded error")
		}
		authStr := base64.URLEncoding.EncodeToString(encodedJSON)
		authMapping[repository.Host] = authStr
	}
	return authMapping
}


func loadConfig(path string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	var _config *Config
	err = viper.Unmarshal(&_config)
	if err != nil {
		panic(err.Error())
	}

	return _config
}
