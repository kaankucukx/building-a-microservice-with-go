package main

import (
	"context"
	"fmt"
	application "github.com/kaankucukx/building-a-microservice-with-go/app"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}


