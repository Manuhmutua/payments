package main

import (
	"payments/handler"
	pb "payments/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("payments"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPaymentsHandler(srv.Server(), new(handler.Payments))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
