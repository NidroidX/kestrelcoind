package main

import (
	"fmt"
	"github.com/NidroidX/kestrelcoind/cmd/kestrelcoinwallet/libkestrelcoinwallet"
	"github.com/NidroidX/kestrelcoind/util"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	privateKey, publicKey, err := libkestrelcoinwallet.CreateKeyPair(false)
	if err != nil {
		panic(err)
	}

	addr, err := util.NewAddressPublicKey(publicKey, cfg.NetParams().Prefix)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Private key: %x\n", privateKey)
	fmt.Printf("Address: %s\n", addr)
}
