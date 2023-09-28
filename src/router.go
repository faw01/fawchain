package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// returns the blocks in the blockchain as a JSON response.
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(blockchainInstance.blocks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// creates a new block based on the provided data.
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]float64
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	basePrice, exists := requestData["Data"]
	if !exists {
		http.Error(w, "Data not provided", http.StatusBadRequest)
		return
	}

	blockchainInstance.AddBlock(basePrice)
	w.WriteHeader(http.StatusCreated)
}

// initializes the HTTP router.
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}
