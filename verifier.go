package light

import (
	"encoding/binary"
	"fmt"

	"github.com/electron-labs/near-light-client-go/nearprimitive"
	borsh "github.com/near/borsh-go"
)

func calculate_merklelization_hashes(h nearprimitive.HostFunction, eo nearprimitive.ExecutionOutcomeView) ([]nearprimitive.CryptoHash, error) {
	res := []nearprimitive.CryptoHash{}

	logs_payload := []byte{}
	ser_recipet_ids, err := borsh.Serialize(eo.ReceiptIds)
	if err != nil {
		return res, fmt.Errorf("Failed to serialize receipt ids: %s", err)
	}

	ser_gas_burnt, err := borsh.Serialize(eo.GasBurnt)
	if err != nil {
		return res, fmt.Errorf("Failed to serialize gas burnt: %s", err)
	}

	ser_tokens_burnt, err := borsh.Serialize(eo.TokensBurnt)
	if err != nil {
		return res, fmt.Errorf("Failed to serialize tokens burnt: %s", err)
	}

	ser_executor_id, err := borsh.Serialize(eo.ExecutorId)
	if err != nil {
		return res, fmt.Errorf("Failed to serialize executor id: %s", err)
	}

	ser_status, err := borsh.Serialize(eo.Status)
	if err != nil {
		return res, fmt.Errorf("Failed to serialize status: %s", err)
	}

	logs_payload = append(logs_payload, ser_recipet_ids...)
	logs_payload = append(logs_payload, ser_gas_burnt...)
	logs_payload = append(logs_payload, ser_tokens_burnt...)
	logs_payload = append(logs_payload, ser_executor_id...)
	logs_payload = append(logs_payload, ser_status...)

	first_elem_merkelization_hashes := h.Sha256(logs_payload)

	res = append(res, first_elem_merkelization_hashes)

	for _, log := range eo.Logs {
		res = append(res, h.Sha256([]byte(log)))
	}
	return res, nil
}

func calculate_execution_outcome_hash(h nearprimitive.HostFunction, eo nearprimitive.ExecutionOutcomeView, tx_hash nearprimitive.CryptoHash) (nearprimitive.CryptoHash, error) {
	res := &nearprimitive.CryptoHash{}
	merkelization_hashes, err := calculate_merklelization_hashes(h, eo)
	if err != nil {
		return *res, fmt.Errorf("Failed to calculate merkelization hashes: %s", err)
	}

	pack_merkelization_hash := []byte{}

	for _, hash := range merkelization_hashes {
		pack_merkelization_hash = append(pack_merkelization_hash, hash.AsBytes()...)
	}

	final_hash := []byte{}

	hash_len := make([]byte, binary.MaxVarintLen32)
	binary.LittleEndian.PutUint32(hash_len, uint32(len(merkelization_hashes)))

	final_hash = append(final_hash, hash_len...)
	final_hash = append(final_hash, tx_hash.AsBytes()...)
	final_hash = append(final_hash, pack_merkelization_hash...)

	tmp := h.Sha256(final_hash)
	err = res.TryFromRaw(tmp[:])
	if err != nil {
		return *res, fmt.Errorf("Failed to create CryptoHash from Sha256")
	}

	return *res, nil
}
