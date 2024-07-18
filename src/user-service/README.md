Go Kit는 마이크로서비스 아키텍처를 구현하기 위한 라이브러리입니다. 이를 사용하여 사용자 정보를 등록, 수정, 삭제, 조회하는 API를 만드는 예제를 작성하겠습니다. 이 예제는 Go Kit의 핵심 개념인 서비스, 엔드포인트, 트랜스포트 레이어를 포함합니다.

# 프로젝트 구조
프로젝트 구조는 다음과 같습니다:

```
user-service/
├── go.mod
├── go.sum
├── main.go
├── service/
│   ├── user.go
│   └── user_service.go
├── endpoint/
│   └── user_endpoint.go
└── transport/
    └── http.go
```
# main.go

프로젝트의 진입점입니다. 서비스와 엔드포인트를 설정하고 HTTP 트랜스포트를 시작합니다.

```
package main

import (
    "log"
    "net/http"
    "os"

    "user-service/endpoint"
    "user-service/service"
    "user-service/transport"

    kitlog "github.com/go-kit/kit/log"
    "github.com/gorilla/mux"
)

func main() {
    logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))

    userService := service.NewUserService()
    userEndpoints := endpoint.MakeUserEndpoints(userService)

    r := mux.NewRouter()
    transport.MakeHTTPHandler(r, userEndpoints, logger)

    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":9080", nil))
}
```

# service/user.go
사용자 정보를 저장할 데이터 구조체를 정의합니다.

```
package service

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserService interface {
    CreateUser(User) (string, error)
    GetUser(string) (User, error)
    GetUsers() ([]User, error)
    UpdateUser(User) error
    DeleteUser(string) error
}
```

# service/user_service.go
UserService 인터페이스를 구현하는 구조체를 정의합니다.

```
package service

import (
    "errors"
    "sync"
)

var (
    ErrUserNotFound = errors.New("user not found")
)

type userService struct {
    mu    sync.Mutex
    users map[string]User
}

func NewUserService() UserService {
    return &userService{
        users: make(map[string]User),
    }
}

func (s *userService) CreateUser(user User) (string, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.users[user.ID] = user
    return user.ID, nil
}

func (s *userService) GetUser(id string) (User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    user, ok := s.users[id]
    if !ok {
        return User{}, ErrUserNotFound
    }
    return user, nil
}

func (s *userService) GetUsers() ([]User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    users := make([]User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users, nil
}

func (s *userService) UpdateUser(user User) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, ok := s.users[user.ID]; !ok {
        return ErrUserNotFound
    }
    s.users[user.ID] = user
    return nil
}

func (s *userService) DeleteUser(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, ok := s.users[id]; !ok {
        return ErrUserNotFound
    }
    delete(s.users, id)
    return nil
}
```

# endpoint/user_endpoint.go
서비스 메소드를 호출하는 엔드포인트를 정의합니다.

```
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
```

# transport/http.go
HTTP 핸들러를 정의하고, 엔드포인트와 연결합니다.

```
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
```

# decodeCreateUserRequest 함수 설명
이 함수는 http.Request를 디코딩하여 endpoint.CreateUserRequest 구조체로 변환하는 역할을 합니다. 이는 Go의 웹 서버에서 요청 본문을 읽고, 해당 데이터를 사용하여 사용자 요청을 처리하는 데 사용됩니다.

## 함수의 전체 구조

```
func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req endpoint.CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
        return nil, err
    }
    return req, nil
}
```

## 세부 설명
> 함수 시그니처
>> decodeCreateUserRequest 함수는 두 개의 매개변수를 받습니다:
>>> context.Context: 주로 시간 초과, 취소 신호 또는 기타 요청 스코프 데이터를 전달하는 데 사용됩니다.
>>> *http.Request: HTTP 요청을 나타내며, 클라이언트로부터 받은 요청에 대한 모든 정보를 포함합니다.
>> 반환값은 두 개:
>>> interface{}: 요청 본문을 디코딩한 후 생성된 CreateUserRequest 구조체를 반환합니다.
>>> error: 요청 본문을 디코딩하는 동안 발생할 수 있는 오류를 반환합니다.
> endpoint.CreateUserRequest 구조체 선언
>> req 변수는 endpoint.CreateUserRequest 타입의 빈 구조체로 선언됩니다. 이 구조체는 사용자 생성 요청에 대한 데이터를 담고 있습니다.
> JSON 디코딩
>> json.NewDecoder(r.Body).Decode(&req.User)는 요청 본문(r.Body)을 JSON으로 디코딩하여 req.User 구조체에 할당합니다.
>> r.Body는 HTTP 요청의 본문을 나타냅니다. 보통 JSON 형식으로 된 데이터를 포함합니다.
>> Decode 메서드는 JSON 데이터를 읽고 지정된 구조체(req.User)에 매핑합니다.
>> 디코딩 중에 오류가 발생하면, 해당 오류를 반환합니다(return nil, err).
> 성공적인 디코딩 후 반환
>> 디코딩이 성공하면 req를 반환하고 오류는 nil로 설정됩니다.

## 예제
예를 들어, 클라이언트가 다음과 같은 JSON 데이터를 포함하는 POST 요청을 보냈다고 가정해 봅시다:

```
{
    "id": "1",
    "name": "Alice",
    "email": "alice@example.com"
}
```
이 요청 본문은 다음과 같은 구조체로 매핑됩니다:

```
type CreateUserRequest struct {
    User User `json:"user"`
}

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

decodeCreateUserRequest 함수는 요청 본문을 읽고 JSON 데이터를 CreateUserRequest 구조체의 User 필드에 매핑합니다.

## 실행 방법
위의 모든 파일을 작성한 후, 프로젝트 디렉토리에서 go mod init user-service 명령어를 실행하여 Go 모듈을 초기화합니다. 그런 다음 go get 명령어를 사용하여 필요한 패키지를 설치합니다.

이제 main.go 파일을 실행하여 서버를 시작할 수 있습니다.


먼저 "user-service"라는 디렉토리를 만들고 위의 코드로 디렉토리와 파일을 만듭니다.

프로젝트를 초기화하고 명령을 실행하여 필요한 모듈을 가져옵니다.

```
go mod init user-service // 이렇게 하면 go.mod와, go.sum 파일이 생성되는 걸 확인 할 있습니다.
go mod tidy  // 이 명령을 통해 필요한 모듈을 자동으로 가져옵니다.
```

위의 명령으로도 설치되지 않은 모듈이 있는 경우 로그를 보고 수동으로 설치해야 합니다.

```
go get github.com/go-kit/kit/log@v0.13.0
go get github.com/gorilla/mux
```

## 실행

```
go run main.go
```

## 테스트

```
curl 명령어를 사용하여 API를 테스트할 수 있습니다. 예를 들어, 새로운 사용자를 생성하려면 다음과 같이 할 수 있습니다.

curl -X POST http://localhost:9080/users -d '{"id":"1","name":"Alice","email":"alice@example.com"}' -H "Content-Type: application/json"
curl -X POST http://localhost:9080/users -d '{"id":"2","name":"KaOk","email":"kaok@example.com"}' -H "Content-Type: application/json"
curl -X GET http://localhost:9080/users -H "Content-Type: application/json"
curl -X GET http://localhost:9080/users/1 -H "Content-Type: application/json"
```
다른 API 엔드포인트도 유사한 방법으로 테스트할 수 있습니다.



# 소스 중에 'sync.Mutex'라는 부분이 있는데 이건 뭘까?
sync.Mutex는 Go에서 상호 배제를 제공하는 동기화 원시 자료형입니다. 이를 사용하여 여러 고루틴이 동일한 자원에 동시에 접근할 때 발생할 수 있는 데이터 레이스를 방지할 수 있습니다. sync.Mutex를 이용하면 한 고루틴이 자원을 사용하고 있을 때 다른 고루틴이 해당 자원에 접근하지 못하게 할 수 있습니다.

## sync.Mutex 사용법
sync.Mutex는 두 가지 주요 메서드를 제공합니다:

> Lock(): Mutex를 잠급니다. 만약 다른 고루틴이 이미 잠근 상태라면, 이 메서드는 해당 고루틴이 Unlock할 때까지 블록됩니다.
> Unlock(): Mutex를 잠금 해제합니다.

## 예제
다음은 sync.Mutex를 사용하는 간단한 예제입니다.

```
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()   // 임계 구역을 잠금
    c.count++
    c.mu.Unlock() // 임계 구역을 잠금 해제
}

func (c *Counter) Value() int {
    c.mu.Lock()   // 임계 구역을 잠금
    defer c.mu.Unlock() // 임계 구역을 잠금 해제
    return c.count
}

func main() {
    c := Counter{}

    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final count:", c.Value())
}
```

## 설명
- Counter 구조체에는 sync.Mutex와 count 필드가 있습니다.
- Increment 메서드는 mu.Lock()을 호출하여 임계 구역을 잠근 후 count를 증가시키고, mu.Unlock()을 호출하여 잠금을 해제합니다.
- Value 메서드는 mu.Lock()을 호출하여 임계 구역을 잠근 후 현재 count 값을 반환하기 전에 mu.Unlock()을 호출합니다.
- main 함수는 1000개의 고루틴을 생성하여 Increment 메서드를 호출합니다.
- 모든 고루틴이 완료될 때까지 sync.WaitGroup을 사용하여 기다린 후 최종 count 값을 출력합니다.

## 주의 사항
- 반드시 Lock과 Unlock을 쌍으로 사용해야 합니다. 만약 Unlock이 호출되지 않으면 다른 고루틴이 임계 구역에 접근할 수 없게 되어 데드락이 발생할 수 있습니다.
- defer를 사용하여 Unlock을 호출하면 Unlock이 보장되므로, 코드가 더 안전해집니다.

이렇게 sync.Mutex를 사용하여 여러 고루틴 간의 자원 접근을 안전하게 제어할 수 있습니다.

