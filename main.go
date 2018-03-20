package main

import (
	"log"

	"github.com/bscott/k8s_buffalo_example/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
