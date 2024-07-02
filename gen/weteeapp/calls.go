package weteeapp

import types "github.com/wetee-dao/go-sdk/gen/types"

// App create
// 注册任务
func MakeCreateCall(name0 []byte, image1 []byte, meta2 []byte, port3 []types.Service, command4 types.Command, env5 []types.EnvInput, cpu6 uint32, memory7 uint32, disk8 []types.Disk, sideContainer9 []types.Container, level10 byte, teeVersion11 types.TEEVersion) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEApp: true,
		AsWeTEEAppField0: &types.WeteeAppPalletCall{
			IsCreate:               true,
			AsCreateName0:          name0,
			AsCreateImage1:         image1,
			AsCreateMeta2:          meta2,
			AsCreatePort3:          port3,
			AsCreateCommand4:       command4,
			AsCreateEnv5:           env5,
			AsCreateCpu6:           cpu6,
			AsCreateMemory7:        memory7,
			AsCreateDisk8:          disk8,
			AsCreateSideContainer9: sideContainer9,
			AsCreateLevel10:        level10,
			AsCreateTeeVersion11:   teeVersion11,
		},
	}
}

// App update
// 更新任务
func MakeUpdateCall(appId0 uint64, newName1 types.OptionTByteSlice, newImage2 types.OptionTByteSlice, newPort3 types.OptionTServiceSlice, newCommand4 types.OptionTCommand, newEnv5 []types.EnvInput, withRestart6 bool) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEApp: true,
		AsWeTEEAppField0: &types.WeteeAppPalletCall{
			IsUpdate:             true,
			AsUpdateAppId0:       appId0,
			AsUpdateNewName1:     newName1,
			AsUpdateNewImage2:    newImage2,
			AsUpdateNewPort3:     newPort3,
			AsUpdateNewCommand4:  newCommand4,
			AsUpdateNewEnv5:      newEnv5,
			AsUpdateWithRestart6: withRestart6,
		},
	}
}

// App restart
// 更新任务
func MakeRestartCall(appId0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEApp: true,
		AsWeTEEAppField0: &types.WeteeAppPalletCall{
			IsRestart:       true,
			AsRestartAppId0: appId0,
		},
	}
}

// update price
// 更新价格
func MakeUpdatePriceCall(level0 byte, price1 types.Price) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEApp: true,
		AsWeTEEAppField0: &types.WeteeAppPalletCall{
			IsUpdatePrice:       true,
			AsUpdatePriceLevel0: level0,
			AsUpdatePricePrice1: price1,
		},
	}
}
