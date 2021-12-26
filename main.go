package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

// Variables
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

// Main
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		genesisBlock := Block{0, time.Now().String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	// log.Fatal(run())
}

// calculate hash by using index, timestamp, bpm, previous hash
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generate block
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

// validate block by using Index, PrevHash and double check calculateHash function
func isBlockValid(newBlock, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// replace chain if that is longer than blockchain
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
