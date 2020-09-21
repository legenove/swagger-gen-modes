package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getStoreInventoryDecors = []grpccore.GrpcDecoratorFunc{}

func getStoreInventory(ctx context.Context, req *pb.EmptyMessage) (*pb.GetStoreInventoryReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

