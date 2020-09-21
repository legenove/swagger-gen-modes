package services

import (
    "context"
    "fmt"

    "github.com/legenove/nano-server-sdk/grpccore"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)


var getUserLoginDecors = []grpccore.GrpcDecoratorFunc{}

func getUserLogin(ctx context.Context, req *pb.GetUserLoginRequest) (*pb.GetUserLoginReply, error) {
	fmt.Println("in", req)
	return nil, nil
}

