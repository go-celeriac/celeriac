package main

import (
	"fmt"
	"log"

	"github.com/go-celeriac/celeriac"
	_ "github.com/go-celeriac/celeriac/drivers/amqp"
)

func main() {
	fmt.Println("Doing a thing")

	b, err := celeriac.NewBroker("amqp://user:pass@0.0.0.0:5672/")
	if err != nil {
		log.Fatal(err)
	}

	//defer b.Close()

	b.Enqueue("test", "my-queue")
}
