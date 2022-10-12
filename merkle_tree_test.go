// Copyright © 2022, Electron Labs

package light

import (
	"bytes"
	"testing"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
)

func TestComputeFromPath(t *testing.T) {
	path := []nearprimitive.MerklePathItem{}

	path_item_hash := &nearprimitive.CryptoHash{}
	err := path_item_hash.TryFromRaw(base58.Decode("3hbd1r5BK33WsN6Qit7qJCjFeVZfDFBZL3TnJt2S2T4T"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	path = append(path, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})

	path_item_hash = &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("4A9zZ1umpi36rXiuaKYJZgAjhUH9WoTrnSBXtA3wMdV2"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	path = append(path, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})

	item_hash := &nearprimitive.CryptoHash{}
	err = item_hash.TryFromRaw(base58.Decode("2gvBz5DDhPVuy7fSPAu8Xei8oc92W2JtVf4SQRjupoQF"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}

	expected_block_outcome_root := &nearprimitive.CryptoHash{}
	err = expected_block_outcome_root.TryFromRaw(base58.Decode("AZYywqmo6vXvhPdVyuotmoEDgNb2tQzh2A1kV5f4Mxmq"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}

	computed_block_outcome_root, err := Compute_root_from_path(mock.MockHostFunction{}, path, nearprimitive.MerkleHash(*item_hash))
	if err != nil {
		t.Errorf("Failed to compute outcome root: %s", err)
	}

	computed_block_outcome_root_hash := &nearprimitive.CryptoHash{}
	err = computed_block_outcome_root_hash.TryFromRaw(computed_block_outcome_root[:])
	if err != nil {
		t.Errorf("Failed to convert to CryptoHash")
	}

	if !bytes.Equal(computed_block_outcome_root_hash.AsBytes(), expected_block_outcome_root.AsBytes()) {
		t.Errorf("Failed to validate compute root path")
	}
}
