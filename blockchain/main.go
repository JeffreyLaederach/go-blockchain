package main

import (
	"fmt"
)

func main() {
	newblockchain := NewBlockchain() // Initialize the blockchain

	// Create 2 blocks and add 2 sample transactions
	newblockchain.AddBlock("First transaction")
	newblockchain.AddBlock("Second transaction")

	// Print all the blocks and their contents
	for i, block := range newblockchain.Blocks { // Iterate on each block
		fmt.Printf("Block ID : %d \n", i)                                       // Print the block ID
		fmt.Printf("Timestamp : %d \n", block.Timestamp+int64(i))               // Print the timestamp of the block, to make them different, +i
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)                 // Print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash) // Print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData)                 // Print the transactions
	} // The blockchain will be printed
}
