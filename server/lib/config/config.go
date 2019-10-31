package config

import (
	"fmt"
)

type Config struct {
	APIServerPort int `environment_key:"TORCH_API_PORT" type:"int"`

	// MySQL
	MySQLHost     string `environment_key:"TORCH_MYSQL_HOST" type:"string"`
	MySQLPort     int    `environment_key:"TORCH_MYSQL_PORT" type:"int"`
	MySQLUser     string `environment_key:"TORCH_API_USER" type:"string"`
	MySQLPassword string `environment_key:"TORCH_API_PASS" type:"string"`
}

func (config *Config) GetAPIServerLocation() string {
	loc := fmt.Sprintf(":%d", config.APIServerPort)
	return loc
}

func (config *Config) GetDBServerLocation() string {
	loc := fmt.Sprintf("%s:%s@tcp(%s:%d)/allergy?parseTime=true",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
	)

	return loc
}
