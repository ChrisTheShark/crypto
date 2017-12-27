package domain

import "testing"

func TestBlock(t *testing.T) {
	block := NewBlock("This is mock data.", NewOriginBlock().Hash)
	t.Logf("Obtained hash: %x", block.Hash)
	if block.Hash == nil {
		t.Errorf("Calculated hash is empty: %x", block.Hash)
	}
}