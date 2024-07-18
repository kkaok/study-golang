SQLite3를 사용하여 사용자 정보를 관리하는 Go Kit 서비스로 변경하려면 SQLite3 데이터베이스를 초기화하고 CRUD 작업을 위한 메서드를 구현해야 합니다. 다음은 SQLite3를 사용하여 사용자 정보를 등록, 수정, 삭제, 조회 및 전체 조회하는 Go Kit API 예제입니다.

# 프로젝트 구조
프로젝트 구조는 다음과 같습니다:

```
user-service-db/
├── go.mod
├── go.sum
├── main.go
├── service/
│   ├── models.go
│   └── service.go
├── endpoint/
│   └── endpoint.go
└── transport/
    └── http.go
```

# SQLite3 초기화 및 서비스 구현
먼저, SQLite3 데이터베이스를 초기화하고, 이를 이용해 서비스 메서드를 구현합니다.

```
% go mod init user-service-db
go: creating new go.mod: module user-service-db
go: to add module requirements and sums:
        go mod tidy

% go get -u gorm.io/gorm
go: downloading gorm.io/gorm v1.25.11
go: downloading github.com/jinzhu/now v1.1.5
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading golang.org/x/text v0.14.0
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added golang.org/x/text v0.16.0
go: added gorm.io/gorm v1.25.11

% go get -u gorm.io/driver/sqlite
go: downloading gorm.io/driver/sqlite v1.5.6
go: added github.com/mattn/go-sqlite3 v1.14.22
go: added gorm.io/driver/sqlite v1.5.6

% go mod tidy
go: finding module for package github.com/gorilla/mux
go: finding module for package github.com/go-kit/kit/log
go: finding module for package github.com/go-kit/kit/endpoint
go: finding module for package github.com/go-kit/kit/transport/http
go: found github.com/go-kit/kit/log in github.com/go-kit/kit v0.13.0
go: found github.com/go-kit/kit/endpoint in github.com/go-kit/kit v0.13.0
go: found github.com/go-kit/kit/transport/http in github.com/go-kit/kit v0.13.0
go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.1
```

# main.go

```
package main

import (
    "log"
    "net/http"
    "user-service-db/endpoint"
    "user-service-db/service"
    "user-service-db/transport"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    if err := db.AutoMigrate(&service.User{}); err != nil {
        log.Fatal(err)
    }

    svc := service.NewUserService(db)

    endpoints := endpoint.Endpoints{
        CreateUserEndpoint:  endpoint.MakeCreateUserEndpoint(svc),
        GetUserEndpoint:     endpoint.MakeGetUserEndpoint(svc),
        UpdateUserEndpoint:  endpoint.MakeUpdateUserEndpoint(svc),
        DeleteUserEndpoint:  endpoint.MakeDeleteUserEndpoint(svc),
        GetAllUsersEndpoint: endpoint.MakeGetAllUsersEndpoint(svc),
    }

    r := transport.MakeHTTPHandler(endpoints)

    log.Println("Starting server on :9080...")
    log.Fatal(http.ListenAndServe(":9080", r))
}
```

# service/models.go
사용자 정보를 저장할 데이터 구조체를 정의합니다.

```
package service

type User struct {
    ID    string `gorm:"primaryKey" json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
service/service.go
UserService 인터페이스를 구현하는 구조체를 정의합니다.

package service

import (
    "context"

    "gorm.io/gorm"
)

type Service interface {
    CreateUser(ctx context.Context, user User) error
    GetUser(ctx context.Context, id string) (User, error)
    UpdateUser(ctx context.Context, user User) error
    DeleteUser(ctx context.Context, id string) error
    GetAllUsers(ctx context.Context) ([]User, error)
}

type userService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) Service {
    return &userService{db: db}
}

func (s *userService) CreateUser(ctx context.Context, user User) error {
    return s.db.Create(&user).Error
}

func (s *userService) GetUser(ctx context.Context, id string) (User, error) {
    var user User
    if err := s.db.First(&user, "id = ?", id).Error; err != nil {
        return User{}, err
    }
    return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user User) error {
    return s.db.Save(&user).Error
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
    return s.db.Delete(&User{}, "id = ?", id).Error
}

func (s *userService) GetAllUsers(ctx context.Context) ([]User, error) {
    var users []User
    if err := s.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
```

# endpoint/endpoint.go
서비스 메소드를 호출하는 엔드포인트를 정의합니다.

```
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
```

# transport/http.go
HTTP 핸들러를 정의하고, 엔드포인트와 연결합니다.

```
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
```

이 예제는 Go Kit을 사용하여 SQLite3 데이터베이스를 기반으로 사용자 정보를 등록, 조회, 수정, 삭제 및 전체 조회하는 기능을 제공합니다. 이를 통해 보다 영구적인 데이터 저장소를 사용할 수 있습니다.

# 실행

```
% go run main.go
2024/07/18 10:39:23 Starting server on :9080...
```

# 테스트

```
curl -X POST http://localhost:9080/users -d '{"id":"1","name":"Alice","email":"alice@example.com"}' -H "Content-Type: application/json"
curl -X POST http://localhost:9080/users -d '{"id":"2","name":"KaOk","email":"kaok@example.com"}' -H "Content-Type: application/json"
curl -X GET http://localhost:9080/users -H "Content-Type: application/json"
curl -X GET http://localhost:9080/users/1 -H "Content-Type: application/json"
```
