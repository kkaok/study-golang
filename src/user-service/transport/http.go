package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"user-service/endpoint"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(r *mux.Router, endpoints endpoint.UserEndpoints, logger log.Logger) {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}

	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.GetUserEndpoint,
		decodeGetUserRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		endpoints.GetUsersEndpoint,
		decodeGetUsersRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUserEndpoint,
		decodeDeleteUserRequest,
		encodeResponse,
		options...,
	))
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetUserRequest{ID: vars["id"]}, nil
}

func decodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateUserRequest
	vars := mux.Vars(r)
	req.User.ID = vars["id"]
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteUserRequest{ID: vars["id"]}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
