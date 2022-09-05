// 2018-03-02 14:15
package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Usage() {
	fmt.Println("usage: createblockchain -address address")
	fmt.Println("usage: createwallet -address address")
	fmt.Println("usage: getbalance -address address")
	fmt.Println("usage: send -from xxx -to xxx -amount xxx")
	fmt.Println("usage: createwallet")
	fmt.Println("usage: listaddress")
	fmt.Println("usage: printchain")
}

func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		cli.Usage()
		os.Exit(1)
	}
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	getBalanceAddress := getBalanceCmd.String("address", "", "The address of get balance for")

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendFrom := sendCmd.String("from", "", "send from")
	sendTo := sendCmd.String("to", "", "send to")
	sendAmount := sendCmd.Int("amount", 0, "send to")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	listAddressCmd := flag.NewFlagSet("listaddress", flag.ExitOnError)

	switch os.Args[1] {
	case "createblockchain":
		createBlockchainCmd.Parse(os.Args[2:])
	case "getbalance":
		getBalanceCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	case "send":
		sendCmd.Parse(os.Args[2:])
	case "createwallet":
		createWalletCmd.Parse(os.Args[2:])
	case "listaddress":
		listAddressCmd.Parse(os.Args[2:])
	default:
		cli.Usage()
		os.Exit(1)
	}
	if createBlockchainCmd.Parsed() {
		if len(*createBlockchainAddress) == 0 {
			cli.Usage()
			os.Exit(1)
		}
		CreateBlockChain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		bc := NewBlockChain("")
		bc.PrintChain()
	}

	if sendCmd.Parsed() {
		if len(*sendFrom) == 0 || len(*sendTo) == 0 {
			cli.Usage()
			os.Exit(1)
		}
		bc := NewBlockChain(*sendFrom)
		tx := NewUTXOTransaction(*sendFrom, *sendTo, *sendAmount, bc)
		bc.MineBlock([]*Transaction{tx})
		fmt.Println("done")
	}

	if getBalanceCmd.Parsed() {
		if len(*getBalanceAddress) == 0 {
			cli.Usage()
			os.Exit(1)
		}
		pubKeyHash := Base58Decode([]byte(*getBalanceAddress))
		pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

		balance := NewBlockChain(*getBalanceAddress).GetBalance(pubKeyHash)
		fmt.Printf("Balance of %s is %d\n", *getBalanceAddress, balance)
	}

	if createWalletCmd.Parsed() {
		wallets, _ := NewWallets()
		address := wallets.CreateWallet()
		wallets.SaveToFile()

		fmt.Printf("Your new address: %s\n", address)
	}

	if listAddressCmd.Parsed() {
		wallets, _ := NewWallets()
		fmt.Println(wallets.GetAddresses())
	}
}

func main() {
	cli := CLI{}
	cli.Run()
}
