package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
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
