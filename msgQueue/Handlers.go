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

func AddADomainHandler(msg *stan.Msg) {
	message := make(map[string]interface{})
	json.Unmarshal(msg.Data, &message)
	message["shop_id"]
	err := models.AddCustomDomain(message["shop_id"].(string), message["domain"].(string))
	if err != nil {
		log.Println(err)
	}
}
