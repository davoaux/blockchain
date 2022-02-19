package bc

import (
	"fmt"
	"os"
	"strings"

	"github.com/davoaux/blockchain/util"
)

type Block struct {
	Header       BlockHeader
	Hash         string
	Transactions []Transaction
}

type BlockHeader struct {
	PrevHash   string
	MerkleRoot string
	Timestamp  int64
	Nonce      uint64
	Height     uint64
}

func NewBlock(transactions []Transaction, prevHash string, timestamp int64, height uint64) *Block {
	b := &Block{
		Transactions: transactions,
		Header: BlockHeader{
			PrevHash:  prevHash,
			Timestamp: timestamp,
			Height:    height,
		},
	}
	b.Header.MerkleRoot = b.calculateMerkleRoot()
	b.Hash = b.calculateBlockHash()
	return b
}

func (b Block) calculateBlockHash() string {
	if b.Header.MerkleRoot == "" {
		fmt.Fprintf(os.Stderr, "calculateBlockHash: merkle root is missing from block #%d", b.Header.Height)
		os.Exit(1)
	}
	return util.Encode(fmt.Sprintf("%s%d%d%s", b.Header.PrevHash, b.Header.Timestamp, b.Header.Nonce, b.Header.MerkleRoot))
}

func (b Block) calculateMerkleRoot() string {
	size := len(b.Transactions)
	// Hashes of the block's transactions (last one duplicated if the amount of transactions is odd)
	htx := make([]string, size)
	for i, tx := range b.Transactions {
		htx[i] = tx.Txid
	}
	if size%2 != 0 {
		htx = append(htx, b.Transactions[size-1].Txid)
	}
	return b.getMerkleRoot(htx)
}

func (b Block) getMerkleRoot(arr []string) string {
	if len(arr) <= 1 {
		return arr[0]
	}
	if len(arr)%2 != 0 {
		arr = append(arr, arr[len(arr)-1])
	}
	for i := 1; i <= len(arr); i++ {
		arr[i-1] = util.Encode(arr[i-1] + arr[i])
		arr = append(arr[:i], arr[i+1:]...)
	}
	return b.getMerkleRoot(arr)
}

func (b *Block) MineBlock() string {
	prefix := strings.Repeat("0", 5)
	for b.Hash[:len(prefix)] != prefix {
		b.Header.Nonce++
		b.Hash = b.calculateBlockHash()
	}
	return b.Hash
}

func (b Block) isValid() bool {
	return b.Header.Height == 0 || b.Hash < b.Header.PrevHash
}
