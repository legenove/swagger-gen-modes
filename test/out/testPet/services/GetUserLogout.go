package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getUserLogoutDecors = []grpccore.GrpcDecoratorFunc{}

func getUserLogout(ctx context.Context, req *pb.EmptyMessage) (*pb.CommonReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

