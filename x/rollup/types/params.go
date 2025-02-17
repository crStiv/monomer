package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	const (
		defaultL1FeeRecipient string = "0x000000000000000000000000000000000000dEaD"
		// defaultL1CrossDomainMessenger uses the devnet address of the L1 cross domain messenger contract as the default value.
		defaultL1CrossDomainMessenger string = "0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE"
		defaultMinFeeWithdrawalAmount uint64 = 400_000
		defaultFeeWithdrawalGasLimit  uint64 = 400_000
	)

	return Params{
		L1FeeRecipient:         defaultL1FeeRecipient,
		L1CrossDomainMessenger: defaultL1CrossDomainMessenger,
		MinFeeWithdrawalAmount: defaultMinFeeWithdrawalAmount,
		FeeWithdrawalGasLimit:  defaultFeeWithdrawalGasLimit,
	}
}

// Validate checks that the parameters have valid values.
func (p *Params) Validate() error {
	if err := validateEthAddress(p.L1FeeRecipient); err != nil {
		return fmt.Errorf("validate L1 fee recipient address: %w", err)
	}
	if err := validateEthAddress(p.L1CrossDomainMessenger); err != nil {
		return fmt.Errorf("validate L1 cross domain messenger address: %w", err)
	}

	return nil
}

func validateEthAddress(addr string) error {
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("validate ethereum address: %s", addr)
	}
	return nil
}
