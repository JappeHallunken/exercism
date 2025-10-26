package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	two := big.NewInt(2)
	max := new(big.Int).Sub(p, two)

	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return new(big.Int).Add(r, two)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBig := big.NewInt(g)
	return new(big.Int).Exp(gBig, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	return priv, PublicKey(priv, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
