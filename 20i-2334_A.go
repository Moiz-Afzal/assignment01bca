//Name: Moiz Ali Afzal
//Roll Number: 20i-2334
// Assignment 01

//Github Link: "github.com/Moiz-Afzal/assignment01bca"
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        int
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}

// Blockchain represents the entire blockchain.
var Blockchain []Block

// CalculateHash generates the hash for a block based on its data.
func (b *Block) CalculateHash() {
	data := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, b.Data, b.PreviousHash)
	hashInBytes := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hashInBytes[:])
}

// CreateBlock creates a new block and adds it to the blockchain.
func CreateBlock(data string, previousBlock Block) Block {
	var newBlock Block

	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.PreviousHash = previousBlock.Hash

	newBlock.CalculateHash()

	return newBlock
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := &Blockchain[i]
		previousBlock := &Blockchain[i-1]

		currentBlock.CalculateHash() // Update the current block's hash
		if currentBlock.Hash != currentBlock.Hash {
			return false // Hash is not valid
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false // Previous hash is not valid
		}
	}
	return true
}

func main() {
	// Create the genesis block
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Data:         "Genesis Block",
		PreviousHash: "",
	}
	genesisBlock.CalculateHash()
	Blockchain = append(Blockchain, genesisBlock)

	// Create some sample blocks
	newData := []string{"Transaction 1", "Transaction 2", "Transaction 3"}
	previousBlock := Blockchain[0]

	for _, data := range newData {
		newBlock := CreateBlock(data, previousBlock)
		Blockchain = append(Blockchain, newBlock)
		previousBlock = newBlock
	}

	// Verify the integrity of the blockchain
	isValid := VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}

	// Print all blocks in the blockchain
	for _, block := range Blockchain {
		fmt.Printf("Block %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

