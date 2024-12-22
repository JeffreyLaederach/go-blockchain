package main

// Create the "block" data structure
// A block contains this info:
type Block struct {
	Timestamp         int64  // The time when the block was created
	PreviousBlockHash []byte // The hash of the previous block
	MyBlockHash       []byte // The hash of the current block
	AllData           []byte // The data or transactions (body info)
}

// Prepare the blockchain data structure:
type Blockchain struct {
	Blocks []*Block // blockchain = series of blocks
}
