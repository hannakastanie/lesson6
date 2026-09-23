package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// create a new logger
	l := log.New(os.Stdout, "",  log.LstdFlags)
	// create a new server
	s := server.Router(l)
	
	// start the server with the logger
	if err := s.Serv.ListenAndServe(); err != nil{
		l.Fatal(err)
	}
}
