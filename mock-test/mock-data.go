package mocktest

import (
	pbu "EXAM3/api-gateway/genproto/user_service"
	"context"

	"google.golang.org/grpc"
)

type UserServiceClient interface {
	CreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	GetUserByUsername(ctx context.Context, in *pbu.Username, opts ...grpc.CallOption) (*pbu.User, error)
	GetUserByEmail(ctx context.Context, in *pbu.Email, opts ...grpc.CallOption) (*pbu.User, error)
	UpdateUserById(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	GetUserById(ctx context.Context, in *pbu.UserId, opts ...grpc.CallOption) (*pbu.User, error)
	ListUser(ctx context.Context, in *pbu.GetAllUserRequest, opts ...grpc.CallOption) (*pbu.GetAllUserResponse, error)
	DeleteUser(ctx context.Context, in *pbu.UserId, opts ...grpc.CallOption) (*pbu.Empty, error)
	CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error)
}

type userServiceClient struct {
}

func NewUserServiceClient() UserServiceClient {
	return &userServiceClient{}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *pbu.Username, opts ...grpc.CallOption) (*pbu.User, error) {
	user := pbu.User{
		Id:       "dsfhdsjfhl",
		Name:     "Nodirbek",
		Age:      17,
		Username: "rarebek",
		Email:    "nomonovn2@gmail.com",
	}
	return &user, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *pbu.Email, opts ...grpc.CallOption) (*pbu.User, error) {
	user := pbu.User{
		Id:       "dsfhdsjfhl",
		Name:     "Nodirbek",
		Age:      17,
		Username: "rarebek",
		Email:    "nomonovn2@gmail.com",
	}
	return &user, nil
}

func (c *userServiceClient) UpdateUserById(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *pbu.UserId, opts ...grpc.CallOption) (*pbu.User, error) {
	user := pbu.User{
		Id:       "dsfhdsjfhl",
		Name:     "Nodirbek",
		Age:      17,
		Username: "rarebek",
		Email:    "nomonovn2@gmail.com",
	}
	return &user, nil
}

func (c *userServiceClient) ListUser(ctx context.Context, in *pbu.GetAllUserRequest, opts ...grpc.CallOption) (*pbu.GetAllUserResponse, error) {
	user := pbu.User{
		Id:       "dsfhdsjfhl",
		Name:     "Nodirbek",
		Age:      17,
		Username: "rarebek",
		Email:    "nomonovn2@gmail.com",
	}
	resp := pbu.GetAllUserResponse{
		Count: 1,
		Users: []*pbu.User{
			&user,
			&user,
			&user,
		},
	}
	return &resp, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *pbu.UserId, opts ...grpc.CallOption) (*pbu.Empty, error) {
	return &pbu.Empty{}, nil
}

func (c *userServiceClient) CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error) {
	return &pbu.CheckFieldResponse{
		Status: false,
	}, nil
}
