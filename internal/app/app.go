package app

import (
	"allowsync/allowsync/internal/config"
	"allowsync/allowsync/internal/services"
)

func init() {
	config.CONFIG = config.MustLoadConfig()
}

func Run() {
	services.SetupServer()
}
