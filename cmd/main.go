package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yejiayu/simple_blockchain/blockchain"
)

var cmd = &cobra.Command{
	Use: "simple blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := blockchain.NewBlockChain()
		if err != nil {
			panic(err)
		}

		fmt.Println("start block chain")
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
