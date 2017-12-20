package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonyalaribe/440sites/config"
	"github.com/tonyalaribe/440sites/rpc"
	"github.com/tonyalaribe/440sites/web"
)

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config.Init() //Init Config.yaml
	// msgQueue.Init()
	RPC_PORT := os.Getenv("RPC_PORT")
	if RPC_PORT == "" {
		RPC_PORT = "9001"
	}
	log.Println("RPC port:", RPC_PORT)
	go rpc.StartRPCServer(RPC_PORT)

	router := web.StartRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")
		PORT = "8081"
	}

	log.Println("serving at :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
