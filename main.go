package main

import "log"

type SiteInfo struct {
	Dir string
}

var siteInfo = SiteInfo{}

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	siteInfo.Dir = "./sites"
	siteInfo.NewSite("past3")
}
