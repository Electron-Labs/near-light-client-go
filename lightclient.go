package light

import (
	// "github.com/near/borsh-go"
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"github.com/near/borsh-go"
)

const LightClientValidationError = Error("validation failed")

type ApprovalInner struct {
	Endorsement [32]byte
	BlockHeight uint64
}

type PublicKey struct {
	K [32]byte
}

type Signature struct {
	R [32]byte
	S [32]byte
}

type BlockProducer struct {
	PublickKey  PublicKey
	Stake       uint // Should be u128
	IsChunkOnly bool
}

type OptionalBlockProducers struct {
	Some            bool
	BlockeProducers []BlockProducer
}

type OptionalSignature struct {
	Some      bool
	Signature Signature
}

type BlockHeaderInnerLite struct {
	Height          uint64
	EpochId         [32]byte
	NextEpochId     [32]byte
	PrevStateRoot   [32]byte
	OutcomeRoot     [32]byte
	TimeStamp       uint64
	NextBpHash      [32]byte
	BlockMerkleRoot [32]byte
}

type LightClientBlock struct {
	PrevBlockHash      [32]byte
	NextBlockInnerHash [32]byte
	InnerLite          BlockHeaderInnerLite
	InnerRestHash      [32]byte
	NextBps            OptionalBlockProducers
	ApprovalsAfterNext []OptionalSignature
}

func reconstructLightClientBlockViewField(blockView LightClientBlock) ([]byte, []byte, []byte, error) {
	innerLiteSerialized, err := borsh.Serialize(blockView.InnerLite)
	if err != nil {
		return nil, nil, nil, err
	}

	hashInnerLite := sha256.Sum256(innerLiteSerialized)
	hashInnerLiteHashAndRestHash := sha256.Sum256(append(
		hashInnerLite[:],
		blockView.InnerRestHash[:]...,
	))

	currentBlockHash := sha256.Sum256(append(
		hashInnerLiteHashAndRestHash[:],
		blockView.PrevBlockHash[:]...,
	))

	nextBlockHash := sha256.Sum256(append(
		blockView.NextBlockInnerHash[:],
		currentBlockHash[:]...,
	))

	nextBlockHashSerialized, err := borsh.Serialize(nextBlockHash)
	if err != nil {
		return nil, nil, nil, err
	}

	var newHeightSerialized [8]byte
	binary.LittleEndian.PutUint64(newHeightSerialized[:], blockView.InnerLite.Height+2)
	approvalMessage := append(
		nextBlockHashSerialized,
		newHeightSerialized[:]...,
	)

	return currentBlockHash[:], nextBlockHash[:], approvalMessage, nil
}

func ValidateAndUpdateHead(
	blockView LightClientBlock,
	head LightClientBlock,
	epochBlockProducersMap map[[32]byte][]BlockProducer,
) (LightClientBlock, map[[32]byte][]BlockProducer, error) {

	_, _, _, err := reconstructLightClientBlockViewField(blockView)
	if err != nil {
		return LightClientBlock{}, nil, err
	}

	if blockView.InnerLite.Height <= head.InnerLite.Height {
		return LightClientBlock{}, nil, LightClientValidationError
	}

	if !bytes.Equal(blockView.InnerLite.EpochId[:], head.InnerLite.EpochId[:]) ||
		!bytes.Equal(blockView.InnerLite.EpochId[:], head.InnerLite.NextEpochId[:]) {
		return LightClientBlock{}, nil, LightClientValidationError
	}

	if bytes.Equal(blockView.InnerLite.EpochId[:], head.InnerLite.NextEpochId[:]) &&
		!blockView.NextBps.Some {
		return LightClientBlock{}, nil, LightClientValidationError
	}

	totalStake := uint(0)
	approvedStake := uint(0)

	epochBlockProducers := epochBlockProducersMap[blockView.InnerLite.EpochId]

	for i := 0; i < len(epochBlockProducers); i++ {
		stake := epochBlockProducers[i].Stake
		totalStake += stake

		if len(blockView.ApprovalsAfterNext) <= i {
			continue
		}

		approvedStake += stake
		//if !verifySignature(epochBlockProducers[i].PublickKey, blockView.ApprovalsAfterNext[i], approvalMessage) {
		//	return LightClientBlock{}, nil, LightClientValidationError
		//}
	}

	threshold := (totalStake * 2) / 3
	if approvedStake < threshold {
		return LightClientBlock{}, nil, LightClientValidationError
	}

	if blockView.NextBps.Some {
		nextBpsSerialized, err := borsh.Serialize(blockView.NextBps)
		if err != nil {
			return LightClientBlock{}, nil, err
		}
		nextBpsHash := sha256.Sum256(nextBpsSerialized)
		if !bytes.Equal(nextBpsHash[:], blockView.InnerLite.NextBpHash[:]) {
			return LightClientBlock{}, nil, LightClientValidationError
		}

		epochBlockProducersMap[blockView.InnerLite.NextEpochId] = blockView.NextBps.BlockeProducers
	}

	return blockView, epochBlockProducersMap, nil
}
