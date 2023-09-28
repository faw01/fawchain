package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// the blockchain
var blockchainInstance *Blockchain

func main() {
	err := godotenv.Load() // load environment variables
	if err != nil {
		log.Fatal(err)
	}

	blockchainInstance = NewBlockchain() // initialize blockchain

	run() // start HTTP server
}

// run starts the HTTP server
func run() {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	if httpAddr == "" {
		httpAddr = "8080"
	}
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
