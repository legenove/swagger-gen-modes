package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var postStoreOrderDecors = []grpccore.GrpcDecoratorFunc{}

func postStoreOrder(ctx context.Context, req *pb.PostStoreOrderRequest) (*pb.PostStoreOrderReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

