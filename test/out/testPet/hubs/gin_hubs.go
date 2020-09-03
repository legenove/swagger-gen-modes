package hubs
/*
### DO NOT CHANGE THIS FILE
### The code is auto generated, your change will be overwritten by
### code generating.
*/
import (
    "context"
    pb "/testPet/pb"
    "/testPet/schemas"
    "/testPet/services"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc/metadata"
)

var server = services.NewServer()

func PostPet(c *gin.Context) (int, interface{}) {
    in := new(pb.PostPetRequest)
    headers, err := schemas.GetPostPetParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostPet(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func PutPet(c *gin.Context) (int, interface{}) {
    in := new(pb.PutPetRequest)
    headers, err := schemas.GetPutPetParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PutPet(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func GetPetFindByStatus(c *gin.Context) (int, interface{}) {
    in := new(pb.GetPetFindByStatusRequest)
    headers, err := schemas.GetGetPetFindByStatusParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetPetFindByStatus(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func GetPetFindByTags(c *gin.Context) (int, interface{}) {
    in := new(pb.GetPetFindByTagsRequest)
    headers, err := schemas.GetGetPetFindByTagsParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetPetFindByTags(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func DeletePetPetId(c *gin.Context) (int, interface{}) {
    in := new(pb.DeletePetPetIdRequest)
    headers, err := schemas.GetDeletePetPetIdParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.DeletePetPetId(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func GetPetPetId(c *gin.Context) (int, interface{}) {
    in := new(pb.GetPetPetIdRequest)
    headers, err := schemas.GetGetPetPetIdParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetPetPetId(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func PostPetPetId(c *gin.Context) (int, interface{}) {
    in := new(pb.PostPetPetIdRequest)
    headers, err := schemas.GetPostPetPetIdParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostPetPetId(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func PostPetPetIdUploadImage(c *gin.Context) (int, interface{}) {
    in := new(pb.PostPetPetIdUploadImageRequest)
    headers, err := schemas.GetPostPetPetIdUploadImageParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostPetPetIdUploadImage(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func GetStoreInventory(c *gin.Context) (int, interface{}) {
    in := new(pb.EmptyMessage)
    headers, err := schemas.GetGetStoreInventoryParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetStoreInventory(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func PostStoreOrder(c *gin.Context) (int, interface{}) {
    in := new(pb.PostStoreOrderRequest)
    headers, err := schemas.GetPostStoreOrderParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostStoreOrder(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func DeleteStoreOrderOrderId(c *gin.Context) (int, interface{}) {
    in := new(pb.DeleteStoreOrderOrderIdRequest)
    headers, err := schemas.GetDeleteStoreOrderOrderIdParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.DeleteStoreOrderOrderId(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func GetStoreOrderOrderId(c *gin.Context) (int, interface{}) {
    in := new(pb.GetStoreOrderOrderIdRequest)
    headers, err := schemas.GetGetStoreOrderOrderIdParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetStoreOrderOrderId(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func PostUser(c *gin.Context) (int, interface{}) {
    in := new(pb.PostUserRequest)
    headers, err := schemas.GetPostUserParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostUser(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func PostUserCreateWithArray(c *gin.Context) (int, interface{}) {
    in := new(pb.PostUserCreateWithArrayRequest)
    headers, err := schemas.GetPostUserCreateWithArrayParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostUserCreateWithArray(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func PostUserCreateWithList(c *gin.Context) (int, interface{}) {
    in := new(pb.PostUserCreateWithListRequest)
    headers, err := schemas.GetPostUserCreateWithListParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PostUserCreateWithList(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func GetUserLogin(c *gin.Context) (int, interface{}) {
    in := new(pb.GetUserLoginRequest)
    headers, err := schemas.GetGetUserLoginParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetUserLogin(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func GetUserLogout(c *gin.Context) (int, interface{}) {
    in := new(pb.EmptyMessage)
    headers, err := schemas.GetGetUserLogoutParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetUserLogout(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

func DeleteUserUsername(c *gin.Context) (int, interface{}) {
    in := new(pb.DeleteUserUsernameRequest)
    headers, err := schemas.GetDeleteUserUsernameParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.DeleteUserUsername(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func GetUserUsername(c *gin.Context) (int, interface{}) {
    in := new(pb.GetUserUsernameRequest)
    headers, err := schemas.GetGetUserUsernameParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.GetUserUsername(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), res.Data
}

func PutUserUsername(c *gin.Context) (int, interface{}) {
    in := new(pb.PutUserUsernameRequest)
    headers, err := schemas.GetPutUserUsernameParams(c, in)
    // header设置
    ctx := metadata.NewOutgoingContext(context.Background(), headers)
    res, err := server.PutUserUsername(ctx, in)
    if err != nil {
        panic(err)
    }
    if res == nil {
        return 200, nil
    }
    return int(res.HttpCode), nil
}

