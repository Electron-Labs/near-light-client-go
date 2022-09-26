// Copyright © 2022, Electron Labs

package light

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/electron-labs/near-light-client-go/nearprimitive"

	base58 "github.com/btcsuite/btcutil/base58"
	num "github.com/shabbyrobe/go-num"
)

type NearInnerLightView struct {
	BlockMerkleRoot  string `json:"block_merkle_root"`
	EpochId          string `json:"epoch_id"`
	Height           uint64 `json:"height"`
	NextEpochId      string `json:"next_epoch_id"`
	PrevStateRoot    string `json:"prev_state_root"`
	OutcomeRoot      string `json:"outcome_root"`
	Timestamp        uint64 `json:"timestamp"`
	TimestampNanosec string `json:"timestamp_nanosec"`
	NextBpHash       string `json:"next_bp_hash"`
}

type NearNextBps struct {
	AccountId                   string `json:"account_id"`
	PublicKey                   string `json:"public_key"`
	Stake                       string `json:"stake"`
	ValidatorStakeStructVersion string `json:"validator_stake_struct_version"`
}

type Result struct {
	ApprovalsAfterNext []*string          `json:"approvals_after_next"`
	InnerLite          NearInnerLightView `json:"inner_lite"`
	InnerRestHash      string             `json:"inner_rest_hash"`
	NextBlockInnerHash string             `json:"next_block_inner_hash"`
	NextBps            []NearNextBps      `json:"next_bps"`
	PrevBlockHash      string             `json:"prev_block_hash"`
}

type NearLightClientBlockView struct {
	Jsonrpc string
	Result  Result
	Id      string
}

func (n *NearLightClientBlockView) parse() nearprimitive.LightClientBlockView {
	lb := nearprimitive.LightClientBlockView{}

	// Parse the signatures
	for _, approval := range n.Result.ApprovalsAfterNext {
		if approval == nil {
			lb.ApprovalsAfterNext = append(lb.ApprovalsAfterNext, nil)
		} else {
			// TODO: Improve error handling here
			signature := strings.Split(*approval, ":")[1]
			decode_sig := base58.Decode(signature)

			sig := &nearprimitive.Signature{}
			err := sig.TryFromRaw([]byte(decode_sig))

			if err != nil {
				fmt.Printf("Failed to decode signature: %s", err)
			}
			lb.ApprovalsAfterNext = append(lb.ApprovalsAfterNext, sig)
		}
	}

	// Parse InnerLightView
	decoded_epoch_id := base58.Decode(n.Result.InnerLite.EpochId)
	epoch_id := &nearprimitive.CryptoHash{}
	err := epoch_id.TryFromRaw(decoded_epoch_id)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	decoded_next_epoch_id := base58.Decode(n.Result.InnerLite.NextEpochId)
	next_epoch_id := &nearprimitive.CryptoHash{}
	err = next_epoch_id.TryFromRaw(decoded_next_epoch_id)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	decoded_prev_state := base58.Decode(n.Result.InnerLite.PrevStateRoot)
	prev_state := &nearprimitive.CryptoHash{}
	err = prev_state.TryFromRaw(decoded_prev_state)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	decoded_outcome_root := base58.Decode(n.Result.InnerLite.OutcomeRoot)
	outcome_root := &nearprimitive.CryptoHash{}
	err = outcome_root.TryFromRaw(decoded_outcome_root)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	decoded_next_bp_hash := base58.Decode(n.Result.InnerLite.NextBpHash)
	next_bp_hash := &nearprimitive.CryptoHash{}
	err = next_bp_hash.TryFromRaw(decoded_next_bp_hash)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	decoded_block_merkle_root := base58.Decode(n.Result.InnerLite.BlockMerkleRoot)
	block_merkle_root := &nearprimitive.CryptoHash{}
	err = block_merkle_root.TryFromRaw(decoded_block_merkle_root)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	timestamp_nanosec, err := strconv.ParseUint(n.Result.InnerLite.TimestampNanosec, 10, 64)
	if err != nil {
		fmt.Printf("Failed to parse timestamp nanosec: %s", err)
	}

	lb.InnerLite = nearprimitive.BlockHeaderInnerLiteView{
		Height:           nearprimitive.BlockHeight(n.Result.InnerLite.Height),
		EpochId:          *epoch_id,
		NextEpochId:      *next_epoch_id,
		PrevStateRoot:    *prev_state,
		BlockMerkleRoot:  *block_merkle_root,
		NextBpHash:       *next_bp_hash,
		OutcomeRoot:      *outcome_root,
		Timestamp:        n.Result.InnerLite.Timestamp,
		TimestampNanosec: timestamp_nanosec,
	}

	// Parse previous block hash
	decoded_prev_block_hash := base58.Decode(n.Result.PrevBlockHash)
	prev_block_hash := &nearprimitive.CryptoHash{}
	err = prev_block_hash.TryFromRaw(decoded_prev_block_hash)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	lb.PrevBlockHash = *prev_block_hash

	// Parse next block hash
	decoded_next_block_inner_hash := base58.Decode(n.Result.NextBlockInnerHash)
	next_block_inner_hash := &nearprimitive.CryptoHash{}
	err = next_block_inner_hash.TryFromRaw(decoded_next_block_inner_hash)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	lb.NextBlockInnerHash = *next_block_inner_hash

	// Parse inner rest hash
	decoded_inner_rest_hash := base58.Decode(n.Result.InnerRestHash)
	inner_rest_hash := &nearprimitive.CryptoHash{}
	err = inner_rest_hash.TryFromRaw(decoded_inner_rest_hash)
	if err != nil {
		fmt.Printf("Failed to decode epoch id: %s", err)
	}

	lb.InnerRestHash = *inner_rest_hash

	// Parse next BPS
	for _, bps := range n.Result.NextBps {
		vs := nearprimitive.ValidatorStakeView{}
		if bps.ValidatorStakeStructVersion == "V1" {
			vs.Version = nearprimitive.V1
		}

		vs.V1.AccountId = nearprimitive.AccountId(bps.AccountId)

		encoded_pub_key := strings.Split(bps.PublicKey, ":")[1]
		decoded_pub_key := base58.Decode(encoded_pub_key)

		pubkey := &nearprimitive.PublicKey{}
		err := pubkey.TryFromRaw(decoded_pub_key)
		if err != nil {
			fmt.Printf("Failed to parse pub key: %s", err)
		}
		vs.V1.PublicKey = *pubkey

		vs.V1.Stake, _, _ = num.U128FromString(bps.Stake)

		lb.NextBps = append(lb.NextBps, vs)
	}

	return lb
}
