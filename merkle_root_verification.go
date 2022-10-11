package light

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
	"github.com/near/borsh-go"
)

func BlockMerkleRoot(lc_block string) nearprimitive.MerkleHash {
	bp := NearLightClientBlockView{}

	block_merkle_root := nearprimitive.CryptoHash{}
	err := block_merkle_root.TryFromRaw(base58.Decode(bp.Result.InnerLite.BlockMerkleRoot))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	return nearprimitive.MerkleHash(block_merkle_root)
}
func FromJsonToLCliteView(res string) (nearprimitive.LightClientBlockLiteView, []nearprimitive.MerklePathItem) {
	bp := TxRpcResponse{}

	err := json.Unmarshal([]byte(res), &bp)
	if err != nil {
		fmt.Printf("Failed to unmarshal RpcResponse: %s", err)
	}

	LCLiteView := nearprimitive.LightClientBlockLiteView{}
	inner_rest_hash := nearprimitive.CryptoHash{}
	err = inner_rest_hash.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerRestHash))
	prev_block_hash := nearprimitive.CryptoHash{}
	err = prev_block_hash.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.PrevBlockHash))

	epoch_Id := nearprimitive.CryptoHash{}
	err = epoch_Id.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed decode: %s", err)
	}

	next_epoch_Id := nearprimitive.CryptoHash{}
	err = next_epoch_Id.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	prev_state_root := nearprimitive.CryptoHash{}
	err = prev_state_root.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	outcome_root := nearprimitive.CryptoHash{}
	err = outcome_root.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	timestamp_nanosec, err := strconv.ParseUint(bp.Result.BlockHeaderLite.InnerLite.TimestampNanosec, 10, 64)
	if err != nil {
		fmt.Printf("Failed to parse timestamp nanosec: %s", err)
	}

	next_bp_hash := nearprimitive.CryptoHash{}
	err = next_bp_hash.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	block_merkle_root := nearprimitive.CryptoHash{}
	err = block_merkle_root.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.EpochId))
	if err != nil {
		fmt.Printf("Failed to decode: %s", err)
	}

	inner_lite := nearprimitive.BlockHeaderInnerLiteView{
		Height:           nearprimitive.BlockHeight(bp.Result.BlockHeaderLite.InnerLite.Height),
		EpochId:          epoch_Id,
		NextEpochId:      next_epoch_Id,
		PrevStateRoot:    prev_state_root,
		OutcomeRoot:      outcome_root,
		Timestamp:        timestamp_nanosec,
		TimestampNanosec: timestamp_nanosec,
		NextBpHash:       next_bp_hash,
		BlockMerkleRoot:  block_merkle_root,
	}

	LCLiteView.InnerLite = inner_lite
	LCLiteView.InnerRestHash = inner_rest_hash
	LCLiteView.PrevBlockHash = prev_block_hash

	// parse block_proof

	block_proof := GetProof(bp.Result.BlockProof)

	return LCLiteView, block_proof

}

func ComputeBlockHash(h nearprimitive.HostFunction, LCliteView nearprimitive.LightClientBlockLiteView) (nearprimitive.MerkleHash, error) {
	block_hash := nearprimitive.MerkleHash{}
	bytes_inner_lite, err := borsh.Serialize(LCliteView.InnerLite)
	if err != nil {
		return block_hash, err
	}
	sha_inner_lite := h.Sha256(bytes_inner_lite)

	bytes_inner_rest, err := borsh.Serialize(LCliteView.InnerRestHash)
	if err != nil {
		return block_hash, err
	}
	sha_inner_rest := h.Sha256(bytes_inner_rest)

	comb_lite_rest_hash, err := combine_hash(mock.MockHostFunction{}, sha_inner_lite, sha_inner_rest)
	if err != nil {
		return block_hash, err
	}
	comb_prev_lite_rest, err := combine_hash(mock.MockHostFunction{}, comb_lite_rest_hash, nearprimitive.MerkleHash(LCliteView.PrevBlockHash))
	if err != nil {
		return block_hash, err
	}
	block_hash = comb_prev_lite_rest
	return block_hash, nil
}
