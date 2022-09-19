package nearprimitive

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	borsh "github.com/near/borsh-go"
	num "github.com/shabbyrobe/go-num"
)

type CryptoHash [32]byte

func (c *CryptoHash) HashBytes(byteArray []byte) {
	digest := sha256.Sum256(byteArray[:])
	copy(c[:], digest[:32])
}

func (c *CryptoHash) AsBytes() []byte {
	return c[:]
}

func (c *CryptoHash) TryFromRaw(byteArray []byte) error {
	if len(byteArray) != 32 {
		return fmt.Errorf("Ill-formed byte array, size: %d", len(byteArray))
	}

	copy(c[:], byteArray[:32])

	return nil
}

func (c *CryptoHash) HashBorsh(borshSerializedArray []byte) error {
	data := []byte{}
	err := borsh.Deserialize(&data, borshSerializedArray)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	c.HashBytes(data)

	return nil
}

type PublicKey [32]byte

func (p *PublicKey) TryFromRaw(data []byte) error {
	if len(data) != 32 {
		return fmt.Errorf("Ill-formed public key, wrong size: %d", len(data))
	}

	copy(p[:], data[:32])

	return nil
}

func (p *PublicKey) GetEd25519PubKey() ed25519.PublicKey {
	return p[:]
}

type Signature [64]byte

func (s *Signature) AsBytes() []byte {
	return s[:]
}

func (s *Signature) TryFromRaw(data []byte) error {
	if len(data) != 64 {
		return fmt.Errorf("Ill-formed signature, wrong size: %d", len(data))
	}

	copy(s[:], data[:64])

	return nil
}

func (s *Signature) Verify(data []byte, public_key *PublicKey) bool {
	return ed25519.Verify(public_key.GetEd25519PubKey(), data, s.AsBytes())
}

type BlockHeight uint64
type AccountId string
type Balance num.U128
type Gas uint64

type MerkleHash CryptoHash

type Direction uint8

const (
	Left Direction = iota
	Right
)

type MerklePathItem struct {
	hash      MerkleHash
	direction Direction
}

type MerklePath []MerklePathItem

type BlockHeaderInnerLiteView struct {
	height            BlockHeight
	epoch_id          CryptoHash
	next_epoch_id     CryptoHash
	prev_state_root   CryptoHash
	outcome_root      CryptoHash
	timestamp         uint64
	timestamp_nanosec uint64
	next_bp_hash      CryptoHash
	block_merkle_root CryptoHash
}

func (b BlockHeaderInnerLiteView) ToBlockHeaderInnerLiteViewFinal() BlockHeaderInnerLiteViewFinal {
	return BlockHeaderInnerLiteViewFinal{
		Height:          b.height,
		EpochId:         b.epoch_id,
		NextEpochId:     b.next_epoch_id,
		PrevStateRoot:   b.prev_state_root,
		OutcomeRoot:     b.outcome_root,
		Timestamp:       b.timestamp,
		NextBpHash:      b.next_bp_hash,
		BlockMerkleRoot: b.block_merkle_root,
	}
}

type LightClientBlockLiteView struct {
	prev_block_hash CryptoHash
	inner_rest_hash CryptoHash
	inner_lite      BlockHeaderInnerLiteView
}

type ValidatorStakeViewVersion uint

const (
	V1 ValidatorStakeViewVersion = iota
)

type ValidatorStakeViewV1 struct {
	accountId  AccountId
	public_key PublicKey
	stake      Balance
}

type ValidatorStakeView struct {
	version ValidatorStakeViewVersion
	v1      ValidatorStakeViewV1
}

type LightClientBlockView struct {
	prev_block_hash       CryptoHash
	next_block_inner_hash CryptoHash
	inner_lite            BlockHeaderInnerLiteView
	inner_rest_hash       CryptoHash
	next_bps              []ValidatorStakeView
	approvals_after_next  []Signature
}

type BlockHeaderInnerLiteViewFinal struct {
	Height          BlockHeight
	EpochId         CryptoHash
	NextEpochId     CryptoHash
	PrevStateRoot   CryptoHash
	OutcomeRoot     CryptoHash
	Timestamp       uint64
	NextBpHash      CryptoHash
	BlockMerkleRoot CryptoHash
}

func (bf BlockHeaderInnerLiteViewFinal) serialize() ([]byte, error) {
	data, err := borsh.Serialize(bf)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (bf *BlockHeaderInnerLiteViewFinal) deserialize(data []byte) error {
	err := borsh.Deserialize(bf, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type ExecutionOutcomeView struct {
	logs         []string
	receipt_ids  []CryptoHash
	gas_burnt    Gas
	tokens_burnt num.U128
	executor_id  AccountId
	status       []uint8
}

type OutcomeProof struct {
	proof      []MerklePathItem
	block_hash CryptoHash
	id         CryptoHash
	outcome    ExecutionOutcomeView
}

type ApprovalInnerType uint

const (
	Endorsement ApprovalInnerType = iota
	Skip
)

type ApprovalInner struct {
	inner_type  ApprovalInnerType
	endorsement CryptoHash
	skip        BlockHeight
}

type HostFunction interface {
	sha256(data []byte) [32]byte
	verify(sig Signature, data []byte, public_key PublicKey) bool
}

//func (lb *LightClientBlockView) CurrentBlockHash(h HostFunction) CryptoHash {
//
//}
