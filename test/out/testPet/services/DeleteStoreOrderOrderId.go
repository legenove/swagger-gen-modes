package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var deleteStoreOrderOrderIdDecors = []grpccore.GrpcDecoratorFunc{}

func deleteStoreOrderOrderId(ctx context.Context, req *pb.DeleteStoreOrderOrderIdRequest) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

