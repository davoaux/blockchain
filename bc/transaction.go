package bc

import (
	"fmt"

	"github.com/davoaux/blockchain/util"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float32
	Confirmed bool
	Txid      string
}

func NewTransaction(sender, receiver string, amount float32) *Transaction {
	t := &Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Confirmed: false,
	}
	t.Txid = util.Encode(fmt.Sprintf("%s%s%f", sender, receiver, amount))
	return t
}

// TODO impl
func (tx Transaction) isValid() bool {
	return true
}
