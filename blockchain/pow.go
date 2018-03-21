package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

const targetBits = 8

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewPow(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBits)

	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	b := pow.block
	timestampB := make([]byte, 64)
	binary.LittleEndian.PutUint64(timestampB, uint64(b.Timestamp))
	nonceB := make([]byte, 64)
	binary.LittleEndian.PutUint64(nonceB, nonce)

	data := bytes.Join([][]byte{
		b.PrevBlockHash,
		b.Data,
		timestampB,
		nonceB,
		pow.target.Bytes(),
	}, []byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}

func (pow *ProofOfWork) Run() (uint64, []byte) {
	var nonce uint64
	var hashInt big.Int
	var hash []byte

	for nonce < math.MaxUint64 {
		hash = pow.prepareData(nonce)
		hashInt.SetBytes(hash)

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}

	return nonce, hash
}

func (pow *ProofOfWork) Validate() bool {
	hash := pow.prepareData(pow.block.Nonce)
	var hashInt big.Int
	hashInt.SetBytes(hash)

	return hashInt.Cmp(pow.target) == -1
}
