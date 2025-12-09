# Curio Localnet Configuration

This document describes the environment variables required to run Curio in localnet mode for local development.

## Building for Localnet

To build Curio for localnet mode, use the `localnet` build tag:

```bash
go build -tags localnet ./cmd/curio
```

## Required Environment Variables

### Contract Addresses

The following contract addresses **must** be set when running Curio in localnet mode:

- **`FOC_LOCALNET_CONTRACT_PAY`** - Filecoin Pay contract address
  - Used for payment processing
  - Must be a valid hex address (e.g., `0x1234...`)

- **`FOC_LOCALNET_CONTRACT_FWSS`** - Filecoin Warm Storage Service contract address
  - Used as a whitelist service for record keepers
  - Must be a valid hex address

- **`FOC_LOCALNET_CONTRACT_MULTICALL`** - Multicall3 contract address
  - Used for batching multiple contract calls
  - Must be a valid hex address

### Optional Contract Addresses

- **`FOC_LOCALNET_CONTRACT_SIMPLE`** - Simple service address
  - Defaults to `0x0000000000000000000000000000000000000000` if not set
  - Used for service registry

- **`FOC_LOCALNET_CONTRACT_USDFC`** - USDFC token contract address
  - Defaults to `0x0000000000000000000000000000000000000000` if not set
  - Used for USDFC token operations

### Network Parameters

The following network parameters can be configured via environment variables. If not set, default values will be used:

- **`FOC_LOCALNET_BLOCK_DELAY`** - Block delay in seconds (default: `4`)
  - Controls the time between blocks in the network

- **`FOC_LOCALNET_PROPAGATION_DELAY`** - Propagation delay in seconds (default: `1`)
  - Network message propagation delay

- **`FOC_LOCALNET_EQUIVOCATION_DELAY`** - Equivocation delay in seconds (default: `0`)
  - Time delay for equivocation checks

## Example Configuration

```bash
# Required contract addresses
export FOC_LOCALNET_CONTRACT_PAY="0x1234567890123456789012345678901234567890"
export FOC_LOCALNET_CONTRACT_FWSS="0x2345678901234567890123456789012345678901"
export FOC_LOCALNET_CONTRACT_MULTICALL="0x3456789012345678901234567890123456789012"

# Optional contract addresses
export FOC_LOCALNET_CONTRACT_SIMPLE="0x4567890123456789012345678901234567890123"
export FOC_LOCALNET_CONTRACT_USDFC="0x5678901234567890123456789012345678901234"

# Optional network parameters (these values are the defaults)
export FOC_LOCALNET_BLOCK_DELAY="4"
export FOC_LOCALNET_PROPAGATION_DELAY="1"
export FOC_LOCALNET_EQUIVOCATION_DELAY="0"
```

## Error Handling

If any required environment variable is not set or contains an invalid value, Curio will panic at startup with a descriptive error message indicating which variable needs to be configured.

For contract addresses, the value must be a valid Ethereum hex address (starting with `0x` followed by 40 hexadecimal characters).

For network parameters, values must be valid unsigned 64-bit integers.
