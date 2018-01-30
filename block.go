package main

/**
This is a fork of the blockchain article found at https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc

I made some changes that I think make things easier to understand. The article was concise and (rightly) focused on the blockchain aspects. My changes are
mostly for usability, maintainability, and clearer delineation of responsibilities. For example, we made sure all returns are the same shape and that we don't
directly return arrays (OWASP JSON Array security risk).
*/
import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

/*Block represents a single block in the chain*/
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	BPM       int    `json:"bpm"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
}

//Blockchain is the actual blockchain, held in memory
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
