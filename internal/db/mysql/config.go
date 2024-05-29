package mysql

import (
	"errors"
	"os"
)

type Config struct {
	Host     string `mapstructure:"MYSQLHOST"`
	Password string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	User     string `mapstructure:"MYSQLUSER"`
	DBName   string `mapstructure:"MYSQLDATABASE"`
	Port     string `mapstructure:"MYSQLPORT"`
	Url      string
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.New("DB Connect Host is empty")
	}
	return nil
}

func (c *Config) LoadFromEnv() error {
	// // dbName := os.Getenv("MYSQLDATABASE")
	// // if dbName == "" {
	// // 	return errors.New("MYSQLDATABASE is empty")
	// // }
	// // c.DBName = dbName
	// // password := os.Getenv("MYSQL_ROOT_PASSWORD")
	// // if password == "" {
	// // 	return errors.New("MYSQL_ROOT_PASSWORD empty")
	// // }
	// // c.Password = password
	// // host := os.Getenv("MYSQLHOST")
	// // if host == "" {
	// // 	return errors.New("MYSQLHOST empty")
	// // }
	// // c.Host = host
	// // user := os.Getenv("MYSQLUSER")
	// // if user == "" {
	// // 	return errors.New("MYSQLUSER empty")
	// // }
	// // c.User = user
	// // port := os.Getenv("MYSQLPORT")
	// // if port == "" {
	// // 	return errors.New("MYSQLPORT empty")
	// // }
	// c.Port = port

	c.Url = os.Getenv("DATABASE_PRIVATE_URL")
	if c.Url == "" {
		return errors.New("DATABASE_PRIVATE_URL empty")
	}
	return nil
}
