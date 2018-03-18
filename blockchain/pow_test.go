package blockchain

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	pods := NewPow(NewBlock([]byte("test"), []byte{}))
	nonce, hash := pods.Run()
	fmt.Println(nonce, hex.EncodeToString(hash))
}

func TestValidate(t *testing.T) {
	p := NewPow(NewBlock([]byte("test"), []byte{}))
	if !p.Validate() {
		t.Fatal("illegal block")
	}
}
