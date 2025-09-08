package config

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type DatabaseConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  int           `mapstructure:"port"`
	User                  string        `mapstructure:"user"`
	Pass                  string        `mapstructure:"pass"`
	Name                  string        `mapstructure:"name"`
	SSL                   bool          `mapstructure:"ssl"`
	MaxOpenConnections    int           `mapstructure:"max_open_connections"`
	MaxIdleConnections    int           `mapstructure:"max_idle_connections"`
	ConnectionMaxLifeTime time.Duration `mapstructure:"connection_max_life_time"`
	ConnectionMaxIdleTime time.Duration `mapstructure:"connection_max_idle_time"`
}

func (c *DatabaseConfig) Validate() error {
	if c.Host == "" {
		return errors.New("host cannot be empty")
	}
	if c.Port == 0 {
		return errors.New("port cannot be empty")
	}
	if c.User == "" {
		return errors.New("user cannot be empty")
	}
	if c.Pass == "" {
		return errors.New("pass cannot be empty")
	}
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}
	if c.MaxOpenConnections <= 0 {
		c.MaxOpenConnections = 25
	}
	if c.MaxIdleConnections <= 0 {
		c.MaxIdleConnections = 5
	}
	if c.ConnectionMaxLifeTime <= 0 {
		c.ConnectionMaxLifeTime = 5 * time.Minute
	}
	if c.ConnectionMaxIdleTime <= 0 {
		c.ConnectionMaxIdleTime = 5 * time.Minute
	}
	return nil
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Pass, c.Name, c.SSL,
	)
}
