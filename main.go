package main

import (
	"crypto/sha256"

	"fmt"
)

// the block details consists of the following
type BlockDetails struct {
	transaction string //the transactiob consists of the string as well as the nonce
	hash        string //hash of that block
	prevhash    string //hash of the previous block
}

type blockchain struct {
	gblock BlockDetails   //This is the genisis block which is the first block of that block chain
	chain  []BlockDetails //This is an array of the blocks which is the "the blockchain"
}

func createbc() blockchain {
	gblock := BlockDetails{
		transaction: "Nothing", //created the 'block 0' with transaction string nothing
	}
	gblock.hash = gblock.CalHash() //Hash of the block is calculated
	return blockchain{
		gblock,
		[]BlockDetails{gblock}, //gblock is the first block or index of the blockchain array
	}
}

func (bc BlockDetails) CalHash() string {
	data := bc.transaction + bc.prevhash //Data is the transaction and previous hash is the hash calculated
	hash := sha256.Sum256([]byte(data))  //The hash of the data is stored in the hash which is an element of the struct. This is the hash of the current block
	return fmt.Sprintf("%x", hash)

}

func (bc *blockchain) newblock(trans string, x int) {

	data := fmt.Sprint(trans, x)
	prevblock := bc.chain[len(bc.chain)-1] //previous block is minus 1 from the current block
	newbc := BlockDetails{
		transaction: data, //The data and the previous hash is added into the new block
		prevhash:    prevblock.CalHash(),
	}
	bc.chain = append(bc.chain, newbc) //array is appended as new blocks are added

	bc.chain[(len(bc.chain) - 1)].hash = bc.chain[(len(bc.chain) - 1)].CalHash() //hash of this new block is calculated

}

func (bc *blockchain) listblock() string {

	for i := 0; i < len(bc.chain); i++ {
		if i == 0 {
			//this is block 0 which is the genesis block

			fmt.Println("======================================================")
			fmt.Println("\n")
			fmt.Println("Genesis Block is block number 0  and data is", bc.chain[i])
			fmt.Println("\n")

		}
		if i > 0 {

			fmt.Println("======================================================")
			fmt.Println("BLOCK ", i)
			fmt.Println("\n")
			fmt.Println("Transaction and nonce of ", bc.chain[i].transaction)
			fmt.Println("Hash of previous block is", bc.chain[i].prevhash)
			fmt.Println("Hash of current block is", bc.chain[i].hash)
			fmt.Println("\n")

		}

	}
	return ""

}
func (bc *blockchain) verify() bool {

	for i := 1; i < len(bc.chain); i++ {

		block := bc.chain[i]  //current block
		prev := bc.chain[i-1] // prev block
		//if the current hash and the previous hash are not different return true
		if block.hash != block.CalHash() || block.prevhash != prev.hash {
			return false
		}
	}
	return true
}
func (bc *blockchain) modify(index int, mtransaction string) {

	//modify the block and then calculate the hash again.
	bc.chain[index].transaction = mtransaction
	bc.chain[index].hash = bc.chain[index].CalHash()

}

func main() {

	bc := createbc()
	bc.newblock("Bob to Alice", 6)
	bc.newblock("Ali to Zoha", 10)
	//bc.newblock("amna to ali", 5)
	fmt.Println(bc.verify())
	bc.modify(1, "ali amna")
	fmt.Println(bc.verify())

	bc.listblock()

}
