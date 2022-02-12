package main

import (
	"log"
	"os"

	"github.com/andygeiss/meridian59-build/cli/internal/maintenance"
)

func main() {

	cmd := os.Args[1]

	handler := maintenance.NewHandler()
	handler.Connect("127.0.0.1:59595")
	handler.Send(cmd)
	out := handler.Receive()
	if err := handler.Error(); err != nil {
		log.Fatal(err)
	}
	handler.Close()

	println(out)
}
