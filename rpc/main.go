package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type SiteInfo struct {
	Dir string
}

type Args struct {
}

type Reply struct {
}

var siteInfo = SiteInfo{}

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (*SiteInfo) NewShopHandler(*Args, *Reply) {

}

func main() {
	siteInfo.Dir = "../sites"
	siteInfo.NewSite("past3")

	rpc.Register(siteInfo)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
