package mysql

import (
	"errors"
	"os"
)

type Config struct {
	DBurl string
}

func (c *Config) Validate() error {
	if c.DBurl == "" {
		return errors.New("DB URL is empty")
	}
	return nil
}

func (c *Config) LoadFromEnv() error {

	c.DBurl = os.Getenv("DATABASE_PRIVATE_URL")
	if c.DBurl == "" {
		return errors.New("DATABASE_PRIVATE_URL empty")
	}
	return nil
}
