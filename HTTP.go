package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

//HTTPReturn represents a standardized return object
type HTTPReturn struct {
	Data interface{} `json:"data"`
}

//BlockPost is the post body for a new block in the chain
type BlockPost struct {
	BPM int `json:"bpm"`
}

//sendResponse is a single exit method for HTTP handlers that formats a response to send to the client
func sendResponse(w http.ResponseWriter, code int, payload interface{}) {
	var response HTTPReturn
	response.Data = payload
	toSend, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(toSend)
}

//sets up the mux and routes
func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")

	log.Printf("\nListening on %s", httpAddr)

	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	return err
}

//makeMuxRouter sets up the routes we will be using
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

/* Our HTTP handlers */

//handleGetBlockchain handles getting the current Blockchain
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, 200, Blockchain)
}

//handleWriteBlock takes in a POSTed block and adds it to the chain if successful
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	m := BlockPost{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		sendResponse(w, http.StatusBadRequest, "Could not decode BPM")
		return
	}
	defer r.Body.Close()

	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, m)
		return
	}

	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
	}
	sendResponse(w, http.StatusCreated, newBlock)
	return
}
