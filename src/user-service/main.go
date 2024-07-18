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
