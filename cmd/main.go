package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	counterHTTP "github.com/gnumi34/word-counter/pkg/counter/delivery/http"
	"github.com/gnumi34/word-counter/pkg/counter/repository/rest"
	"github.com/gnumi34/word-counter/pkg/counter/usecase"
	"github.com/gnumi34/word-counter/pkg/utils"
)

func main() {
	mux := http.DefaultServeMux

	configureRouter(mux)

	go func() {
		log.Println("Server started! Listening to 0.0.0.0:8082")
		http.ListenAndServe("0.0.0.0:8082", mux)
	}()

	// Wait for shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")
}

func configureRouter(mux *http.ServeMux) {
	defaultClient := utils.GetClient()

	counterAPIRepository := rest.NewRESTClient(defaultClient)
	counterUseCase := usecase.NewUseCase(counterAPIRepository)
	counterHandler := counterHTTP.NewHandler(counterUseCase)

	mux.HandleFunc("/", counterHandler.HelloWorld)
	mux.HandleFunc("/count-word", counterHandler.CountWordsFromText)
}
