package identity

import (
	"crypto/rand"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func CreateIdentity() (crypto.PrivKey, error) {
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Error generating key pair: %v", err)
	}
	return priv, nil
}
