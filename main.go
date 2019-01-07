package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

const (
	defaultBind = ":9999"
)

var (
	version = "0.1.0"
)

func main() {
	bind := os.Getenv("RAAS_BIND")
	if bind == "" {
		bind = defaultBind
	}

	router := httprouter.New()

	router.POST("/raas/subscribe", subscribeHandler)
	router.POST("/raas/publish", publishHandler)

	log.Printf("Starting raas v%s on %s\n", version, bind)

	log.Fatal(http.ListenAndServe(bind, router))
}
