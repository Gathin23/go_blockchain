package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return b
}

//override function to create printing format for blocks
func (b *Block) Print() {
	fmt.Printf("Timestamp        %d\n",b.timestamp)
	fmt.Printf("Nonce            %d\n",b.nonce)
	fmt.Printf("Previous_Hash    %s\n",b.previousHash)
	fmt.Printf("Transactions     %s\n",b.transactions)
}


func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
