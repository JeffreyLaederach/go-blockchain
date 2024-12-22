package main

import (
	// Necessary libraries:
	"bytes"         // Convert data into bytes in order to be sent on the network
	"crypto/sha256" // Crypto library to hash the data
	"strconv"       // For conversions
	"time"          // The time for timestamp
)

// Generating a hash of the block
// Concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // Get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // Concatenate all the block data
	hash := sha256.Sum256(headers)                                                               // Hash the whole thing
	block.MyBlockHash = hash[:]                                                                  // Now set the hash of the block
}

// A function for new block generation that then returns that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // Block is received
	block.SetHash()                                                           // Block is hashed
	return block                                                              // Block is returned with all information in it
}

// Create the "genesis" block function that will return the first block. The genesis block is the first block on the chain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // The genesis block is made with some data in it
}
