package rpc

import (
	"log"

	"golang.org/x/net/context"

	"github.com/tonyalaribe/440sites/models"
	pb "github.com/tonyalaribe/440sites/rpc/s440proto"
)

func (s s440server) CreateSiteHandler(ctx context.Context, shop *pb.Shop) (*pb.Response, error) {
	log.Printf("%#v", ctx)
	log.Printf("%#v", shop)

	log.Println("create site")
	err := models.NewSite(shop.Slug)
	message := ""
	if err != nil {
		message = err.Error()
	}

	return &pb.Response{Message: message}, nil
}
