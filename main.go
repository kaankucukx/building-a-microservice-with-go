package main

import (
	"context"
	"fmt"
	application "github.com/kaankucukx/building-a-microservice-with-go/app"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}


