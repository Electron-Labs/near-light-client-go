// Copyright © 2022, Electron Labs

package light

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/electron-labs/near-light-client-go/nearprimitive"
	"github.com/near/borsh-go"

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

func (il NearInnerLightView) parse() (nearprimitive.BlockHeaderInnerLiteView, error) {
	decoded_epoch_id := base58.Decode(il.EpochId)
	epoch_id := &nearprimitive.CryptoHash{}
	err := epoch_id.TryFromRaw(decoded_epoch_id)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode epoch id: %s", err)
	}

	decoded_next_epoch_id := base58.Decode(il.NextEpochId)
	next_epoch_id := &nearprimitive.CryptoHash{}
	err = next_epoch_id.TryFromRaw(decoded_next_epoch_id)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode next epoch id: %s", err)
	}

	decoded_prev_state := base58.Decode(il.PrevStateRoot)
	prev_state := &nearprimitive.CryptoHash{}
	err = prev_state.TryFromRaw(decoded_prev_state)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode prev state root: %s", err)
	}

	decoded_outcome_root := base58.Decode(il.OutcomeRoot)
	outcome_root := &nearprimitive.CryptoHash{}
	err = outcome_root.TryFromRaw(decoded_outcome_root)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode outcome root: %s", err)
	}

	decoded_next_bp_hash := base58.Decode(il.NextBpHash)
	next_bp_hash := &nearprimitive.CryptoHash{}
	err = next_bp_hash.TryFromRaw(decoded_next_bp_hash)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode next bp hash: %s", err)
	}

	decoded_block_merkle_root := base58.Decode(il.BlockMerkleRoot)
	block_merkle_root := &nearprimitive.CryptoHash{}
	err = block_merkle_root.TryFromRaw(decoded_block_merkle_root)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to decode block merkle root: %s", err)
	}

	timestamp_nanosec, err := strconv.ParseUint(il.TimestampNanosec, 10, 64)
	if err != nil {
		return nearprimitive.BlockHeaderInnerLiteView{}, fmt.Errorf("Failed to parse timestamp nanosec: %s", err)
	}

	bh := nearprimitive.BlockHeaderInnerLiteView{
		Height:           nearprimitive.BlockHeight(il.Height),
		EpochId:          *epoch_id,
		NextEpochId:      *next_epoch_id,
		PrevStateRoot:    *prev_state,
		BlockMerkleRoot:  *block_merkle_root,
		NextBpHash:       *next_bp_hash,
		OutcomeRoot:      *outcome_root,
		Timestamp:        timestamp_nanosec,
		TimestampNanosec: timestamp_nanosec,
	}

	return bh, nil
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

type BlockHeaderLite struct {
	InnerLite     NearInnerLightView `json:"inner_lite"`
	InnerRestHash string             `json:"inner_rest_hash"`
	PrevBlockHash string             `json:"prev_block_hash"`
}

func (b BlockHeaderLite) parse() (nearprimitive.LightClientBlockLiteView, error) {
	inner_lite, err := b.InnerLite.parse()
	if err != nil {
		return nearprimitive.LightClientBlockLiteView{}, fmt.Errorf("Failed to parse inner lite: %s", err)
	}

	inner_rest_hash := &nearprimitive.CryptoHash{}
	err = inner_rest_hash.TryFromRaw(base58.Decode(b.InnerRestHash))
	if err != nil {
		return nearprimitive.LightClientBlockLiteView{}, fmt.Errorf("Failed to parse inner rest hash: %s", err)
	}

	prev_block_hash := &nearprimitive.CryptoHash{}
	err = prev_block_hash.TryFromRaw(base58.Decode(b.PrevBlockHash))
	if err != nil {
		return nearprimitive.LightClientBlockLiteView{}, fmt.Errorf("Failed to parse inner rest hash: %s", err)
	}

	lb := nearprimitive.LightClientBlockLiteView{
		PrevBlockHash: *prev_block_hash,
		InnerRestHash: *inner_rest_hash,
		InnerLite:     inner_lite,
	}

	return lb, nil
}

type MetaData struct {
	GasProfile *string `json:"gas_profile"`
	Version    uint64  `json:"version"`
}

type Status struct {
	SuccessReceiptId string `json:"SuccessReceiptId"`
}

type Outcome struct {
	ExecutorId  string                     `json:"executor_id"`
	GasBurnt    uint64                     `json:"gas_burnt"`
	Logs        []string                   `json:"logs"`
	MetaData    MetaData                   `json:"metadata"`
	ReceiptIds  []string                   `json:"receipt_ids"`
	Status      map[string]json.RawMessage `json:"status"`
	TokensBurnt string                     `json:"tokens_burnt"`
}

type Proof []struct {
	Direction string `json:"direction"`
	Hash      string `json:"hash"`
}

type OutcomeProof struct {
	BlockHash string  `json:"block_hash"`
	Id        string  `json:"id"`
	Outcome   Outcome `json:"outcome"`
	Proof     Proof   `json:"proof"`
}

type TxResult struct {
	BlockHeaderLite  BlockHeaderLite `json:"block_header_lite"`
	BlockProof       Proof           `json:"block_proof"`
	OutcomeProof     OutcomeProof    `json:"outcome_proof"`
	OutcomeRootProof Proof           `json:"outcome_root_proof"`
}

type NearTxResult struct {
	BlockHeaderLite  nearprimitive.LightClientBlockLiteView
	BlockProof       nearprimitive.MerklePath
	OutcomeProof     nearprimitive.OutcomeProof
	OutcomeRootProof nearprimitive.MerklePath
}

func (tx TxResult) Parse() (NearTxResult, error) {
	near_tx_result := NearTxResult{}

	lite_header, err := tx.BlockHeaderLite.parse()
	if err != nil {
		return near_tx_result, fmt.Errorf("Failed to parse lite header: %s", err)
	}

	block_proof, err := tx.BlockProof.parse()
	if err != nil {
		return near_tx_result, fmt.Errorf("Failed to parse block proof: %s", err)
	}

	outcome_proof, err := tx.OutcomeProof.parse()
	if err != nil {
		return near_tx_result, fmt.Errorf("Failed to parse outcome proof: %s", err)
	}

	outcome_root_proof, err := tx.OutcomeRootProof.parse()
	if err != nil {
		return near_tx_result, fmt.Errorf("Failed to parse outcome root proof: %s", err)
	}

	near_tx_result.BlockHeaderLite = lite_header
	near_tx_result.BlockProof = block_proof
	near_tx_result.OutcomeProof = outcome_proof
	near_tx_result.OutcomeRootProof = outcome_root_proof

	return near_tx_result, nil
}

func (bp Proof) parse() (nearprimitive.MerklePath, error) {
	merklePath := nearprimitive.MerklePath{}

	for i := 0; i < len(bp); i++ {
		path_item_hash := &nearprimitive.CryptoHash{}
		err := path_item_hash.TryFromRaw(base58.Decode(bp[i].Hash))
		if err != nil {
			return merklePath, fmt.Errorf("Failed to decode hash to base58: %s", err)
		}
		if bp[i].Direction == "Right" {
			merklePath = append(merklePath, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Right})
		} else {
			merklePath = append(merklePath, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})
		}
	}
	return merklePath, nil
}

func (op OutcomeProof) parse() (nearprimitive.OutcomeProof, error) {
	receipt_ids := []nearprimitive.CryptoHash{}
	single_receipt := nearprimitive.CryptoHash{}

	for i := 0; i < len(op.Outcome.ReceiptIds); i++ {
		err := single_receipt.TryFromRaw(base58.Decode(op.Outcome.ReceiptIds[i]))
		receipt_ids = append(receipt_ids, single_receipt)
		if err != nil {
			return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to get receipt Id: %s", err)
		}
	}

	token_burnt, _, _ := num.U128FromString(op.Outcome.TokensBurnt)

	serialized_status, err := nearprimitive.IntoExecutionStatusView(op.Outcome.Status)
	if err != nil {
		return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to parse status: %s", err)
	}

	ser_status, err := borsh.Serialize(serialized_status)
	if err != nil {
		return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to serialize status: %s", err)
	}

	execution_outcome := nearprimitive.ExecutionOutcomeView{
		Logs:        op.Outcome.Logs,
		ReceiptIds:  receipt_ids,
		GasBurnt:    nearprimitive.Gas(op.Outcome.GasBurnt),
		TokensBurnt: token_burnt,
		ExecutorId:  nearprimitive.AccountId(op.Outcome.ExecutorId),
		Status:      ser_status,
	}

	id := &nearprimitive.CryptoHash{}
	err = id.TryFromRaw(base58.Decode(op.Id))
	if err != nil {
		return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to decode Outcome Proof Id: %s", err)
	}

	block_hash := &nearprimitive.CryptoHash{}
	err = block_hash.TryFromRaw(base58.Decode(op.BlockHash))
	err = id.TryFromRaw(base58.Decode(op.Id))

	if err != nil {
		return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to Decode block Hash: %s", err)
	}

	proof, err := op.Proof.parse()
	if err != nil {
		return nearprimitive.OutcomeProof{}, fmt.Errorf("Failed to parse proof: %s", err)
	}

	outcome_proof := nearprimitive.OutcomeProof{
		Proof:     proof,
		BlockHash: *block_hash,
		Id:        *id,
		Outcome:   execution_outcome,
	}

	return outcome_proof, nil
}

type TxRpcResponse struct {
	Id      string   `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Result  TxResult `json:"result"`
}

// GetProof will form Merkle path from Json response of merkle path
func GetProof(bp Proof) []nearprimitive.MerklePathItem {

	merklePath := []nearprimitive.MerklePathItem{}
	for i := 0; i < len(bp); i++ {
		path_item_hash := &nearprimitive.CryptoHash{}
		err := path_item_hash.TryFromRaw(base58.Decode(bp[i].Hash))
		if err != nil {
			fmt.Printf("Failed to decode hash to base58: %s", err)
		}
		if bp[i].Direction == "Right" {
			merklePath = append(merklePath, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Right})
		} else {
			merklePath = append(merklePath, nearprimitive.MerklePathItem{Hash: nearprimitive.MerkleHash(*path_item_hash), Direction: nearprimitive.Left})
		}
	}
	return merklePath
}

func GetTxProof(response string) (TxResult, error) {
	bp := TxRpcResponse{}

	err := json.Unmarshal([]byte(response), &bp)
	if err != nil {
		return bp.Result, fmt.Errorf("Failed to unmarshal RpcResponse: %s", err)
	}

	return bp.Result, nil
}

// GetOutcomeProof will give outcome proof and outcome root proof from Rpc json response from near node
func GetOutcomeProof(response string) (nearprimitive.OutcomeProof, []nearprimitive.MerklePathItem, nearprimitive.CryptoHash) {
	bp := TxRpcResponse{}

	err := json.Unmarshal([]byte(response), &bp)
	if err != nil {
		fmt.Printf("Failed to unmarshal RpcResponse: %s", err)
	}

	expected_outcome_root := &nearprimitive.CryptoHash{}
	err = expected_outcome_root.TryFromRaw(base58.Decode(bp.Result.BlockHeaderLite.InnerLite.OutcomeRoot))
	if err != nil {
		fmt.Printf("Failed to get expected_outcome_root: %s", err)
	}

	receipt_ids := []nearprimitive.CryptoHash{}
	single_receipt := nearprimitive.CryptoHash{}

	for i := 0; i < len(bp.Result.OutcomeProof.Outcome.ReceiptIds); i++ {
		err = single_receipt.TryFromRaw(base58.Decode(bp.Result.OutcomeProof.Outcome.ReceiptIds[i]))
		receipt_ids = append(receipt_ids, single_receipt)
		if err != nil {
			fmt.Printf("Failed to get receipt Id: %s", err)
		}
	}

	token_burnt, _, _ := num.U128FromString(bp.Result.OutcomeProof.Outcome.TokensBurnt)

	serialized_status, err := nearprimitive.IntoExecutionStatusView(bp.Result.OutcomeProof.Outcome.Status)
	if err != nil {
		fmt.Printf("Failed to parse status: %s", err)
	}

	ser_status, err := borsh.Serialize(serialized_status)
	if err != nil {
		fmt.Printf("Failed to serialize status: %s", err)
	}

	execution_outcome := nearprimitive.ExecutionOutcomeView{
		Logs:        bp.Result.OutcomeProof.Outcome.Logs,
		ReceiptIds:  receipt_ids,
		GasBurnt:    nearprimitive.Gas(bp.Result.OutcomeProof.Outcome.GasBurnt),
		TokensBurnt: token_burnt,
		ExecutorId:  nearprimitive.AccountId(bp.Result.OutcomeProof.Outcome.ExecutorId),
		Status:      ser_status,
	}

	id := &nearprimitive.CryptoHash{}
	err = id.TryFromRaw(base58.Decode(bp.Result.OutcomeProof.Id))
	if err != nil {
		fmt.Printf("Failed to decode Outcome Proof Id: %s", err)
	}

	block_hash := &nearprimitive.CryptoHash{}
	err = block_hash.TryFromRaw(base58.Decode(bp.Result.OutcomeProof.BlockHash))
	err = id.TryFromRaw(base58.Decode(bp.Result.OutcomeProof.Id))

	if err != nil {
		fmt.Printf("Failed to Decode block Hash: %s", err)
	}

	outcome_proof := nearprimitive.OutcomeProof{
		Proof:     GetProof(bp.Result.OutcomeProof.Proof),
		BlockHash: *block_hash,
		Id:        *id,
		Outcome:   execution_outcome,
	}

	return outcome_proof, GetProof(bp.Result.OutcomeRootProof), *expected_outcome_root
}

func (n *NearLightClientBlockView) Parse() nearprimitive.LightClientBlockView {
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
		Timestamp:        timestamp_nanosec,
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
