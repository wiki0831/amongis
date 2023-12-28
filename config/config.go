package config

import (
	"flag"

	"github.com/joho/godotenv"
)

// static vars
var (
	Env  = flag.String("envfile", "./config/.env", "envfile path")
	PORT = flag.String("port", "3000", "api port")
)

func init() {
	// Parse Flags
	flag.Parse()
	// Load .env file
	err := godotenv.Load(*Env)
	if err != nil {
		panic("Error loading .env file")
	}
}
