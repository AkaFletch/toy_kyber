package kyber

import (
	"math/big"
)

type KyberScheme int

const (
	KYBER512 KyberScheme = iota
	BABYKYBER
)

type KyberConfig struct {
	// The maximum degree used by polynomials
	N big.Int
	// The polynomials per vector
	K big.Int
	// The modulus for numbers
	Q big.Int
	// How big coefficients of small polynomials can be
	N1 big.Int
	N2 big.Int
}

func GetConfig(scheme KyberScheme) KyberConfig {
	switch scheme {
	case BABYKYBER:
		return KyberConfig{
			N:  *big.NewInt(4),
			K:  *big.NewInt(2),
			Q:  *big.NewInt(17),
			N1: *big.NewInt(4),
			N2: *big.NewInt(5),
		}
	case KYBER512:
		return KyberConfig{
			N:  *big.NewInt(256),
			K:  *big.NewInt(2),
			Q:  *big.NewInt(3329),
			N1: *big.NewInt(4),
			N2: *big.NewInt(3),
		}
	}
	return KyberConfig{}
}
