package light

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/electron-labs/near-light-client-go/mock"
	"github.com/electron-labs/near-light-client-go/nearprimitive"
	borsh "github.com/near/borsh-go"
)

func BlockMerkleRootVerification(lcResp string, execResp string) error {
	nlc_json := NearLightClientBlockView{}
	err := json.Unmarshal([]byte(lcResp), &nlc_json)
	if err != nil {
		return fmt.Errorf("Failed to parse light client block: %s", err)
	}

	tx_proof_json, err := GetTxProof(execResp)
	if err != nil {
		return fmt.Errorf("Failed to parse tx proof: %s", err)
	}

	tx_proof, err := tx_proof_json.parse()
	if err != nil {
		return fmt.Errorf("Failed to parse tx_proof: %s", err)
	}

	ser_inner_lite, err := borsh.Serialize(tx_proof.BlockHeaderLite.InnerLite.ToBlockHeaderInnerLiteViewFinal())
	if err != nil {
		return fmt.Errorf("Failed to serialize: %s", err)
	}

	h := mock.MockHostFunction{}

	sha_inner_lite := h.Sha256(ser_inner_lite)
	re := CurrentBlockHash(h, sha_inner_lite, tx_proof.BlockHeaderLite.InnerRestHash, tx_proof.BlockHeaderLite.PrevBlockHash)

	near_light_client_block_view := nlc_json.parse()

	root, err := compute_root_from_path(h, tx_proof.BlockProof, nearprimitive.MerkleHash(re))
	if err != nil {
		return fmt.Errorf("Failed to compute root: %s", err)
	}

	if !bytes.Equal(near_light_client_block_view.InnerLite.BlockMerkleRoot[:], root[:]) {
		return fmt.Errorf("Failed to verify merkle root!")
	} else {
		return nil
	}
}
