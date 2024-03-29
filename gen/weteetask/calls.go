package weteetask

import (
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

// See [`Pallet::create`].
func MakeCreateCall(name0 []byte, image1 []byte, port2 []uint32, cpu3 uint32, memory4 uint32, disk5 uint32, level6 byte, deposit7 types.UCompact) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types1.WeteeTaskPalletCall{
			IsCreate:         true,
			AsCreateName0:    name0,
			AsCreateImage1:   image1,
			AsCreatePort2:    port2,
			AsCreateCpu3:     cpu3,
			AsCreateMemory4:  memory4,
			AsCreateDisk5:    disk5,
			AsCreateLevel6:   level6,
			AsCreateDeposit7: deposit7,
		},
	}
}

// See [`Pallet::rerun`].
func MakeRerunCall(id0 uint64) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types1.WeteeTaskPalletCall{
			IsRerun:    true,
			AsRerunId0: id0,
		},
	}
}

// See [`Pallet::update`].
func MakeUpdateCall(appId0 uint64, name1 []byte, image2 []byte, port3 []uint32) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types1.WeteeTaskPalletCall{
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
		IsWeteeTask: true,
		AsWeteeTaskField0: &types1.WeteeTaskPalletCall{
			IsSetSettings:       true,
			AsSetSettingsAppId0: appId0,
			AsSetSettingsValue1: value1,
		},
	}
}

// See [`Pallet::recharge`].
func MakeRechargeCall(id0 uint64, deposit1 types.U128) types1.RuntimeCall {
	return types1.RuntimeCall{
		IsWeteeTask: true,
		AsWeteeTaskField0: &types1.WeteeTaskPalletCall{
			IsRecharge:         true,
			AsRechargeId0:      id0,
			AsRechargeDeposit1: deposit1,
		},
	}
}
