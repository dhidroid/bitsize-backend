package config

import "fmt"

type Config struct {
	db, port string
}

// AppConfig is the global configuration instance.
var AppConfig *Config

func init() {
	AppConfig = &Config{
		db:   "https://dhidroid.vercel.app",
		port: "8080",
	}
}

// GetDB returns the database URL.
func (c *Config) GetDB() string {
	fmt.Println("Getting database URL...")
	return c.db
}

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) SetDB(db string) {
	c.db = db
}

func (c *Config) SetPort(port string) {
	c.port = port
}
