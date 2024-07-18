package endpoint

import (
	"context"
	"user-service/service"

	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	CreateUserEndpoint endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
	GetUsersEndpoint   endpoint.Endpoint
	UpdateUserEndpoint endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
}

func MakeUserEndpoints(s service.UserService) UserEndpoints {
	return UserEndpoints{
		CreateUserEndpoint: MakeCreateUserEndpoint(s),
		GetUserEndpoint:    MakeGetUserEndpoint(s),
		GetUsersEndpoint:   MakeGetUsersEndpoint(s),
		UpdateUserEndpoint: MakeUpdateUserEndpoint(s),
		DeleteUserEndpoint: MakeDeleteUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(req.User)
		if err != nil {
			return CreateUserResponse{ID: "", Err: err.Error()}, nil
		}
		return CreateUserResponse{ID: id, Err: ""}, nil
	}
}

func MakeGetUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := s.GetUser(req.ID)
		if err != nil {
			return GetUserResponse{User: service.User{}, Err: err.Error()}, nil
		}
		return GetUserResponse{User: user, Err: ""}, nil
	}
}

func MakeGetUsersEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetUsers()
		if err != nil {
			return GetUsersResponse{Users: nil, Err: err}, err
		}
		return GetUsersResponse{Users: users, Err: nil}, nil
	}
}

func MakeUpdateUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := s.UpdateUser(req.User)
		if err != nil {
			return UpdateUserResponse{Err: err.Error()}, nil
		}
		return UpdateUserResponse{Err: ""}, nil
	}
}

func MakeDeleteUserEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := s.DeleteUser(req.ID)
		if err != nil {
			return DeleteUserResponse{Err: err.Error()}, nil
		}
		return DeleteUserResponse{Err: ""}, nil
	}
}

type CreateUserRequest struct {
	User service.User
}

type CreateUserResponse struct {
	ID  string `json:"id"`
	Err string `json:"error,omitempty"`
}

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User service.User `json:"user"`
	Err  string       `json:"error,omitempty"`
}

type GetUsersResponse struct {
	Users []service.User `json:"users"`
	Err   error          `json:"err,omitempty"`
}

type UpdateUserRequest struct {
	User service.User
}

type UpdateUserResponse struct {
	Err string `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	ID string
}

type DeleteUserResponse struct {
	Err string `json:"error,omitempty"`
}
