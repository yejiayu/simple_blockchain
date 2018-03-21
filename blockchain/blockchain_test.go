package blockchain

import (
	"bytes"
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc, _ := NewBlockChain()
	b1, _ := bc.AddBlockWithStr("block1")
	b2, _ := bc.AddBlockWithStr("block2")

	if bytes.Compare(b2.PrevBlockHash, b1.Hash) != 0 {
		t.Fatal("The block prevHash don't match prevBlock hash")
	}
}
