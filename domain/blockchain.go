package domain

import (
	"math/big"
	"bytes"
	"fmt"
	"crypto/sha256"
	"log"
	"math"
)

const targetBits = 10

type Blockchain struct {
	blocks []*Block
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewOriginBlock()}}
}

func NewOriginBlock() *Block {
	return NewBlock("Origin Block", []byte{})
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	return &ProofOfWork{ block, target }
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			[]byte(fmt.Sprintf("%x", pow.block.Timestamp)),
			[]byte(fmt.Sprintf("%x", int64(targetBits))),
			[]byte(fmt.Sprintf("%x", int64(nonce))),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		log.Printf("%x\n", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}