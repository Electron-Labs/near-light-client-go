package light

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/electron-labs/near-light-client-go/nearprimitive"
	borsh "github.com/near/borsh-go"
	num "github.com/shabbyrobe/go-num"
)

func get_client_block_view(client_block_response string) (nearprimitive.LightClientBlockView, error) {
	var block_view NearLightClientBlockView

	err := json.Unmarshal([]byte(client_block_response), &block_view)
	if err != nil {
		return nearprimitive.LightClientBlockView{}, fmt.Errorf("Failed to parse client block: %s", err)
	}

	return block_view.parse(), nil
}

func next_block_hash(h nearprimitive.HostFunction, next_block_inner_hash nearprimitive.CryptoHash, current_block_hash nearprimitive.CryptoHash) (nearprimitive.CryptoHash, error) {
	final_hash := []byte{}
	final_hash = append(final_hash, next_block_inner_hash.AsBytes()...)
	final_hash = append(final_hash, current_block_hash.AsBytes()...)

	sha_hash := h.Sha256(final_hash)

	res := &nearprimitive.CryptoHash{}
	err := res.TryFromRaw(sha_hash[:])

	if err != nil {
		return *res, fmt.Errorf("Failed to convert to CryptoHash: %s", err)
	}

	return *res, nil
}

func reconstruct_light_client_block_view_fields(h nearprimitive.HostFunction, block_view nearprimitive.LightClientBlockView) (nearprimitive.CryptoHash, nearprimitive.CryptoHash, []byte, error) {
	current_block_hash, err := block_view.CurrentBlockHash(h)
	if err != nil {
		return nearprimitive.CryptoHash{}, nearprimitive.CryptoHash{}, []byte{}, fmt.Errorf("Failed to get current block hash: %s", err)
	}

	next_block_hash, err := next_block_hash(h, block_view.NextBlockInnerHash, current_block_hash)
	if err != nil {
		return nearprimitive.CryptoHash{}, nearprimitive.CryptoHash{}, []byte{}, fmt.Errorf("Failed to get next block hash: %s", err)
	}

	// TODO: Might need to borsh serialize ApprovalMessage
	// HACK: Add 0 to indicate first type of Enum
	approval_message := []byte{0}
	approval_message = append(approval_message, next_block_hash.AsBytes()...)

	lite_height_to_bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lite_height_to_bytes, uint64(block_view.InnerLite.Height)+2)
	approval_message = append(approval_message, lite_height_to_bytes...)

	return current_block_hash, next_block_hash, approval_message, nil
}

func ValidateLightBlock(h nearprimitive.HostFunction, head *nearprimitive.LightClientBlockView, block_view *nearprimitive.LightClientBlockView, epoch_block_producers_map map[nearprimitive.CryptoHash][]nearprimitive.ValidatorStakeView) error {
	_, _, approval_message, err := reconstruct_light_client_block_view_fields(h, *block_view)
	if err != nil {
		return fmt.Errorf("Failed to reconstruct light client block view fields: %s", err)
	}

	if block_view.InnerLite.Height <= head.InnerLite.Height {
		return fmt.Errorf("Block view height is not ahead of the head's height")
	}

	if !(block_view.InnerLite.EpochId == head.InnerLite.EpochId || block_view.InnerLite.EpochId == head.InnerLite.NextEpochId) {
		return fmt.Errorf("Block view epoch id not present in the head %v %v %v", block_view.InnerLite.EpochId, head.InnerLite.EpochId, head.InnerLite.NextEpochId)
	}

	if block_view.InnerLite.EpochId == head.InnerLite.NextEpochId && len(block_view.NextBps) == 0 {
		return fmt.Errorf("Block view epoch id is not the next epoch")
	}

	total_stake := num.U128{}
	approved_stake := num.U128{}

	epoch_block_producers := epoch_block_producers_map[block_view.InnerLite.EpochId]

	for i, signature := range block_view.ApprovalsAfterNext {
		block_producer := epoch_block_producers[i]
		bp_stake_view, err := block_producer.GetValidatorStake()

		if err != nil {
			return fmt.Errorf("Failed to retrieve validator stake %v", i)
		}

		bp_stake := bp_stake_view.Stake
		total_stake = total_stake.Add(bp_stake)

		if signature == nil {
			continue
		}

		approved_stake = approved_stake.Add(bp_stake)

		validator_pub_key := bp_stake_view.PublicKey
		if !signature.Verify(approval_message, &validator_pub_key) {
			return fmt.Errorf("Failed to verify the signature for %s", approval_message)
		}
	}

	threshold := total_stake.Mul64(2).Quo64(3)
	if approved_stake.LessOrEqualTo(threshold) {
		return fmt.Errorf("Block is not final: stake threshold is not reached")
	}

	// HACK
	next_bps := []nearprimitive.ValidatorStakeView{}

	for _, bps := range block_view.NextBps {
		stake := bps.V1.Stake
		stake_lo := stake.AsUint64()
		stake_hi := stake.Rsh(64).AsUint64()
		bps.V1.Stake = num.U128FromRaw(stake_lo, stake_hi)

		next_bps = append(next_bps, bps)
	}

	if len(block_view.NextBps) > 0 {
		ser_block_view_next_bps, err := borsh.Serialize(next_bps)
		if err != nil {
			return fmt.Errorf("Failed to serialize block view next bps")
		}

		next_bps_hash := h.Sha256(ser_block_view_next_bps)
		if next_bps_hash != block_view.InnerLite.NextBpHash {
			return fmt.Errorf("Incorrect next bp hash in block view")
		}
	}

	return nil
}
