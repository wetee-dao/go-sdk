package weteeapp

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

// See [`Pallet::create`].
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []uint32, cpu4 uint32, memory5 uint32, disk6 uint32, level7 byte, deposit8 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeApp: true,
		AsWeteeAppField0: &types1.WeteeAppPalletCall{
			IsCreate:         true,
			AsCreateName0:    name0,
			AsCreateImage1:   image1,
			AsCreateMeta2:    meta2,
			AsCreatePort3:    port3,
			AsCreateCpu4:     cpu4,
			AsCreateMemory5:  memory5,
			AsCreateDisk6:    disk6,
			AsCreateLevel7:   level7,
			AsCreateDeposit8: deposit8,
		},
	}
}

// See [`Pallet::update`].
func MakeUpdateCall(appId0 uint64, name1 []byte, image2 []byte, port3 []uint32) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeApp: true,
		AsWeteeAppField0: &types1.WeteeAppPalletCall{
			IsUpdate:       true,
			AsUpdateAppId0: appId0,
			AsUpdateName1:  name1,
			AsUpdateImage2: image2,
			AsUpdatePort3:  port3,
		},
	}
}

// See [`Pallet::set_settings`].
func MakeSetSettingsCall(appId0 uint64, value1 []types1.AppSettingInput) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeApp: true,
		AsWeteeAppField0: &types1.WeteeAppPalletCall{
			IsSetSettings:       true,
			AsSetSettingsAppId0: appId0,
			AsSetSettingsValue1: value1,
		},
	}
}

// See [`Pallet::recharge`].
func MakeRechargeCall(id0 uint64, deposit1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeApp: true,
		AsWeteeAppField0: &types1.WeteeAppPalletCall{
			IsRecharge:         true,
			AsRechargeId0:      id0,
			AsRechargeDeposit1: deposit1,
		},
	}
}

// See [`Pallet::restart`].
func MakeRestartCall(appId0 uint64) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeApp: true,
		AsWeteeAppField0: &types1.WeteeAppPalletCall{
			IsRestart:       true,
			AsRestartAppId0: appId0,
		},
	}
}
