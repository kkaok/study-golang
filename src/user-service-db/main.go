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
