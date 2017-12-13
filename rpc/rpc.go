package rpc

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	// "google/protobuf/empty.proto"
	pb "github.com/tonyalaribe/440sites/rpc/s440proto"
)

type s440server struct {
}

func (s s440server) CreateSiteHandler(ctx context.Context, shop *pb.Shop) (*pb.Response, error) {
	log.Printf("%#v", ctx)
	log.Printf("%#v", shop)

	log.Println("create site")
	err := NewSite(shop.Slug)
	message := ""
	if err != nil {
		message = err.Error()
	}

	return &pb.Response{Message: message}, nil
}

func init() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func StartRPCServer(PORT string) {
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	// pb.
	pb.RegisterS440SiteServer(server, s440server{})

	server.Serve(lis)
}
