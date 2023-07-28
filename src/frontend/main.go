/**
 * @description: frontend service for all requests that render pages
 * @author: Xie Yuting
 * @date: 2023-07-28
 */

package main

import (
	"fmt"
	"net/http"

	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	mhttp "github.com/go-micro/plugins/v4/server/http"
	"github.com/gorilla/mux"
	handler "github.com/lovelydett/tt-home/frontend/handler"
	micro "go-micro.dev/v4"
)

func main() {
	homeHandler := &handler.HomeHandler{}

	server := micro.NewService(
		micro.Server(mhttp.NewServer()),
		micro.Client(mgrpc.NewClient()),
	)
	options := []micro.Option{
		micro.Name("tt-home-frontend"),
		micro.Version("1.0.0"),
		micro.Address(":8080"),
	}
	server.Init(options...)

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler.GetHomePage).Methods(http.MethodGet)
	// Getting static resources
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	micro.RegisterHandler(server.Server(), router)

	fmt.Println("Server starts at: 8080")
	server.Run()
}
