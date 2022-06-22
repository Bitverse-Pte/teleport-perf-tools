package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"time"
)

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func QueryHeight(c *jsonrpc.Client, startHeight, endHeight int) {
	for i := startHeight; i < endHeight; i++ {
		block, err := c.Eth().GetBlockByNumber(ethgo.BlockNumber(i), true)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("block height:%d, block time: %d, total transactions:%d\n",
			i, block.Timestamp, len(block.Transactions))
		fmt.Println(time.Unix(int64(block.Timestamp), 0))
	}
}
