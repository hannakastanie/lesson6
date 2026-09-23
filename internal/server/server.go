package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type MyServer struct{
	Logg log.Logger
	Serv http.Server
}

func Router(l log.Logger) *MyServer{
	var m MyServer

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandlerRoot)
	mux.HandleFunc("/upload", handlers.HandlerUpload)

	s := http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: &l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	m.Logg = l
	m.Serv = s

	return &m
}