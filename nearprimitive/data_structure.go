// Copyright © 2022, Electron Labs

package nearprimitive

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"

	base58 "github.com/btcsuite/btcutil/base58"
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
type Gas uint64

type MerkleHash CryptoHash

type Direction uint8

const (
	Left Direction = iota
	Right
)

type MerklePathItem struct {
	Hash      MerkleHash
	Direction Direction
}

func (mp MerklePathItem) serialize() ([]byte, error) {
	data, err := borsh.Serialize(mp)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (mp *MerklePathItem) deserialize(data []byte) error {
	err := borsh.Deserialize(mp, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type MerklePath []MerklePathItem

type BlockHeaderInnerLiteView struct {
	Height           BlockHeight
	EpochId          CryptoHash
	NextEpochId      CryptoHash
	PrevStateRoot    CryptoHash
	OutcomeRoot      CryptoHash
	Timestamp        uint64
	TimestampNanosec uint64
	NextBpHash       CryptoHash
	BlockMerkleRoot  CryptoHash
}

func (bh BlockHeaderInnerLiteView) serialize() ([]byte, error) {
	data, err := borsh.Serialize(bh)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (bh *BlockHeaderInnerLiteView) deserialize(data []byte) error {
	err := borsh.Deserialize(bh, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

func (b BlockHeaderInnerLiteView) ToBlockHeaderInnerLiteViewFinal() BlockHeaderInnerLiteViewFinal {
	return BlockHeaderInnerLiteViewFinal{
		Height:          b.Height,
		EpochId:         b.EpochId,
		NextEpochId:     b.NextEpochId,
		PrevStateRoot:   b.PrevStateRoot,
		OutcomeRoot:     b.OutcomeRoot,
		Timestamp:       b.Timestamp,
		NextBpHash:      b.NextBpHash,
		BlockMerkleRoot: b.BlockMerkleRoot,
	}
}

type LightClientBlockLiteView struct {
	PrevBlockHash CryptoHash
	InnerRestHash CryptoHash
	InnerLite     BlockHeaderInnerLiteView
}

func (lc LightClientBlockLiteView) serialize() ([]byte, error) {
	data, err := borsh.Serialize(lc)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (lc *LightClientBlockLiteView) deserialize(data []byte) error {
	err := borsh.Deserialize(lc, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type ValidatorStakeViewVersion uint8

const (
	V1 ValidatorStakeViewVersion = iota
)

type ValidatorStakeViewV1 struct {
	AccountId AccountId
	// HACK to borsh serialize it properly
	Dummy     uint8
	PublicKey PublicKey
	Stake     num.U128
}

type ValidatorStakeView struct {
	Version ValidatorStakeViewVersion
	V1      ValidatorStakeViewV1
}

func (vs ValidatorStakeView) serialize() ([]byte, error) {
	data, err := borsh.Serialize(vs)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (vs *ValidatorStakeView) deserialize(data []byte) error {
	err := borsh.Deserialize(vs, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

func (v *ValidatorStakeView) GetValidatorStake() (ValidatorStakeViewV1, error) {
	if v.Version == V1 {
		return v.V1, nil
	}

	return v.V1, fmt.Errorf("Invalid version %v", v.Version)
}

type LightClientBlockView struct {
	PrevBlockHash      CryptoHash
	NextBlockInnerHash CryptoHash
	InnerLite          BlockHeaderInnerLiteView
	InnerRestHash      CryptoHash
	NextBps            []ValidatorStakeView
	ApprovalsAfterNext []*Signature
}

func (lb LightClientBlockView) serialize() ([]byte, error) {
	data, err := borsh.Serialize(lb)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (lb *LightClientBlockView) deserialize(data []byte) error {
	err := borsh.Deserialize(lb, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
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
	Logs        []string
	ReceiptIds  []CryptoHash
	GasBurnt    Gas
	TokensBurnt num.U128
	ExecutorId  AccountId
	Status      []uint8
}

func (eo ExecutionOutcomeView) serialize() ([]byte, error) {
	data, err := borsh.Serialize(eo)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (eo *ExecutionOutcomeView) deserialize(data []byte) error {
	err := borsh.Deserialize(eo, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type OutcomeProof struct {
	Proof     []MerklePathItem
	BlockHash CryptoHash
	Id        CryptoHash
	Outcome   ExecutionOutcomeView
}

func (op OutcomeProof) serialize() ([]byte, error) {
	data, err := borsh.Serialize(op)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (op *OutcomeProof) deserialize(data []byte) error {
	err := borsh.Deserialize(op, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type ApprovalInnerType uint8

const (
	Endorsement ApprovalInnerType = iota
	Skip
)

type ApprovalInner struct {
	InnerType   ApprovalInnerType
	Endorsement CryptoHash
	Skip        BlockHeight
}

func (ai ApprovalInner) serialize() ([]byte, error) {
	data, err := borsh.Serialize(ai)
	if err != nil {
		return data, fmt.Errorf("Failed to serialize: %s", err)
	}

	return data, nil
}

func (ai *ApprovalInner) deserialize(data []byte) error {
	err := borsh.Deserialize(ai, data)
	if err != nil {
		return fmt.Errorf("Failed to deserialize: %s", err)
	}

	return nil
}

type HostFunction interface {
	Sha256(data []byte) [32]byte
	Verify(sig Signature, data []byte, public_key PublicKey) bool
}

func (lb *LightClientBlockView) CurrentBlockHash(h HostFunction) (CryptoHash, error) {
	inner_lite_ser, err := lb.InnerLite.ToBlockHeaderInnerLiteViewFinal().serialize()
	if err != nil {
		return CryptoHash{}, fmt.Errorf("Failed to serialize inner lite: %s", err)
	}

	inner_lite_hash := h.Sha256(inner_lite_ser)
	c := &CryptoHash{}

	appended_hashes := append(inner_lite_hash[:], lb.InnerRestHash.AsBytes()...)
	new_inner_hash := h.Sha256(appended_hashes[:])
	appended_hashes = append(new_inner_hash[:], lb.PrevBlockHash.AsBytes()...)

	c.HashBytes(appended_hashes)

	return *c, nil
}

type Unknown struct{}

type SuccessValue struct {
	Inner string
}

type SuccessReceiptID struct {
	Inner CryptoHash
}

type ExecutionStatusView struct {
	Enum borsh.Enum `borsh_enum:"true"`
	// The execution is pending or unknown.
	Unknown Unknown
	// The execution has failed.
	Failure []byte
	// The final action succeeded and returned some value or an empty vec encoded in base64.
	SuccessValue SuccessValue
	// The final action of the receipt returned a promise or the signed transaction was converted
	// to a receipt. Contains the receipt_id of the generated receipt.
	SuccessReceiptID SuccessReceiptID
}

func IntoExecutionStatusView(raw_status map[string]json.RawMessage) (ExecutionStatusView, error) {
	var status ExecutionStatusView
	const unknown = 0
	const executionStatusView = 2
	const successReceiptID = 3
	for k, v := range raw_status {
		switch k {
		case "Uknonwn":
			var s string
			err := json.Unmarshal([]byte(v), &s)
			if err != nil {
				return status, err
			}
			status = ExecutionStatusView{
				Enum:    unknown,
				Unknown: Unknown{},
			}
		case "Failure":
			fmt.Println("Unsupported failure transaction")
			return status, errors.New("unsupported failure transaction")
		case "SuccessValues":
			var s string
			err := json.Unmarshal([]byte(v), &s)
			if err != nil {
				return status, err
			}
			status = ExecutionStatusView{
				Enum: executionStatusView,
				SuccessValue: SuccessValue{
					Inner: s,
				},
			}
		case "SuccessReceiptId":
			var s string
			err := json.Unmarshal([]byte(v), &s)
			if err != nil {
				return status, err
			}
			cryptoHash := CryptoHash{}
			cryptoHash.TryFromRaw(base58.Decode(s))
			status = ExecutionStatusView{
				Enum: successReceiptID,
				SuccessReceiptID: SuccessReceiptID{
					Inner: cryptoHash,
				},
			}
		}
	}

	return status, nil
}
