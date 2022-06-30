package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"os"
	"os/exec"
	"strconv"
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

		result, err := exec.Command("/bin/bash", "-c",
			"curl -s 172.17.0.1:26657/commit?height="+strconv.Itoa(i)+
				" | jq \"{time: .result.signed_header.header.time}\"").CombinedOutput()
		fmt.Println(string(result))
	}
}
