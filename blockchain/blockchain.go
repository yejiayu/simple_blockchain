package blockchain

import (
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/yejiayu/simple_blockchain/database"
)

const tipBlock = "tip"

type BlockChain struct {
	tip *Block
	db  database.Interface
}

func NewBlockChain() (*BlockChain, error) {
	db, err := database.New("temp/")
	if err != nil {
		return nil, err
	}

	var b *Block
	data, err := db.Get([]byte(tipBlock))
	if err != nil && err != errors.ErrNotFound {
		return nil, err
	}

	if err == errors.ErrNotFound {
		b = NewBlock([]byte("Genesis block"), []byte{})
	} else {
		if err := json.Unmarshal(data, &b); err != nil {
			return nil, err
		}
	}

	return &BlockChain{b, db}, nil
}

func (bc *BlockChain) AddBlockWithStr(data string) (*Block, error) {
	prevBlock := bc.tip
	block := NewBlock([]byte(data), prevBlock.Hash)

	if err := bc.saveBlock(block); err != nil {
		return nil, err
	}
	bc.tip = block
	return block, nil
}

func (bc *BlockChain) saveBlock(b *Block) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}

	if err := bc.db.Put([]byte(tipBlock), data); err != nil {
		return err
	}

	return bc.db.Put(b.Hash, data)
}

func (bc *BlockChain) GetByHash(hash []byte) (*Block, error) {
	data, err := bc.db.Get(hash)
	if err != nil {
		return nil, err
	}

	var b Block
	if err := json.Unmarshal(data, &b); err != nil {
		return nil, err
	}

	return &b, nil
}
