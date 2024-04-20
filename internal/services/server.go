package services

import (
	"allowsync/allowsync/internal/config"
	"allowsync/allowsync/internal/transport/rest"
	"log"
	"net/http"
)

func SetupServer() {
	s := &http.Server{
		Addr: config.CONFIG.Address + ":" + config.CONFIG.Port,
	}
	http.HandleFunc("/ping", rest.Ping)
	log.Fatal(s.ListenAndServe())
}
