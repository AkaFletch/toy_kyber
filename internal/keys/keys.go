package keys

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/AkaFletch/toy_kyber/v2/internal/kyber"
)

var bigintTwo = big.NewInt(2)
var bigintOne = big.NewInt(1)
var bigintZero = big.NewInt(0)

type polynomial struct {
	Components []polynomialComponent
}

type polynomialComponent struct {
	Exponent    *big.Int
	Coefficient *big.Int
}

func genPolynomial(terms, maxCoefficient *big.Int) (*polynomial, error) {
	p := &polynomial{}
	for len(p.Components) == 0 {
		for i := big.NewInt(0); i.Cmp(terms) < 0; i.Add(i, bigintOne) {
			// TODO Should use the modulus here instead of having a max
			coefficient, err := rand.Int(rand.Reader, maxCoefficient)
			if err != nil {
				return nil, err
			}
			negate, err := rand.Int(rand.Reader, bigintTwo)
			if err != nil {
				return nil, err
			}
			if coefficient.Cmp(bigintZero) == 0 {
				continue
			}
			if negate.Cmp(bigintOne) == 0 {
				coefficient.Neg(coefficient)
			}
			exponent := new(big.Int).Set(i)
			comp := polynomialComponent{
				Exponent:    exponent,
				Coefficient: coefficient,
			}
			p.Components = append(p.Components, comp)
		}
	}
	return p, nil
}

func GenerateKeyPair(publicPath, privatePath string, scheme kyber.KyberScheme) {
	keys, _ := generatePrivateKey(kyber.GetConfig(kyber.BABYKYBER))
	fmt.Println("Private key")
	for _, key := range keys {
		for _, term := range key.Components {
			fmt.Printf("%dx^%d ", term.Coefficient, term.Exponent)
		}
		fmt.Printf("\n")
	}
}

func generatePrivateKey(config kyber.KyberConfig) ([]*polynomial, error) {
	var polynomials []*polynomial
	for i := big.NewInt(0); i.Cmp(&config.K) < 0; i.Add(i, bigintOne) {
		p, err := genPolynomial(&config.N1, &config.N2)
		if err != nil {
			return nil, err
		}
		polynomials = append(polynomials, p)
	}
	return polynomials, nil
}
