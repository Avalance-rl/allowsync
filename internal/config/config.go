package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var CONFIG config

type config struct {
	HTTPServer
}

type HTTPServer struct {
	Address string
	Port    string
}

func MustLoadConfig() config {
	var c config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.Address = os.Getenv("SERVER_ADDRESS")
	c.Port = os.Getenv("SERVER_PORT")
	return c
}
