package mock

import (
	"crypto/ed25519"
	"crypto/sha256"

	"github.com/electron-labs/near-light-client-go/nearprimitive"
)

type MockHostFunction struct{}

func (m MockHostFunction) Sha256(data []byte) [32]byte {
	return sha256.Sum256(data[:])
}

func (m MockHostFunction) Verify(sig nearprimitive.Signature, data []byte, public_key nearprimitive.PublicKey) bool {
	return ed25519.Verify(public_key.GetEd25519PubKey(), data, sig.AsBytes())
}
