package main

// The method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // The previous block is needed, so retrieve it
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)        // Create a new block containing the data and the hash of the previous block
	blockchain.Blocks = append(blockchain.Blocks, newBlock)      // Add that block to the chain to create a chain of blocks
}

// This is the function that returns the whole blockchain and initializes it with the "genesis" block (first mined block)
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}} // The genesis block is added first to the chain
}
