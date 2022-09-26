// Copyright © 2022, Electron Labs

package light

import (
	"github.com/electron-labs/near-light-client-go/nearprimitive"
)

type NearLightClientInterface interface {
	NewFromCheckpoint(checkpoint nearprimitive.LightClientBlockView, heights_to_track uint64)
	CurrentBlockHeight() uint64
}
