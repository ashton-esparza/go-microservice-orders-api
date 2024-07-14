package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ashton-esparza/go-microservice-orders-api.git/application"
)

func main() {
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}