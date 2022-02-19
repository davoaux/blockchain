package bc

import (
	"fmt"
	"strings"
	"time"
)

type Chain []*Block

func (ch *Chain) AddBlock(b *Block) error {
	if !b.isValid() {
		return fmt.Errorf("The block #%d is not valid", b.Header.Height)
	}
	if b.Header.PrevHash != ch.Tip().Hash {
		return fmt.Errorf("The block #%d doesn't point to the previous block", b.Header.Height)
	}
	if len(b.Transactions) < 1 {
		return fmt.Errorf("The block #%d has to have at least one transaction", b.Header.Height)
	}
	for _, tx := range b.Transactions {
		if !tx.isValid() {
			return fmt.Errorf("The transaction %s in the block #%d is not valid", tx.Txid, b.Header.Height)
		}
	}
	*ch = append(*ch, b)
	return nil
}

func (ch Chain) IsValid() bool {
	for i := len(ch) - 1; i != 0; i-- {
		if !ch[i].isValid() {
			return false
		}
	}
	return true
}

// Find block by its height
func (ch Chain) Get(height uint64) (*Block, error) {
	for _, b := range ch {
		if b.Header.Height == height {
			return b, nil
		}
	}
	return nil, fmt.Errorf("Block #%d not found", height)
}

// Return the last block or nil if empty
func (ch Chain) Tip() *Block {
	if len(ch) > 0 {
		return ch[len(ch)-1]
	} else {
		return nil
	}
}

func (ch Chain) String() string {
	builder := strings.Builder{}
	for _, b := range ch {
		builder.WriteString(fmt.Sprintf(
			"%d: {Nonce: %8d, Timestamp: %s, hash: %s}\n",
			b.Header.Height,
			b.Header.Nonce,
			time.Unix(b.Header.Timestamp, 0),
			b.Hash,
		))
	}
	return builder.String()
}
