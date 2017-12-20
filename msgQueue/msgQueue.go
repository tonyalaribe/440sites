package msgQueue

import "github.com/tonyalaribe/440sites/Config"

func Init() {
	config.Get().Nats.Subscribe("shop440.site.create", NewSiteHandler)
}
