package main

//Ideally, I would use testify here, but I am sticking to the stdlib for simplicity sake

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	//pre-determined hash for the current object
	expected := "32e8cc21ff0e4334e31b2795b8d21c7418b6153dc0ded2f7cf65ae9ef3037c9d"
	input := Block{
		Index:     0,
		Timestamp: "2018-01-30 16:42:05.455708 -0500 EST",
		BPM:       0,
		PrevHash:  "",
		Hash:      "",
	}
	result := calculateHash(input)
	if expected != result {
		t.Error("Hash does not equal expected hash")
	}
}

func TestGenerateBlock(t *testing.T) {
	oldBlock := Block{
		Index:     0,
		Timestamp: "2018-01-30 16:42:05.455708 -0500 EST",
		BPM:       0,
		PrevHash:  "",
		Hash:      "32e8cc21ff0e4334e31b2795b8d21c7418b6153dc0ded2f7cf65ae9ef3037c9d",
	}
	newBlock, err := generateBlock(oldBlock, 123)

	if err != nil {
		t.Errorf("\n%+v\n", err)
	}
	if newBlock.Index != 1 {
		t.Error("Index is not what is expected")
	}
	if newBlock.BPM != 123 {
		t.Error("BPM is not what is expected")
	}
	if newBlock.PrevHash != oldBlock.Hash {
		t.Error("PrevHash is not what is expected")
	}
	if newBlock.Hash == "" {
		t.Error("Hash is blank is not what is expected")
	}
}

func TestIsBlockValid(t *testing.T) {
	oldBlock := Block{
		Index:     0,
		Timestamp: "2018-01-30 16:42:05.455708 -0500 EST",
		BPM:       0,
		PrevHash:  "",
		Hash:      "32e8cc21ff0e4334e31b2795b8d21c7418b6153dc0ded2f7cf65ae9ef3037c9d",
	}
	newBlock, _ := generateBlock(oldBlock, 123)

	valid := isBlockValid(newBlock, oldBlock)
	if !valid {
		t.Error("Expected block to be valid")
	}

	newBlock.Index = 4
	valid = isBlockValid(newBlock, oldBlock)
	if valid {
		t.Error("Expected block to not be valid")
	}

	newBlock.Index = 1
	newBlock.PrevHash = "bad_hash"

	valid = isBlockValid(newBlock, oldBlock)
	if valid {
		t.Error("Expected block to not be valid")
	}

	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = "bad_hash"

	valid = isBlockValid(newBlock, oldBlock)
	if valid {
		t.Error("Expected block to not be valid")
	}
}

func TestReplaceChain(t *testing.T) {
	oldBlock := Block{
		Index:     0,
		Timestamp: "2018-01-30 16:42:05.455708 -0500 EST",
		BPM:       0,
		PrevHash:  "",
		Hash:      "32e8cc21ff0e4334e31b2795b8d21c7418b6153dc0ded2f7cf65ae9ef3037c9d",
	}
	newBlock, _ := generateBlock(oldBlock, 123)
	newBlocks := []Block{
		oldBlock,
		newBlock,
	}
	replaceChain(newBlocks)
	if len(Blockchain) != 2 {
		t.Error("Expected Blockchain length to be 2")
	}
}
