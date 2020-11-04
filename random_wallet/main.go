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

	privateKeyBytes, err := hex.DecodeString("8752669eb70022ef0064dbe61ce47ba84e7673b69faa5bff9a70dbe525ba903ea04d19cd0e00d4ec8f7a81028672ea4dc342f6330013a7542bbbcce92af32c5d")
	privateKey := ed25519.PrivateKey(privateKeyBytes)
	publicKey := privateKey.Public().(ed25519.PublicKey)
	hash = sha3.Sum256(publicKey)

	fmt.Println("Fee Wallet Address: ", base58.CheckEncode(hash[1:], hash[0]))

}
