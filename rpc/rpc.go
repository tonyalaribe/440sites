package rpc

import (
	"log"
	"net"

	pb "github.com/tonyalaribe/440sites/rpc/s440proto"
	"google.golang.org/grpc"
)

type s440server struct {
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
