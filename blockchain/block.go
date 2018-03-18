package blockchain

import (
	"time"
)

func NewBlock(data, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	p := NewPow(b)
	b.Nonce, b.Hash = p.Run()

	return b
}

type Block struct {
	Timestamp     int64  `json:"timestamp"`
	Hash          []byte `json:"hash"`
	PrevBlockHash []byte `json:"prevBlockHash"`
	Data          []byte `json:"data"`
	Nonce         uint64 `json:"nonce"`
}
