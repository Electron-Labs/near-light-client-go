package nearprimitive

import (
	"bytes"
	"crypto/ed25519"
	borsh "github.com/near/borsh-go"
	"reflect"
	"testing"
)

func TestCryptoHashBytes(t *testing.T) {
	c := CryptoHash{}
	byteArray := []byte("hello world\n")
	c.HashBytes(byteArray)
	expectedHash := []byte{169, 72, 144, 79, 47, 15, 71, 155, 143, 129,
		151, 105, 75, 48, 24, 75, 13, 46, 209, 193, 205, 42, 30, 192,
		251, 133, 210, 153, 161, 146, 164, 71}
	if !bytes.Equal(expectedHash, c.AsBytes()) {
		t.Errorf("Did not match %x", c.AsBytes())
	}
}

func TestCryptoHashBorshBytes(t *testing.T) {
	c := CryptoHash{}
	byteArray := []byte("hello world\n")

	serializedData, err := borsh.Serialize(byteArray)
	if err != nil {
		t.Errorf("Error while serializing: %s", err)
	}

	err = c.HashBorsh(serializedData)
	if err != nil {
		t.Error(err)
	}
	expectedHash := []byte{169, 72, 144, 79, 47, 15, 71, 155, 143, 129,
		151, 105, 75, 48, 24, 75, 13, 46, 209, 193, 205, 42, 30, 192,
		251, 133, 210, 153, 161, 146, 164, 71}
	if !bytes.Equal(expectedHash, c.AsBytes()) {
		t.Errorf("Did not match %x", c.AsBytes())
	}
}

func TestSignatureVerification(t *testing.T) {
	pub_key, priv_key, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Errorf("Failed to generate key pair: %s", err)
	}

	msg := []byte("hello world\n")

	signature := ed25519.Sign(priv_key, msg)

	s := &Signature{}
	err = s.TryFromRaw(signature)
	if err != nil {
		t.Errorf("Failed to generate a signature: %s", err)
	}

	p := &PublicKey{}
	err = p.TryFromRaw(pub_key)
	if err != nil {
		t.Errorf("Failed to generate public key: %s", err)
	}

	if !s.Verify(msg, p) {
		t.Errorf("Failed to verify the signature")
	}
}

func TestBlockHeaderInnerLiteViewFinalSerde(t *testing.T) {
	c := &CryptoHash{}
	data := []byte("hello world\n")
	c.HashBytes(data)

	bf := &BlockHeaderInnerLiteViewFinal{
		Height:          31,
		EpochId:         *c,
		NextEpochId:     *c,
		PrevStateRoot:   *c,
		OutcomeRoot:     *c,
		Timestamp:       4,
		NextBpHash:      *c,
		BlockMerkleRoot: *c,
	}

	der_bf := &BlockHeaderInnerLiteViewFinal{}

	ser_bf, err := bf.serialize()
	if err != nil {
		t.Errorf("Failed to serialize bf: %s", err)
	}

	err = der_bf.deserialize(ser_bf)
	if err != nil {
		t.Errorf("Failed to deserialize bf: %s", err)
	}

	if !reflect.DeepEqual(bf, der_bf) {
		t.Errorf("bf: %v\nder_bf: %v", bf, der_bf)
	}
}
