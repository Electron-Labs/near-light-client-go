// Copyright © 2022, Electron Labs

package light

import (
	"testing"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
	num "github.com/shabbyrobe/go-num"
)

func TestValidateTransaction(t *testing.T) {
	decoded_data := &nearprimitive.CryptoHash{}
	err := decoded_data.TryFromRaw(base58.Decode("8HoqDvJGYrSjaejXpv2PsK8c5NUvqhU3EcUFkgq18jx9"))
	if err != nil {
		t.Errorf("Failed to read TX hash: %s", err)
	}

	outcome_proof_proot := []nearprimitive.MerklePathItem{}

	path_item_hash := &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("B1Kx1mFhCpjkhon9iYJ5BMdmBT8drgesumGZoohWhAkL"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	outcome_proof_proot = append(outcome_proof_proot, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Right})

	path_item_hash = &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("3tTqGEkN2QHr1HQdctpdCoJ6eJeL6sSBw4m5aabgGWBT"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	outcome_proof_proot = append(outcome_proof_proot, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Right})

	path_item_hash = &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("FR6wWrpjkV31NHr6BvRjJmxmL4Y5qqmrLRHT42sidMv5"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	outcome_proof_proot = append(outcome_proof_proot, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Right})

	outcome_root_proof := []nearprimitive.MerklePathItem{}

	path_item_hash = &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("3hbd1r5BK33WsN6Qit7qJCjFeVZfDFBZL3TnJt2S2T4T"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	outcome_root_proof = append(outcome_root_proof, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})

	path_item_hash = &nearprimitive.CryptoHash{}
	err = path_item_hash.TryFromRaw(base58.Decode("4A9zZ1umpi36rXiuaKYJZgAjhUH9WoTrnSBXtA3wMdV2"))
	if err != nil {
		t.Errorf("Failed to read path item hash: %s", err)
	}
	outcome_root_proof = append(outcome_root_proof, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})

	serialized_status := []uint8{3, 114, 128, 19, 177, 40, 127, 16, 184, 156, 69, 215, 55, 142, 98, 142, 27, 111, 246, 232, 85, 207, 169, 209, 101, 242, 113, 144, 111, 227, 117, 100, 30}
	decoded_hash := base58.Decode("8hxkU4avDWFDCsZckig7oN2ypnYvLyb1qmZ3SA1t8iZK")

	receipt_id := &nearprimitive.CryptoHash{}
	err = receipt_id.TryFromRaw(decoded_hash)
	if err != nil {
		t.Errorf("Failed to read receipt_id: %s", err)
	}

	tokens_burnt, _, _ := num.U128FromString("242839501800800000000")

	execution_outcome := nearprimitive.ExecutionOutcomeView{
		Logs:        []string{},
		ReceiptIds:  []nearprimitive.CryptoHash{*receipt_id},
		GasBurnt:    2428395018008,
		TokensBurnt: tokens_burnt,
		ExecutorId:  "relay.aurora",
		Status:      serialized_status,
	}

	outcome_proof := nearprimitive.OutcomeProof{
		BlockHash: nearprimitive.CryptoHash{},
		Id:        *decoded_data,
		Proof:     outcome_proof_proot,
		Outcome:   execution_outcome,
	}

	expected_block_outcome_root := &nearprimitive.CryptoHash{}
	expected_block_outcome_root.TryFromRaw(base58.Decode("AZYywqmo6vXvhPdVyuotmoEDgNb2tQzh2A1kV5f4Mxmq"))

	err = ValidateTransaction(mock.MockHostFunction{}, outcome_proof, outcome_root_proof, *expected_block_outcome_root)
	if err != nil {
		t.Errorf("Failed to validate transaction: %s", err)
	}
}
