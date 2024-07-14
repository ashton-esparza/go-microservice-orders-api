package main

import (
	"context"
	"fmt"

	"github.com/ashton-esparza/go-microservice-orders-api.git/application"
)

func main() {
	app := application.NewApp()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}