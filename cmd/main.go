package main

import (
	"log"

	"omnicloud.mx/tasks/cmd/provider"
)

func main() {
	container := provider.New()
	if err := container.Build(); err != nil {
		log.Fatalln("container cannot be built because: ", err)
	}
}
