/*
### DO NOT CHANGE THIS FILE
### The code is auto generated, your change will be overwritten by
### code generating.
*/
package services

import (
    "context"

    pb "github.com/legenove/swagger-gen-modes/test/out/testPet/pb"
)

//server is used to implement testPet.TestPetServer.
type testPetServer struct {
  pb.UnimplementedTestPetServer
}
func NewServer() *testPetServer{
    return &testPetServer{}
}

func (s *testPetServer) PostPet(ctx context.Context, req *pb.PostPetRequest) (*pb.CommonReply, error) {
	res, err := postPetHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var postPetHandler = decoratorHandler("postPet", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postPet(ctx, req.(*pb.PostPetRequest))
}, postPetDecors...)


func (s *testPetServer) PutPet(ctx context.Context, req *pb.PutPetRequest) (*pb.CommonReply, error) {
	res, err := putPetHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var putPetHandler = decoratorHandler("putPet", func(ctx context.Context, req interface{}) (interface{}, error) {
	return putPet(ctx, req.(*pb.PutPetRequest))
}, putPetDecors...)


func (s *testPetServer) GetPetFindByStatus(ctx context.Context, req *pb.GetPetFindByStatusRequest) (*pb.GetPetFindByStatusReply, error) {
	res, err := getPetFindByStatusHandler(ctx, req)
	var _r = res.(*pb.GetPetFindByStatusReply)
	if _r == nil {
		_r = &pb.GetPetFindByStatusReply{}
	}
	return _r, err
}

var getPetFindByStatusHandler = decoratorHandler("getPetFindByStatus", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getPetFindByStatus(ctx, req.(*pb.GetPetFindByStatusRequest))
}, getPetFindByStatusDecors...)


func (s *testPetServer) GetPetFindByTags(ctx context.Context, req *pb.GetPetFindByTagsRequest) (*pb.GetPetFindByTagsReply, error) {
	res, err := getPetFindByTagsHandler(ctx, req)
	var _r = res.(*pb.GetPetFindByTagsReply)
	if _r == nil {
		_r = &pb.GetPetFindByTagsReply{}
	}
	return _r, err
}

var getPetFindByTagsHandler = decoratorHandler("getPetFindByTags", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getPetFindByTags(ctx, req.(*pb.GetPetFindByTagsRequest))
}, getPetFindByTagsDecors...)


func (s *testPetServer) DeletePetPetId(ctx context.Context, req *pb.DeletePetPetIdRequest) (*pb.CommonReply, error) {
	res, err := deletePetPetIdHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var deletePetPetIdHandler = decoratorHandler("deletePetPetId", func(ctx context.Context, req interface{}) (interface{}, error) {
	return deletePetPetId(ctx, req.(*pb.DeletePetPetIdRequest))
}, deletePetPetIdDecors...)


func (s *testPetServer) GetPetPetId(ctx context.Context, req *pb.GetPetPetIdRequest) (*pb.GetPetPetIdReply, error) {
	res, err := getPetPetIdHandler(ctx, req)
	var _r = res.(*pb.GetPetPetIdReply)
	if _r == nil {
		_r = &pb.GetPetPetIdReply{}
	}
	return _r, err
}

var getPetPetIdHandler = decoratorHandler("getPetPetId", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getPetPetId(ctx, req.(*pb.GetPetPetIdRequest))
}, getPetPetIdDecors...)


func (s *testPetServer) PostPetPetId(ctx context.Context, req *pb.PostPetPetIdRequest) (*pb.CommonReply, error) {
	res, err := postPetPetIdHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var postPetPetIdHandler = decoratorHandler("postPetPetId", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postPetPetId(ctx, req.(*pb.PostPetPetIdRequest))
}, postPetPetIdDecors...)


func (s *testPetServer) PostPetPetIdUploadImage(ctx context.Context, req *pb.PostPetPetIdUploadImageRequest) (*pb.PostPetPetIdUploadImageReply, error) {
	res, err := postPetPetIdUploadImageHandler(ctx, req)
	var _r = res.(*pb.PostPetPetIdUploadImageReply)
	if _r == nil {
		_r = &pb.PostPetPetIdUploadImageReply{}
	}
	return _r, err
}

var postPetPetIdUploadImageHandler = decoratorHandler("postPetPetIdUploadImage", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postPetPetIdUploadImage(ctx, req.(*pb.PostPetPetIdUploadImageRequest))
}, postPetPetIdUploadImageDecors...)


func (s *testPetServer) GetStoreInventory(ctx context.Context, req *pb.EmptyMessage) (*pb.GetStoreInventoryReply, error) {
	res, err := getStoreInventoryHandler(ctx, req)
	var _r = res.(*pb.GetStoreInventoryReply)
	if _r == nil {
		_r = &pb.GetStoreInventoryReply{}
	}
	return _r, err
}

var getStoreInventoryHandler = decoratorHandler("getStoreInventory", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getStoreInventory(ctx, req.(*pb.EmptyMessage))
}, getStoreInventoryDecors...)


func (s *testPetServer) PostStoreOrder(ctx context.Context, req *pb.PostStoreOrderRequest) (*pb.PostStoreOrderReply, error) {
	res, err := postStoreOrderHandler(ctx, req)
	var _r = res.(*pb.PostStoreOrderReply)
	if _r == nil {
		_r = &pb.PostStoreOrderReply{}
	}
	return _r, err
}

var postStoreOrderHandler = decoratorHandler("postStoreOrder", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postStoreOrder(ctx, req.(*pb.PostStoreOrderRequest))
}, postStoreOrderDecors...)


func (s *testPetServer) DeleteStoreOrderOrderId(ctx context.Context, req *pb.DeleteStoreOrderOrderIdRequest) (*pb.CommonReply, error) {
	res, err := deleteStoreOrderOrderIdHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var deleteStoreOrderOrderIdHandler = decoratorHandler("deleteStoreOrderOrderId", func(ctx context.Context, req interface{}) (interface{}, error) {
	return deleteStoreOrderOrderId(ctx, req.(*pb.DeleteStoreOrderOrderIdRequest))
}, deleteStoreOrderOrderIdDecors...)


func (s *testPetServer) GetStoreOrderOrderId(ctx context.Context, req *pb.GetStoreOrderOrderIdRequest) (*pb.GetStoreOrderOrderIdReply, error) {
	res, err := getStoreOrderOrderIdHandler(ctx, req)
	var _r = res.(*pb.GetStoreOrderOrderIdReply)
	if _r == nil {
		_r = &pb.GetStoreOrderOrderIdReply{}
	}
	return _r, err
}

var getStoreOrderOrderIdHandler = decoratorHandler("getStoreOrderOrderId", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getStoreOrderOrderId(ctx, req.(*pb.GetStoreOrderOrderIdRequest))
}, getStoreOrderOrderIdDecors...)


func (s *testPetServer) PostUser(ctx context.Context, req *pb.PostUserRequest) (*pb.CommonReply, error) {
	res, err := postUserHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var postUserHandler = decoratorHandler("postUser", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postUser(ctx, req.(*pb.PostUserRequest))
}, postUserDecors...)


func (s *testPetServer) PostUserCreateWithArray(ctx context.Context, req *pb.PostUserCreateWithArrayRequest) (*pb.CommonReply, error) {
	res, err := postUserCreateWithArrayHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var postUserCreateWithArrayHandler = decoratorHandler("postUserCreateWithArray", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postUserCreateWithArray(ctx, req.(*pb.PostUserCreateWithArrayRequest))
}, postUserCreateWithArrayDecors...)


func (s *testPetServer) PostUserCreateWithList(ctx context.Context, req *pb.PostUserCreateWithListRequest) (*pb.CommonReply, error) {
	res, err := postUserCreateWithListHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var postUserCreateWithListHandler = decoratorHandler("postUserCreateWithList", func(ctx context.Context, req interface{}) (interface{}, error) {
	return postUserCreateWithList(ctx, req.(*pb.PostUserCreateWithListRequest))
}, postUserCreateWithListDecors...)


func (s *testPetServer) GetUserLogin(ctx context.Context, req *pb.GetUserLoginRequest) (*pb.GetUserLoginReply, error) {
	res, err := getUserLoginHandler(ctx, req)
	var _r = res.(*pb.GetUserLoginReply)
	if _r == nil {
		_r = &pb.GetUserLoginReply{}
	}
	return _r, err
}

var getUserLoginHandler = decoratorHandler("getUserLogin", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getUserLogin(ctx, req.(*pb.GetUserLoginRequest))
}, getUserLoginDecors...)


func (s *testPetServer) GetUserLogout(ctx context.Context, req *pb.EmptyMessage) (*pb.CommonReply, error) {
	res, err := getUserLogoutHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var getUserLogoutHandler = decoratorHandler("getUserLogout", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getUserLogout(ctx, req.(*pb.EmptyMessage))
}, getUserLogoutDecors...)


func (s *testPetServer) DeleteUserUsername(ctx context.Context, req *pb.DeleteUserUsernameRequest) (*pb.DeleteUserUsernameReply, error) {
	res, err := deleteUserUsernameHandler(ctx, req)
	var _r = res.(*pb.DeleteUserUsernameReply)
	if _r == nil {
		_r = &pb.DeleteUserUsernameReply{}
	}
	return _r, err
}

var deleteUserUsernameHandler = decoratorHandler("deleteUserUsername", func(ctx context.Context, req interface{}) (interface{}, error) {
	return deleteUserUsername(ctx, req.(*pb.DeleteUserUsernameRequest))
}, deleteUserUsernameDecors...)


func (s *testPetServer) GetUserUsername(ctx context.Context, req *pb.GetUserUsernameRequest) (*pb.GetUserUsernameReply, error) {
	res, err := getUserUsernameHandler(ctx, req)
	var _r = res.(*pb.GetUserUsernameReply)
	if _r == nil {
		_r = &pb.GetUserUsernameReply{}
	}
	return _r, err
}

var getUserUsernameHandler = decoratorHandler("getUserUsername", func(ctx context.Context, req interface{}) (interface{}, error) {
	return getUserUsername(ctx, req.(*pb.GetUserUsernameRequest))
}, getUserUsernameDecors...)


func (s *testPetServer) PutUserUsername(ctx context.Context, req *pb.PutUserUsernameRequest) (*pb.CommonReply, error) {
	res, err := putUserUsernameHandler(ctx, req)
	var _r = res.(*pb.CommonReply)
	if _r == nil {
		_r = &pb.CommonReply{}
	}
	return _r, err
}

var putUserUsernameHandler = decoratorHandler("putUserUsername", func(ctx context.Context, req interface{}) (interface{}, error) {
	return putUserUsername(ctx, req.(*pb.PutUserUsernameRequest))
}, putUserUsernameDecors...)


