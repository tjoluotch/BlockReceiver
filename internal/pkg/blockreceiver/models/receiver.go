package models

import pb "BlockReceiver/internal/pkg/configs/grpc"

type BlockReceiverServer struct {
	pb.UnimplementedBlockResolverServer
}
