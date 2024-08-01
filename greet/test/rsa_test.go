package test

import (
	"greet/internal/crypto"
	"testing"
)

func TestRSA(t *testing.T) {
	crypto.RsaDecryptInit("../rsa_key/private_key.pem")
	t.Logf("%v", crypto.RSAPrivateKey)
	// err := crypto.RsaDecryptInit("./private.pem")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err := crypto.RsaDecryptInit("./private_key.pem")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// crypto.GenerateRsaKey(2048, "./")
}
