//go:build localnet
// +build localnet

package build

import (
	"os"
	"strconv"

	"github.com/filecoin-project/go-address"
)

var BlockDelaySecs uint64
var EquivocationDelaySecs uint64
var PropagationDelaySecs uint64

const UpgradeSmokeHeight = -2

func init() {
	SetAddressNetwork(address.Testnet)
	BuildType = BuildLocalnet

	// Network parameters from environment variables with defaults
	if v := os.Getenv("FOC_LOCALNET_BLOCK_DELAY"); v != "" {
		if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
			BlockDelaySecs = parsed
		} else {
			panic("FOC_LOCALNET_BLOCK_DELAY must be a valid uint64")
		}
	} else {
		BlockDelaySecs = 4 // default
	}

	if v := os.Getenv("FOC_LOCALNET_PROPAGATION_DELAY"); v != "" {
		if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
			PropagationDelaySecs = parsed
		} else {
			panic("FOC_LOCALNET_PROPAGATION_DELAY must be a valid uint64")
		}
	} else {
		PropagationDelaySecs = 1 // default
	}

	if v := os.Getenv("FOC_LOCALNET_EQUIVOCATION_DELAY"); v != "" {
		if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
			EquivocationDelaySecs = parsed
		} else {
			panic("FOC_LOCALNET_EQUIVOCATION_DELAY must be a valid uint64")
		}
	} else {
		EquivocationDelaySecs = 0 // default
	}
}
