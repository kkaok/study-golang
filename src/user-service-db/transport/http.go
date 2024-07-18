package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"user-service-db/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Request 디코딩 함수 정의
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

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteUserRequest{ID: vars["id"]}, nil
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// HTTP 핸들러 정의
func MakeHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.GetUserEndpoint,
		decodeGetUserRequest,
		encodeResponse,
	))
	r.Methods("PUT").Path("/users").Handler(httptransport.NewServer(
		endpoints.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUserEndpoint,
		decodeDeleteUserRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		endpoints.GetAllUsersEndpoint,
		decodeGetAllUsersRequest,
		encodeResponse,
	))
	return r
}
