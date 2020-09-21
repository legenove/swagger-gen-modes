package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getStoreOrderOrderIdDecors = []grpccore.GrpcDecoratorFunc{}

func getStoreOrderOrderId(ctx context.Context, req *pb.GetStoreOrderOrderIdRequest) (*pb.GetStoreOrderOrderIdReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

