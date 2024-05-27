package mysql

import "errors"

type Config struct {
	Host     string `mapstructure:"MYSQLHOST"`
	Password string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	User     string `mapstructure:"MYSQLUSER"`
	DBName   string `mapstructure:"MYSQLDATABASE"`
	Port     string `mapstructure:"MYSQLPORT"`
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("DB Connect Host is empty")
	}
	return nil
}
