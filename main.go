package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonyalaribe/440sites/config"
	"github.com/tonyalaribe/440sites/web"
)

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config.Init() //Init Config.yaml

	router := web.StartRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")
		PORT = "8081"
	}

	log.Println("serving at :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
