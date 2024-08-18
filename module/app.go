package module

import (
	chain "github.com/wetee-dao/go-sdk"
	"github.com/wetee-dao/go-sdk/core"
	"github.com/wetee-dao/go-sdk/pallet/types"
	"github.com/wetee-dao/go-sdk/pallet/weteeapp"

	"errors"
)

// Worker
type App struct {
	Client *chain.ChainClient
	Signer *core.Signer
}

func (w *App) GetApp(publey []byte, id uint64) (*types.TeeApp, error) {
	if len(publey) != 32 {
		return nil, errors.New("publey length error")
	}

	var mss [32]byte
	copy(mss[:], publey)

	res, ok, err := weteeapp.GetTEEAppsLatest(w.Client.Api.RPC.State, mss, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetK8sClusterAccountsLatest => not start")
	}
	return &res, nil
}

func (w *App) GetAccount(id uint64) ([]byte, error) {
	res, ok, err := weteeapp.GetAppIdAccountsLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return res[:], nil
}

func (w *App) GetVersionLatest(id uint64) (uint64, error) {
	res, ok, err := weteeapp.GetAppVersionLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return uint64(res), nil
}

func (w *App) GetSecretEnv(id uint64) ([]byte, error) {
	res, ok, err := weteeapp.GetSecretEnvsLatest(w.Client.Api.RPC.State, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("GetAppIdAccountsLatest => not ok")
	}
	return res, nil
}
