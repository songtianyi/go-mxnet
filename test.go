package main

import (
	"fmt"
	"github.com/songtianyi/go-mxnet/go"
)

func main() {
	if err := mxnet.MXRandomSeed(1); err != nil {
		fmt.Println(err)
	}
}
