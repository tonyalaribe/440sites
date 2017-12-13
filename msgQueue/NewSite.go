package msgQueue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/tonyalaribe/440sites/config"
	"github.com/tonyalaribe/440sites/utils"
	"github.com/tonyalaribe/shop440/shops"
)

func NewSiteHandler(msg *stan.Msg) {
	// Handle the message
	log.Printf(" DATA:  %+v\n", msg)
	shop := shops.Shop{}
	json.Unmarshal(msg.Data, &shop)
	err := NewSite(shop.ShopID)
	if err != nil {
		log.Println(err)

	}

}

//NewSite runs the hugo new command and generates a new site at given location
func NewSite(shopid string) error {
	// resp := exec.Command("hugo", "new", "site", shopid)
	siteDir := config.Get().SitesDir
	cmd := exec.Command("hugo", "new", "site", shopid)
	cmd.Dir = siteDir
	respByt, err := cmd.Output()
	if err != nil {
		log.Println(err)
		// return err
	}
	log.Println(string(respByt))

	destinationURL := filepath.Join(siteDir, shopid, "themes", "shop440-default")
	err = utils.CopyDir("./defaults/shop440-default", destinationURL)
	if err != nil {
		log.Println(err)
		// return err
	}

	var conf map[string]interface{}
	configTomlFile := filepath.Join(siteDir, shopid, "config.toml")
	if _, err := toml.DecodeFile(configTomlFile, &conf); err != nil {
		// handle error
		log.Println(err)
	}

	conf["theme"] = "shop440-default"
	conf["baseURL"] = "/"

	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(conf)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(configTomlFile, buf.Bytes(), 0655)
	if err != nil {
		log.Println(err)
	}

	log.Println(conf)

	log.Println("just before running hugo")

	cmd = exec.Command("hugo")
	cmd.Dir = filepath.Join(siteDir, shopid)
	respByt, err = cmd.Output()
	if err != nil {
		log.Println(err)
		// return err
	}

	log.Println(string(respByt))
	return nil
}
