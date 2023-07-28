package main

import (
	"log"

	pb "github.com/lovelydett/tt-home/src/service_user/proto"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}