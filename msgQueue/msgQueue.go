package msgQueue

import (
	"log"

	"github.com/tonyalaribe/440sites/config"
)

func Init() {
	log.Println("Registering Nats handlers")
	config.Get().Nats.Subscribe("shop440.site.create", NewSiteHandler)
	config.Get().Nats.Subscribe("shop440.site.add_domain", AddADomainHandler)
}
