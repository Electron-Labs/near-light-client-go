package light

import (
	"fmt"

	"github.com/electron-labs/near-light-client-go/nearprimitive"
	borsh "github.com/near/borsh-go"
)

type CombineHash struct {
	Hash1 nearprimitive.MerkleHash
	Hash2 nearprimitive.MerkleHash
}

func (c CombineHash) Serialize() ([]byte, error) {
	data, err := borsh.Serialize(c)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func combine_hash(h nearprimitive.HostFunction, hash1 nearprimitive.MerkleHash, hash2 nearprimitive.MerkleHash) (nearprimitive.MerkleHash, error) {
	final_hash := nearprimitive.MerkleHash{}
	c := CombineHash{Hash1: hash1, Hash2: hash2}
	combined_hash, err := c.Serialize()
	if err != nil {
		return final_hash, fmt.Errorf("Failed to serialize combine hash: %s", err)
	}

	data := h.Sha256(combined_hash)
	final_hash = data

	return final_hash, nil
}

func compute_root_from_path(h nearprimitive.HostFunction, path []nearprimitive.MerklePathItem, item_hash nearprimitive.MerkleHash) (nearprimitive.MerkleHash, error) {
	res := item_hash
	var err error

	for _, item := range path {
		if item.Direction == nearprimitive.Left {
			res, err = combine_hash(h, item.Hash, res)
			if err != nil {
				return res, fmt.Errorf("Failed to combine hash %s", err)
			}
		} else if item.Direction == nearprimitive.Right {
			res, err = combine_hash(h, res, item.Hash)
			if err != nil {
				return res, fmt.Errorf("Failed to combine hash %s", err)
			}
		}
	}

	return res, nil
}
