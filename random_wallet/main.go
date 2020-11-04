package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	//"flag"
	"fmt"
	//"time"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

/*
var (
	seedArg        = flag.Int64("s", defaultSeedArg, "Random seed arg")
	defaultSeedArg = time.Now().UnixNano()
)
*/

func main() {
	//rand.Seed(*seedArg)

	pKey, sKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	hash := sha3.Sum256(pKey)

	fmt.Println("sKey: ", hex.EncodeToString(sKey))
	fmt.Println("pKey: ", hex.EncodeToString(pKey))
	fmt.Println("Address: ", base58.CheckEncode(hash[1:], hash[0]))
}
