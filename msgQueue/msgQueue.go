package msgQueue

import "github.com/tonyalaribe/440sites/config"

func Init() {
	config.Get().Nats.Subscribe("shop440.site.create", NewSiteHandler)
	config.Get().Nats.Subscribe("shop440.site.add_domain", AddADomainHandler)
}
