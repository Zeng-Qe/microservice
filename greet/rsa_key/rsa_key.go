package main

import (
	"fmt"
	"greet/internal/crypto"
)

func main() {
	// err := crypto.RsaDecryptInit("./private.pem")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := crypto.RsaDecryptInit("./private_key.pem")
	if err != nil {
		fmt.Println(err)
	}
	// crypto.GenerateRsaKey(2048, "./")
}
