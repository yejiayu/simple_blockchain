package blockchain

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlockWithStr("block1")
	bc.AddBlockWithStr("block2")
	bc.AddBlockWithStr("block3")

	blocks := bc.blocks
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		if i != 0 {
			prevBlock := blocks[i-1]
			if bytes.Compare(block.PrevBlockHash, prevBlock.Hash) != 0 {
				t.Fatal("The block prevHash don't match prevBlock hash")
			}
		}
		fmt.Printf("prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("data: %x\n", block.Data)
	}
}
