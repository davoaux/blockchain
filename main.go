package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/davoaux/blockchain/bc"
)

var (
	block   *bc.Block
	chain   bc.Chain
	mempool []bc.Transaction // Pending transactions
)

func init() {
	// Genesis block
	genesis := bc.NewBlock([]bc.Transaction{*bc.NewTransaction("0", "1", 1)}, "", time.Now().Unix(), 0)
	chain = bc.Chain{genesis}

	mempool = append(mempool,
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
		*bc.NewTransaction(rndPerson(), rndPerson(), rand.Float32()),
	)
}

func main() {
	prevBlock := chain[len(chain)-1]
	block = bc.NewBlock(mempool, prevBlock.Hash, time.Now().Unix(), prevBlock.Header.Height+1)

	fmt.Printf("Mining block #%d\n", block.Header.Height)

	start := time.Now()
	mined := block.MineBlock()
	elapsed := time.Since(start)

	fmt.Printf("(nonce=%d): %s\n", block.Header.Nonce, mined)
	fmt.Printf("Took: %s\n", elapsed)

	if err := chain.AddBlock(block); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(chain)
	}
}

func rndPerson() string {
	return fmt.Sprintf("Person %c", rand.Intn(90-65)+65)
}
