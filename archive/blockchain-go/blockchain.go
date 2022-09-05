// 2018-03-02 14:54
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	dbFile              = "/tmp/blockchain.db"
	blocksBucket        = "blocks"
	genesisCoinbaseData = "genesis"
)

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) MineBlock(transactions []*Transaction) {
	fmt.Println(">>>minging")
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})

	newBlock := NewBlock(transactions, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		return err
	})
	_ = err
	fmt.Println("<<<done, nonce:", newBlock.Nonce)
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func NewBlockChain(address string) *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{tip, db}
	return &bc
}

func CreateBlockChain(address string) *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock(NewCoinBaseTX(address, genesisCoinbaseData))
			b, err = tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	return &BlockChain{tip, db}
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.tip, bc.db}
}

func (i *BlockChainIterator) Next() *Block {
	if len(i.currentHash) == 0 {
		return nil
	}
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	_ = err
	i.currentHash = block.PrevBlockHash
	return block
}

func (bc *BlockChain) PrintChain() {
	it := bc.Iterator()
	for {
		block := it.Next()
		if block == nil {
			break
		}
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Printf("valid: %s\n", strconv.FormatBool(NewProofOfWork(block).Validate()))
		fmt.Printf("transaction: %x\n", block.Transactions[0].ID)
		fmt.Println("  out:")
		for _, out := range block.Transactions[0].Vout {
			fmt.Printf("    value: %d, address: %s\n", out.Value, PKH2Address(out.PubKeyHash))
		}
		fmt.Println("  in:")
		for _, input := range block.Transactions[0].Vin {
			fmt.Printf("    transaction: %x, index: %d\n", input.Txid, input.Vout)
		}
		fmt.Println("====================")
	}
}

func (bc *BlockChain) FindUnspendTransactions(pubKeyHash []byte) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		if block == nil {
			break
		}
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				if out.IsLockedWithKey(pubKeyHash) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}
			if !tx.IsCoinbase() {
				for _, input := range tx.Vin {
					if input.UsesKey(pubKeyHash) {
						inTxId := hex.EncodeToString(input.Txid)
						spentTXOs[inTxId] = append(spentTXOs[inTxId], input.Vout)
					}
				}
			}
		}
	}
	return unspentTXs
}

func (bc *BlockChain) FindUTXO(pubKeyHash []byte) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspendTransactions(pubKeyHash)
	for _, tx := range unspentTransactions {
		for _, output := range tx.Vout {
			if output.IsLockedWithKey(pubKeyHash) {
				UTXOs = append(UTXOs, output)
			}
		}
	}
	return UTXOs
}

func (bc *BlockChain) GetBalance(pubKeyHash []byte) int {
	UTXOs := bc.FindUTXO(pubKeyHash)
	ret := 0
	for _, output := range UTXOs {
		ret += output.Value
	}
	return ret
}

func (bc *BlockChain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspendTransactions(pubKeyHash)
	accumulated := 0
work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Vout {
			if out.IsLockedWithKey(pubKeyHash) {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
				if accumulated >= amount {
					break work
				}
			}
		}
	}
	return accumulated, unspentOutputs
}
