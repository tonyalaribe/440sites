package models

import (
	"fmt"
	"log"
	"os"

	"github.com/tonyalaribe/440sites/config"
)

func AddCustomDomain(shopid, domain string) error {
	log.Print(shopid, domain)
	err := config.Get().SetBadgerKV(config.DOMAINSCOLLECTION, []byte(domain), []byte(shopid), 0x00)
	if err != nil {
		log.Println(err)
	}

	filename := "./customsites.conf"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	confString := fmt.Sprintf(`
  %s{
    proxy / localhost:8007
    root /home/ubuntu/server/440sites
  }
  `, domain)
	if _, err = f.WriteString(confString); err != nil {
		log.Println(err)
	}

	return nil
}
