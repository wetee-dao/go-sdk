package client

import (
	"fmt"
	"hash"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"github.com/pkg/errors"

	gtypes "github.com/wetee-dao/go-sdk/gen/types"
)

// 区块链链接
// chain client
type ChainClient struct {
	Api     *gsrpc.SubstrateAPI
	Meta    *types.Metadata
	Hash    types.Hash
	Runtime *types.RuntimeVersion
}

// 初始化区块连链接
// init chain client
func ClientInit() (*ChainClient, error) {
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		return nil, err
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}

	runtime, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}

	return &ChainClient{api, meta, genesisHash, runtime}, nil
}

// 获取区块高度
// get block number
func (c *ChainClient) GetBlockNumber() (types.BlockNumber, error) {
	hash, err := c.Api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return 0, err
	}
	return hash.Number, nil
}

// 获取账户信息
// get account info
func (c *ChainClient) GetAccount(address *signature.KeyringPair) (*types.AccountInfo, error) {
	key, err := types.CreateStorageKey(c.Meta, "System", "Account", address.PublicKey)
	if err != nil {
		panic(err)
	}
	var accountInfo types.AccountInfo
	_, err = c.Api.RPC.State.GetStorageLatest(key, &accountInfo)
	return &accountInfo, err
}

// 签名并提交交易
// sign and submit transaction
func (c *ChainClient) SignAndSubmit(signer *signature.KeyringPair, runtimeCall gtypes.RuntimeCall, untilFinalized bool) error {
	accountInfo, err := c.GetAccount(signer)
	if err != nil {
		return err
	}
	call, err := (runtimeCall).AsCall()
	if err != nil {
		return err
	}

	ext := types.NewExtrinsic(call)
	era := types.ExtrinsicEra{IsMortalEra: false}
	nonce := uint32(accountInfo.Nonce)

	o := types.SignatureOptions{
		BlockHash:          c.Hash,
		Era:                era,
		GenesisHash:        c.Hash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        c.Runtime.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.Runtime.TransactionVersion,
	}

	err = ext.Sign(*signer, o)
	if err != nil {
		return err
	}

	sub, err := c.Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return err
	}

	defer sub.Unsubscribe()
	timeout := time.After(20 * time.Second)
	for {
		select {
		case status := <-sub.Chan():
			// fmt.Printf("%#v\n", status)
			if status.IsInBlock {
				LogWithRed("SubmitAndWatchExtrinsic", " => IsInBlock")
				if !untilFinalized {
					return nil
				}
			}
			if status.IsFinalized {
				LogWithRed("SubmitAndWatchExtrinsic", " => IsFinalized")
				return nil
			}
		case err := <-sub.Err():
			LogWithRed("SubmitAndWatchExtrinsic ERROR", err.Error())
			return err
		case <-timeout:
			fmt.Println("timeout")
			return nil
		}
	}
}

// query map data list
func (c *ChainClient) QueryMapAll(pallet string, method string) ([]types.StorageChangeSet, error) {
	key := createPrefixedKey(pallet, method)

	keys, err := c.Api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return []types.StorageChangeSet{}, errors.Wrap(err, "[GetKeysLatest]")
	}

	set, err := c.Api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return []types.StorageChangeSet{}, errors.Wrap(err, "[QueryStorageAtLatest]")
	}

	return set, nil
}

// query double map data list
func (c *ChainClient) QueryDoubleMapAll(pallet string, method string, keyarg interface{}, at *types.Hash) ([]types.StorageChangeSet, error) {
	key, err := c.GetDoubleMapPrefixKey(pallet, method, keyarg)
	if err != nil {
		return nil, errors.Wrap(err, "[GetDoubleMapPrefixKey]")
	}

	// query key
	var keys []types.StorageKey
	if at == nil {
		keys, err = c.Api.RPC.State.GetKeysLatest(key)
	} else {
		keys, err = c.Api.RPC.State.GetKeys(key, *at)
	}

	if err != nil {
		return nil, errors.Wrap(err, "[GetKeysLatest]")
	}

	// get all data
	var set []types.StorageChangeSet
	if at == nil {
		set, err = c.Api.RPC.State.QueryStorageAtLatest(keys)
	} else {
		set, err = c.Api.RPC.State.QueryStorageAt(keys, *at)
	}
	if err != nil {
		return nil, errors.Wrap(err, "[QueryStorageAtLatest]")
	}

	return set, nil
}

func (c *ChainClient) GetDoubleMapPrefixKey(pallet string, method string, keyarg interface{}) ([]byte, error) {
	arg, err := codec.Encode(keyarg)
	if err != nil {
		return nil, err
	}

	// create key prefix
	key := createPrefixedKey(pallet, method)
	hashers, err := c.GetHashers(pallet, method)
	if err != nil {
		return nil, err
	}

	// write key
	_, err = hashers[0].Write(arg)
	if err != nil {
		return nil, fmt.Errorf("unable to hash args[%d]: %s Error: %v", 0, arg, err)
	}
	// append hash to key
	key = append(key, hashers[0].Sum(nil)...)

	return key, nil
}

func (c *ChainClient) GetHashers(pallet, method string) ([]hash.Hash, error) {
	// get entry metadata
	// 获取储存元数据
	entryMeta, err := c.Meta.FindStorageEntryMetadata(pallet, method)
	if err != nil {
		return nil, err
	}

	// check if it's a map
	// 判断是否为map
	if !entryMeta.IsMap() {
		return nil, errors.New(pallet + "." + method + "is not map")
	}

	// get map hashers
	// 获取储存的 hasher 函数
	hashers, err := entryMeta.Hashers()
	if err != nil {
		return nil, errors.Wrap(err, "[Hashers]")
	}
	return hashers, nil
}

func createPrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
