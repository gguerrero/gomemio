package main

import (
	"flag"
	"log"
	"os"

	"github.com/gguerrero/gomemio"
)

var address = flag.String("h", "0.0.0.0", "Address where the server will be running")
var port = flag.Int("p", 2020, "Port where the server will be running")

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	flag.Parse()
	s := gomemio.NewServer(*address, *port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
