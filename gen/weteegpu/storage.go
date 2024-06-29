package weteegpu

import (
	"encoding/hex"
	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	codec "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	types1 "github.com/wetee-dao/go-sdk/gen/types"
)

// Make a storage key for NextTeeId id={{false [12]}}
//
//	The id of the next app to be created.
//	获取下一个app id
func MakeNextTeeIdStorageKey() (types.StorageKey, error) {
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "NextTeeId")
}

var NextTeeIdResultDefaultBytes, _ = hex.DecodeString("0000000000000000")

func GetNextTeeId(state state.State, bhash types.Hash) (ret uint64, err error) {
	key, err := MakeNextTeeIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextTeeIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}
func GetNextTeeIdLatest(state state.State) (ret uint64, err error) {
	key, err := MakeNextTeeIdStorageKey()
	if err != nil {
		return
	}
	var isSome bool
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	if !isSome {
		err = codec.Decode(NextTeeIdResultDefaultBytes, &ret)
		if err != nil {
			return
		}
	}
	return
}

// Make a storage key for GPUApps
//
//	App
//	应用
func MakeGPUAppsStorageKey(tupleOfByteArray32Uint640 [32]byte, tupleOfByteArray32Uint641 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfByteArray32Uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfByteArray32Uint641)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "GPUApps", byteArgs...)
}
func GetGPUApps(state state.State, bhash types.Hash, tupleOfByteArray32Uint640 [32]byte, tupleOfByteArray32Uint641 uint64) (ret types1.GpuApp, isSome bool, err error) {
	key, err := MakeGPUAppsStorageKey(tupleOfByteArray32Uint640, tupleOfByteArray32Uint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetGPUAppsLatest(state state.State, tupleOfByteArray32Uint640 [32]byte, tupleOfByteArray32Uint641 uint64) (ret types1.GpuApp, isSome bool, err error) {
	key, err := MakeGPUAppsStorageKey(tupleOfByteArray32Uint640, tupleOfByteArray32Uint641)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Prices
//
//	Price of resource
//	价格
func MakePricesStorageKey(byte0 byte) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(byte0)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "Prices", byteArgs...)
}
func GetPrices(state state.State, bhash types.Hash, byte0 byte) (ret types1.Price2, isSome bool, err error) {
	key, err := MakePricesStorageKey(byte0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetPricesLatest(state state.State, byte0 byte) (ret types1.Price2, isSome bool, err error) {
	key, err := MakePricesStorageKey(byte0)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for AppIdAccounts
//
//	App 拥有者账户
//	user's K8sCluster information
func MakeAppIdAccountsStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "AppIdAccounts", byteArgs...)
}
func GetAppIdAccounts(state state.State, bhash types.Hash, uint640 uint64) (ret [32]byte, isSome bool, err error) {
	key, err := MakeAppIdAccountsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAppIdAccountsLatest(state state.State, uint640 uint64) (ret [32]byte, isSome bool, err error) {
	key, err := MakeAppIdAccountsStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for Envs
//
//	App setting
//	App设置
func MakeEnvsStorageKey(tupleOfUint64Uint160 uint64, tupleOfUint64Uint161 uint16) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(tupleOfUint64Uint160)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	encBytes, err = codec.Encode(tupleOfUint64Uint161)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "Envs", byteArgs...)
}
func GetEnvs(state state.State, bhash types.Hash, tupleOfUint64Uint160 uint64, tupleOfUint64Uint161 uint16) (ret types1.Env, isSome bool, err error) {
	key, err := MakeEnvsStorageKey(tupleOfUint64Uint160, tupleOfUint64Uint161)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetEnvsLatest(state state.State, tupleOfUint64Uint160 uint64, tupleOfUint64Uint161 uint16) (ret types1.Env, isSome bool, err error) {
	key, err := MakeEnvsStorageKey(tupleOfUint64Uint160, tupleOfUint64Uint161)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}

// Make a storage key for AppVersion
//
//	App version
//	App 版本
func MakeAppVersionStorageKey(uint640 uint64) (types.StorageKey, error) {
	byteArgs := [][]byte{}
	encBytes := []byte{}
	var err error
	encBytes, err = codec.Encode(uint640)
	if err != nil {
		return nil, err
	}
	byteArgs = append(byteArgs, encBytes)
	return types.CreateStorageKey(&types1.Meta, "WeTEEGpu", "AppVersion", byteArgs...)
}
func GetAppVersion(state state.State, bhash types.Hash, uint640 uint64) (ret uint64, isSome bool, err error) {
	key, err := MakeAppVersionStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorage(key, &ret, bhash)
	if err != nil {
		return
	}
	return
}
func GetAppVersionLatest(state state.State, uint640 uint64) (ret uint64, isSome bool, err error) {
	key, err := MakeAppVersionStorageKey(uint640)
	if err != nil {
		return
	}
	isSome, err = state.GetStorageLatest(key, &ret)
	if err != nil {
		return
	}
	return
}
