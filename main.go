package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"sso/handler"
	"sso/subscriber"

	example "sso/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.hxlb.srv.sso"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.hxlb.srv.sso", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.hxlb.srv.sso", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
