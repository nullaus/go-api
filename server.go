package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TV4/graceful"
	"github.com/nullaus/go-api/utils"
)

const (
	EnvServerAddr    = "SERVER_ADDR"
	EnvServerUseTLS  = "SERVER_USE_TLS"
	EnvServerTLSCert = "SERVER_TLS_CERT"
	EnvServerTLSKey  = "SERVER_TLS_KEY"
)

func startServer() {
	api := newAPI()
	api, err := addRoutes(api)
	if err != nil {
		log.Fatal(err)
	}

	addr := utils.MustGetEnv(EnvServerAddr)
	useTLS := os.Getenv(EnvServerUseTLS) == "1"

	server := &http.Server{
		Addr:    addr,
		Handler: api.MakeHandler(),
	}

	if useTLS {
		graceful.ListenAndServeTLS(server, utils.MustGetEnv(EnvServerTLSCert), utils.MustGetEnv(EnvServerTLSKey))
	} else {
		graceful.ListenAndServe(server)
	}
}
