package msgQueue

import (
	"encoding/json"
	"log"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/tonyalaribe/440sites/models"
	"github.com/tonyalaribe/shop440/shops"
)

func NewSiteHandler(msg *stan.Msg) {
	// Handle the message
	log.Printf(" DATA:  %+v\n", msg)
	shop := shops.Shop{}
	json.Unmarshal(msg.Data, &shop)
	err := models.NewSite(shop.ShopID)
	if err != nil {
		log.Println(err)

	}
}
