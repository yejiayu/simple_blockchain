package blockchain

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlcok := NewBlock([]byte("Genesis block"), []byte{})
	return &BlockChain{[]*Block{genesisBlcok}}
}

func (bc *BlockChain) AddBlockWithStr(data string) *Block {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock([]byte(data), prevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
	return block
}
