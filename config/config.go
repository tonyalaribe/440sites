package config

import (
	"log"
	"os"

	"github.com/dgraph-io/badger"
	stan "github.com/nats-io/go-nats-streaming"
)

const (
	SERVER_NATS_CLIENTID     = "shop440-sites"
	NATS_STREAMING_CLUSTERID = "shop440"
)

type Conf struct {
	Nats     stan.Conn
	SitesDir string

	BadgerDB map[string]*badger.DB
	// BleveIndexes map[string]bleve.Index
}

var config = Conf{}

func Init() {
	var err error
	config.SitesDir = os.Getenv("SITES_ROOT")
	if config.SitesDir == "" {
		config.SitesDir = "./sites"
	}

	config.Nats, err = stan.Connect(NATS_STREAMING_CLUSTERID, SERVER_NATS_CLIENTID)
	if err != nil {
		log.Fatalf("Error starting a nats connection. Error: %v", err)
		return
	}
}

func Get() *Conf {
	return &config
}
