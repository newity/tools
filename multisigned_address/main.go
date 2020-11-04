package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/btcsuite/btcutil/base58"

	"golang.org/x/crypto/sha3"
)

var (
	flagPubKeys    = flag.String("p", defaultPubKeys, "Public Keys for multiaddress")
	defaultPubKeys = "a4d153a9d9f42d64f8619b50e483d5441ec8cf5f1ae4bf3115bc3beb6621e727,  2edbb3ea3ef2f4aa0e8158b0abe7ec1250f9a092799943a954268b5ed75ba8e4 ,973394785ebacc881ee6b38018ac42cd86e7c4a4197fc0402a6d7f608f386161"
)

func main() {
	keys := strings.Split(*flagPubKeys, ",")

	for i := range keys {
		keys[i] = strings.Trim(keys[i], " ")
	}

	binPubKeys := make([][]byte, len(keys))
	for i, key := range keys {
		byteKey, err := hex.DecodeString(key)
		if err != nil {
			panic(err)
		}
		binPubKeys[i] = byteKey
	}
	sort.Slice(binPubKeys, func(i, j int) bool {
		return bytes.Compare(binPubKeys[i], binPubKeys[j]) < 0
	})

	hashedAddr := sha3.Sum256(bytes.Join(binPubKeys, []byte("")))
	address := base58.CheckEncode(hashedAddr[1:], hashedAddr[0])

	fmt.Println("Multiaddress: ", address)
}
