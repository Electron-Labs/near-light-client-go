// Copyright © 2022, Electron Labs

package light

import (
	"testing"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
	num "github.com/shabbyrobe/go-num"
)

const (
	TRANSACTION_PROOF = `
{
        "jsonrpc": "2.0",
        "result": {
                "block_header_lite": {
                        "inner_lite": {
                                "block_merkle_root": "D5nnsEuJ2WA4Fua4QJWXa3LF2TGoAqhrW8fctFh7MW2s",
                                "epoch_id": "7e3Vkbngf36bphkBVX98LoRxpoqhvZJbL5Rgb3Yfccy8",
                                "height": 86697768,
                                "next_bp_hash": "Hib973UH8xTq4ReP2urd1bLEaHGmjwWeHCyfQV4ZbHAv",
                                "next_epoch_id": "7AEtEQErauvaagnmmDsxw9qnYqBVuTKjSW4P7DVwZ5z3",
                                "outcome_root": "AZYywqmo6vXvhPdVyuotmoEDgNb2tQzh2A1kV5f4Mxmq",
                                "prev_state_root": "6BWNcpk4chiEXWRWbWum5D4zutZ9pomfwwbmjanLp4sv",
                                "timestamp": 1649062589965425850,
                                "timestamp_nanosec": "1649062589965425850"
                        },
                        "inner_rest_hash": "DeSCLALKLSEX6pjKVoStCUq3ixkzK4v958TMkdPp1fJJ",
                        "prev_block_hash": "Ae7sLAjvHs3gkiU2vFt8Vdxs5RmVUwyxyCwbnqnTkckQ"
                },
                "block_proof": [
                        {
                                "direction": "Right",
                                "hash": "BNmeYcDcNoVXgXZyzcoyJiN5UiyLeZTvwSHYRpSfw9fF"
                        },
                        {
                                "direction": "Right",
                                "hash": "A7HaT2EGxrhJhDK2muP56b6j6c5JL1VAFPE45iB4cxsf"
                        },
                        {
                                "direction": "Left",
                                "hash": "AjhQk267UxRgxrTtLyjHrVoid7DPRN67aki8GJZttnu4"
                        },
                        {
                                "direction": "Left",
                                "hash": "4qyS6XAo8fNLYeGQJVN31D8ncr4TfmrvSe3cursw8oM7"
                        },
                        {
                                "direction": "Right",
                                "hash": "28y98e3vha3vHmkBhgREgxjLzjP7JzfVeu6H6yDHMh4V"
                        },
                        {
                                "direction": "Left",
                                "hash": "CJRqXDJy8L1oEGJDPxXgPuQhrFmLosoFQAf79Dyfrw3z"
                        },
                        {
                                "direction": "Left",
                                "hash": "CGaUbgtx9UFf7sZAe5fLdy1ggb5ZGg2oC3LmT2SgnCbz"
                        },
                        {
                                "direction": "Left",
                                "hash": "EjFednH4uWzcYNJzrfiBPbcDEvVTi7u7MEDFbcJfdPYf"
                        },
                        {
                                "direction": "Right",
                                "hash": "HAxQFR7SS2gkNUZ4nfSNefo3N1mxsmn3n7sMzhBxxLi"
                        },
                        {
                                "direction": "Left",
                                "hash": "KQa9Nzw7vPnciog75ZGNriVU7r4aAqKErE15mEBd3sS"
                        },
                        {
                                "direction": "Left",
                                "hash": "ByNUgeXrsQpeCNeNEqpe8ASw2bh2BfY7knpLaQe1NtXv"
                        },
                        {
                                "direction": "Left",
                                "hash": "ByrTiguozXfUaufYN8MuWAx7jL1dhZJ7bLzJjpCQjvND"
                        },
                        {
                                "direction": "Left",
                                "hash": "DvV6ak7n9wP1TQ1a97P81b81xJq1EdnERp8r3GFdP7wU"
                        },
                        {
                                "direction": "Left",
                                "hash": "Gga62BEfbomV8ZNz3DkPQEFf6UbEqMKngwNAp5zDDoki"
                        },
                        {
                                "direction": "Left",
                                "hash": "76U6DMh4J4VB5sfVVNRpSTeB4SEVt4HPqhtQi2izGZxt"
                        }
                ],
                "outcome_proof": {
                        "block_hash": "5aZZNiqUVbXXvRjjf1FB8sbXG3gpJeVCw1bYeREXzHk2",
                        "id": "8HoqDvJGYrSjaejXpv2PsK8c5NUvqhU3EcUFkgq18jx9",
                        "outcome": {
                                "executor_id": "relay.aurora",
                                "gas_burnt": 2428395018008,
                                "logs": [],
                                "metadata": {
                                        "gas_profile": null,
                                        "version": 1
                                },
                                "receipt_ids": [
                                        "8hxkU4avDWFDCsZckig7oN2ypnYvLyb1qmZ3SA1t8iZK"
                                ],
                                "status": {
                                        "SuccessReceiptId": "8hxkU4avDWFDCsZckig7oN2ypnYvLyb1qmZ3SA1t8iZK"
                                },
                                "tokens_burnt": "242839501800800000000"
                        },
                        "proof": [
                                {
                                        "direction": "Right",
                                        "hash": "B1Kx1mFhCpjkhon9iYJ5BMdmBT8drgesumGZoohWhAkL"
                                },
                                {
                                        "direction": "Right",
                                        "hash": "3tTqGEkN2QHr1HQdctpdCoJ6eJeL6sSBw4m5aabgGWBT"
                                },
                                {
                                        "direction": "Right",
                                        "hash": "FR6wWrpjkV31NHr6BvRjJmxmL4Y5qqmrLRHT42sidMv5"
                                }
                        ]
                },
                "outcome_root_proof": [
                        {
                                "direction": "Left",
                                "hash": "3hbd1r5BK33WsN6Qit7qJCjFeVZfDFBZL3TnJt2S2T4T"
                        },
                        {
                                "direction": "Left",
                                "hash": "4A9zZ1umpi36rXiuaKYJZgAjhUH9WoTrnSBXtA3wMdV2"
                        }
                ]
        },
        "id": "idontcare"
}
`
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

func TestRealValidateTransaction(t *testing.T) {
	outcome_proof, merkle_path, expected_block_outcome_root := GetOutcomeProof(TRANSACTION_PROOF)

	err := ValidateTransaction(mock.MockHostFunction{}, outcome_proof, merkle_path, expected_block_outcome_root)
	if err != nil {
		t.Errorf("Failed to validate transaction: %s", err)
	}
}
