// Binary to generate kyber public/private keys
package main

import (
	"flag"
	"fmt"

	"github.com/AkaFletch/toy_kyber/v2/internal/keys"
	"github.com/AkaFletch/toy_kyber/v2/internal/kyber"
)

func main() {
	privatePath := flag.String("private", ".kyber_id", "The filepath to write the private key to")
	publicPath := flag.String("public", ".kyber_id.pub", "The filepath to write the public key to")
	flag.Parse()
	fmt.Println("Generating Kyber keypair")
	keys.GenerateKeyPair(*publicPath, *privatePath, kyber.BABYKYBER)
}
