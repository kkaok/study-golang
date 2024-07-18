package endpoint

import (
	"context"
	"user-service-db/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateUserRequest struct {
	User service.User
}

type CreateUserResponse struct {
	Err error `json:"error,omitempty"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"err,omitempty"`
}

type UpdateUserRequest struct {
	User service.User `json:"user"`
}

type UpdateUserResponse struct {
	Err error `json:"err,omitempty"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Err error `json:"err,omitempty"`
}

type GetAllUsersResponse struct {
	Users []service.User `json:"users"`
	Err   error          `json:"err,omitempty"`
}

type Endpoints struct {
	CreateUserEndpoint  endpoint.Endpoint
	GetUserEndpoint     endpoint.Endpoint
	UpdateUserEndpoint  endpoint.Endpoint
	DeleteUserEndpoint  endpoint.Endpoint
	GetAllUsersEndpoint endpoint.Endpoint
}

func MakeCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := svc.CreateUser(ctx, service.User{
			ID:    req.User.ID,
			Name:  req.User.Name,
			Email: req.User.Email,
		})
		return CreateUserResponse{Err: err}, err
	}
}

func MakeGetUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := svc.GetUser(ctx, req.ID)
		return GetUserResponse{
			User: service.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
			Err: err,
		}, err
	}
}

func MakeUpdateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := svc.UpdateUser(ctx, service.User{
			ID:    req.User.ID,
			Name:  req.User.Name,
			Email: req.User.Email,
		})
		return UpdateUserResponse{Err: err}, err
	}
}

func MakeDeleteUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := svc.DeleteUser(ctx, req.ID)
		return DeleteUserResponse{Err: err}, err
	}
}

func MakeGetAllUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetAllUsers(ctx)
		var result []service.User
		for _, u := range users {
			result = append(result, service.User{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			})
		}
		return GetAllUsersResponse{Users: result, Err: err}, err
	}
}
