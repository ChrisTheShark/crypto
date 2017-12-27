package domain

import (
	"testing"
	"math/big"
)

func TestNewBlockChain(t *testing.T) {
	chain := NewBlockChain()

	t.Logf("Block length before add: %v", len(chain.blocks))
	chain.AddBlock("This is mock data.")
	t.Logf("Block length after add: %v\n\n", len(chain.blocks))

	if len(chain.blocks) != 2 {
		t.Errorf("Unexpected block length: %v", len(chain.blocks))
	}

	t.Log(">>> Block Output <<<\n")

	for _, block := range chain.blocks {
		t.Log(">>> Start Block <<<")
		t.Logf("Previous Hash: %x\n", block.PrevBlockHash)
		t.Logf("Data: %s\n", block.Data)
		t.Logf("Hash: %x\n", block.Hash)
		t.Log(">>> End Block <<<\n")
	}

	if string(chain.blocks[0].Hash) != string(chain.blocks[1].PrevBlockHash) {
		t.Errorf("Chain NOT bound, hash not propagated: \n%x\n%x",
			chain.blocks[0].Hash, chain.blocks[1].PrevBlockHash)
	}
}

func TestNewProofOfWork(t *testing.T) {
	proof := NewProofOfWork(NewOriginBlock())
	t.Logf("Target value from ProofOfWork: %v", proof.target)
	if proof.target.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Target value: %x is not populated", proof.target)
	}
	if !proof.Validate() {
		t.Errorf("Unable to validate created proof with nonce: %v", proof.block.Nonce)
	} else {
		t.Logf("Validated proof with nonce: %v", proof.block.Nonce)
	}
}
