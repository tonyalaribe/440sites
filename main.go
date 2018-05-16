package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tonyalaribe/440sites/config"
	"github.com/tonyalaribe/440sites/msgQueue"
	"github.com/tonyalaribe/440sites/web"
)

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config.Init() //Init Config.yaml
	msgQueue.Init()

	var isTest = flag.Bool("test", false, "Set test mode to use ephemeral data storage")
	flag.Parse()

	if *isTest {
		log.Println("In test mode")
		os.RemoveAll("./" + "test_sites")
		os.MkdirAll("./test_sites", os.ModePerm)
		config.Get().SitesDir = "test_sites"

		defer fmt.Println("delete test_sites")
		defer os.RemoveAll("./" + config.Get().SitesDir)
	}

	router := web.StartRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")
		PORT = "8081"
	}

	log.Println("serving at :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
