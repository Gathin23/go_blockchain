package main

import (
	"fmt"
	"log"
	"time"
	"strings"
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

func (b *Block) Print() {
	fmt.Printf("Timestamp        %d\n",b.timestamp)
	fmt.Printf("Nonce            %d\n",b.nonce)
	fmt.Printf("Previous_Hash    %s\n",b.previousHash)
	fmt.Printf("Transactions     %s\n",b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init Hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n",strings.Repeat("=",25),i,strings.Repeat("=",25))
		block.Print()
		fmt.Printf("%s\n",strings.Repeat("*",25))
	}
}


func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.CreateBlock(5, "hash1")
	blockchain.Print()
}
