package msgQueue

import (
	"encoding/json"
	"log"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/tonyalaribe/440sites/models"
	"github.com/tonyalaribe/shop440/features/products"
	"github.com/tonyalaribe/shop440/features/shops"
)

func NewSiteHandler(msg *stan.Msg) {
	// Handle the message
	log.Printf(" new site for data DATA:  %+v\n", msg)
	shop := shops.Shop{}
	json.Unmarshal(msg.Data, &shop)
	err := models.NewSite(shop.ShopID)
	if err != nil {
		log.Println(err)
	}
}

func AddADomainHandler(msg *stan.Msg) {
	log.Println("add custom domain for new site")
	message := make(map[string]interface{})
	json.Unmarshal(msg.Data, &message)
	err := models.AddCustomDomain(message["shop_id"].(string), message["domain"].(string))
	if err != nil {
		log.Println(err)
	}
}

func NewProductHandler(msg *stan.Msg) {
	product := products.Product{}
	json.Unmarshal(msg.Data, &product)
	err := models.NewProduct(product)
	if err != nil {
		log.Println(err)
	}
}
