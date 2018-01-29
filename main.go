package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/*Block represents a single block in the chain*/
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

//Blockcahin is the actual blockchain, held in memory
var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	currentHash := sha256.New()
	currentHash.Write([]byte(record))
	hash := currentHash.Sum(nil)
	return hex.EncodeToString(hash)
}

//Generate a new block for the chain
func generateBlock(oldBlock Block, bpm int) (Block, error) {
	t := time.Now()
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: t.String(),
		BPM:       bpm,
		PrevHash:  oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock, nil
}

//checks is a block is valid. Todo: add an error message
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	calculated := calculateHash(newBlock)
	if calculated != newBlock.Hash {
		return false
	}
	return true
}

//if a suggested blockchain is longer than what we have, we need to update our chain
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

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

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	return muxRouter
}

func main() {
	fmt.Println("Not implemented yet")
}
