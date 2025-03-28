package main

import (
	"fmt"
	"math/big"
	"runtime"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	chain "github.com/wetee-dao/go-sdk"
	"github.com/wetee-dao/go-sdk/pallet/balances"
	gtypes "github.com/wetee-dao/go-sdk/pallet/types"
)

func main() {
	fmt.Println(signature.TestKeyringPairAlice.Address)

	client, err := chain.ClientInit("wss://xiaobai.asyou.me:30011/ws", true)
	if err != nil {
		panic(err)
	}

	var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
	p, err := chain.Ed25519PairFromSecret(testSecretSeed, 42)
	if err != nil {
		panic(err)
	}

	fmt.Println("p.Address:", p.Address)

	minter, _ := types.NewMultiAddressFromAccountID(signature.TestKeyringPairAlice.PublicKey)
	minterWrap := gtypes.MultiAddress{
		IsId:       true,
		AsIdField0: minter.AsID,
	}
	bal, _ := new(big.Int).SetString("500000000000", 10)
	call := balances.MakeTransferAllowDeathCall(minterWrap, types.NewUCompact(bal))
	err = client.SignAndSubmit(&p, call, true)
	if err != nil {
		printErrorStack(err)
		return
	}

	printErrorStack(err)
}

func printErrorStack(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if !ok {
			fmt.Printf("runtime.Caller error\n")
			return
		}
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("Error: %s\nFile: %s\nLine: %d\nFunction: %s\n",
			err.Error(), file, line, funcName)
	}
}
