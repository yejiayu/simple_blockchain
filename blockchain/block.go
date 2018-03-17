package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func NewBlock(data, prevBlockHash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	b.SetHash()

	return b
}

type Block struct {
	Timestamp     int64  `json:"timestamp"`
	Hash          []byte `json:"hash"`
	PrevBlockHash []byte `json:"prevBlockHash"`
	Data          []byte `json:"data"`
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{timestamp, b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
