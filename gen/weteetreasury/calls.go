package weteetreasury

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

func MakeSpendCall(daoId0 uint64, beneficiary1 [32]byte, amount2 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeTEETreasury: true,
		AsWeTEETreasuryField0: &types1.WeteeTreasuryPalletCall{
			IsSpend:             true,
			AsSpendDaoId0:       daoId0,
			AsSpendBeneficiary1: beneficiary1,
			AsSpendAmount2:      amount2,
		},
	}
}
