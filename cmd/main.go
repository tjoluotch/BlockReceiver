package main

import (
	pb "BlockReceiver/internal/pkg/configs/grpc"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	NETWORK      = "tcp"
	PORT     int = 50051
	HOSTNAME     = "localhost"
)

type blockReceiverServer struct {
	pb.UnimplementedBlockResolverServer
}

func getAddress() string {
	return fmt.Sprintf("%s:%d", HOSTNAME, PORT)
}

func newServer() *blockReceiverServer {
	return &blockReceiverServer{}
}

// SendBlock generate a new uuid based on each block's hash and returns this as a BlockId object
func (s *blockReceiverServer) SendBlock(ctx context.Context, block *pb.Block) (*pb.BlockId, error) {
	log.Printf("received Block: %+v\n", block)
	id := uuid.New()
	log.Printf("Generated Block ID:%s\n", id.String())
	log.Println("Sending Block ID")
	log.Printf("\n")
	return &pb.BlockId{Id: id.String()}, nil
}

func main() {
	listener, err := net.Listen(NETWORK, getAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	log.Printf("%s\n", "GRPC server startup...")
	log.Printf("%s\n", listener.Addr().String())
	pb.RegisterBlockResolverServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}
