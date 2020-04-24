package main

import (
	"encoding/hex"
	"fmt"
	layer2_sdk "github.com/ontio/layer2-go-sdk"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
)

func newLayer2Account1(ontsdk *layer2_sdk.OntologySdk) (*layer2_sdk.Account, error) {
	privateKey, err := keypair.WIF2Key([]byte("L5CKUdMTnHQNeBtCzdoEZ1hyRpaCsc7LaesZWvFhfpKbzQV1v7pk"))
	if err != nil {
		return nil, fmt.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	return &layer2_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func newLayer2Account2(ontsdk *layer2_sdk.OntologySdk) (*layer2_sdk.Account, error) {
	// AScExXzLbkZV32tDFdV7Uoq7ZhCT1bRCGp
	privateKey, err := keypair.WIF2Key([]byte("KyxsqZ45MCx3t2UbuG9P8h96TzyrzTXGRQnfs9nZKFx6YkjTfHqb"))
	if err != nil {
		return nil, fmt.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	return &layer2_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func newOntologyAccount1(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
	// AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc
	privateKey, err := keypair.WIF2Key([]byte("L5CKUdMTnHQNeBtCzdoEZ1hyRpaCsc7LaesZWvFhfpKbzQV1v7pk"))
	if err != nil {
		return nil, fmt.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	return &ontology_go_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func newOntologyAccount2(ontsdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
	// AScExXzLbkZV32tDFdV7Uoq7ZhCT1bRCGp
	privateKey, err := keypair.WIF2Key([]byte("KyxsqZ45MCx3t2UbuG9P8h96TzyrzTXGRQnfs9nZKFx6YkjTfHqb"))
	if err != nil {
		return nil, fmt.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	fmt.Printf("address: %s\n", address.ToBase58())
	return &ontology_go_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func layer2DepositTransfer(ontsdk *layer2_sdk.OntologySdk, gasPrice, gasLimit uint64, payer *layer2_sdk.Account, to common.Address, amount uint64) (common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(gasPrice, gasLimit, common.ADDRESS_EMPTY, to, amount)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	if payer != nil {
		ontsdk.SetPayer(tx, payer.Address)
		err = ontsdk.SignToTransaction(tx, payer)
		if err != nil {
			return common.UINT256_EMPTY, err
		}
	}
	return ontsdk.SendTransaction(tx)
}

func layer2WithdrawTransfer(ontsdk *layer2_sdk.OntologySdk, gasPrice, gasLimit uint64, payer *layer2_sdk.Account, from common.Address, amount uint64) (common.Uint256, error) {
	tx, err := ontsdk.Native.Ong.NewTransferTransaction(gasPrice, gasLimit, from, common.ADDRESS_EMPTY, amount)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	if payer != nil {
		ontsdk.SetPayer(tx, payer.Address)
		err = ontsdk.SignToTransaction(tx, payer)
		if err != nil {
			return common.UINT256_EMPTY, err
		}
	}
	return ontsdk.SendTransaction(tx)
}

func getLayer2Balance(ontsdk *layer2_sdk.OntologySdk, addr common.Address) uint64 {
	amount, _ := ontsdk.Native.Ong.BalanceOf(addr)
	return amount
}

func readLayer2Account() {
	// create alliance sdk
	layer2sdk := layer2_sdk.NewOntologySdk()
	layer2sdk.NewRpcClient().SetAddress("http://localhost:40336")

	var wallet *layer2_sdk.Wallet
	var err error
	if !common.FileExisted("./wallet_layer2.dat") {
		wallet, err = layer2sdk.CreateWallet("./wallet_layer2.dat")
		if err != nil {
			return
		}
	} else {
		wallet, err = layer2sdk.OpenWallet("./wallet_layer2.dat")
		if err != nil {
			fmt.Errorf("NewETHManager - wallet open error: %s", err.Error())
			return
		}
	}

	signer, err := wallet.GetDefaultAccount([]byte("1"))
	if err != nil || signer == nil {
		signer, err = wallet.NewDefaultSettingAccount([]byte("1"))
		if err != nil {
			fmt.Errorf("NewETHManager - wallet password error")
			return
		}

		err = wallet.Save()
		if err != nil {
			return
		}
	}
	pri_key, _ := keypair.Key2WIF(signer.PrivateKey)
	addr := signer.Address.ToBase58()
	fmt.Printf("private key: %s, address: %s %s\n", string(pri_key), addr, signer.Address.ToHexString())
}

func readOntologyAccount() {
	// create alliance sdk
	ontsdk := ontology_go_sdk.NewOntologySdk()
	ontsdk.NewRpcClient().SetAddress("http://localhost:40336")

	var wallet *ontology_go_sdk.Wallet
	var err error
	if !common.FileExisted("./wallet_ontology.dat") {
		wallet, err = ontsdk.CreateWallet("./wallet_ontology.dat")
		if err != nil {
			return
		}
	} else {
		wallet, err = ontsdk.OpenWallet("./wallet_ontology.dat")
		if err != nil {
			fmt.Errorf("NewETHManager - wallet open error: %s", err.Error())
			return
		}
	}

	signer, err := wallet.GetDefaultAccount([]byte("1"))
	if err != nil || signer == nil {
		signer, err = wallet.NewDefaultSettingAccount([]byte("1"))
		if err != nil {
			fmt.Errorf("NewETHManager - wallet password error")
			return
		}

		err = wallet.Save()
		if err != nil {
			return
		}
	}
	pri_key, _ := keypair.Key2WIF(signer.PrivateKey)
	addr := signer.Address.ToBase58()
	fmt.Printf("private key: %s, address: %s %s\n", string(pri_key), addr, signer.Address.ToHexString())
}

func layer2Deposit() {
	// create alliance sdk
	layer2_sdk := layer2_sdk.NewOntologySdk()
	layer2_sdk.NewRpcClient().SetAddress("http://localhost:40336")

	//
	account_operator, _ := newLayer2Account1(layer2_sdk)
	account_user, _ := newLayer2Account2(layer2_sdk)

	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %d is: %d\n", account_user.Address.ToBase58(), balance)

	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 10000000
	{
		txhash, err := layer2DepositTransfer(layer2_sdk, gasPrice, gasLimit, account_operator, account_user.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %d is: %d\n", account_user.Address.ToBase58(), balance)
}

func layer2Withdraw() {
	// create alliance sdk
	layer2_sdk := layer2_sdk.NewOntologySdk()
	layer2_sdk.NewRpcClient().SetAddress("http://172.168.3.59:40336")

	//
	account_user, _ := newLayer2Account2(layer2_sdk)

	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %d is: %d\n", account_user.Address.ToBase58(), balance)

	//
	var gasPrice uint64 = 0
	var gasLimit uint64 = 20000
	var amount uint64 = 2000000
	{
		txhash, err := layer2WithdrawTransfer(layer2_sdk, gasPrice, gasLimit, account_user, account_user.Address, amount)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		} else {
			fmt.Printf("tx hash: %s\n", txhash.ToHexString())
		}
	}
	//
	balance = getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %d is: %d\n", account_user.Address.ToBase58(), balance)
}

func layer2Balance() {
	// create alliance sdk
	layer2_sdk := layer2_sdk.NewOntologySdk()
	layer2_sdk.NewRpcClient().SetAddress("http://172.168.3.59:40336")

	account_user, _ := newLayer2Account2(layer2_sdk)

	//
	balance := getLayer2Balance(layer2_sdk, account_user.Address)
	fmt.Printf("amount of address %d is: %d\n", account_user.Address.ToBase58(), balance)
}

func getLayer2State() {
	// create alliance sdk
	height := uint32(21)
	layer2_sdk := layer2_sdk.NewOntologySdk()
	layer2_sdk.NewRpcClient().SetAddress("http://localhost:40336")
	layer2State, pks, _ := layer2_sdk.GetLayer2State(height)
	fmt.Printf("layer2state, state root: %s \n", layer2State.StatesRoot.ToHexString())
	for _, pk := range pks {
		addr := types.AddressFromPubKey(pk)
		fmt.Printf(" %s ", addr.ToBase58())
	}
}

func commitLayer2State2Ontology() {
	testOntSdk := ontology_go_sdk.NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress("http://polaris5.ont.io:20336")

	contractAddress, _ := common.AddressFromHexString("4229a92d90d446d1598e12e35698b681ae4d4642")
	depositids := make([]int, 0)
	for i := 0; i < 2; i++ {
		depositids = append(depositids, 3+i)
	}
	withdrawAmounts := make([]uint64, 0)
	toAddresses := make([]common.Address, 0)
	assetAddress := make([][]byte, 0)
	for i := 0; i < 1; i++ {
		withdrawAmounts = append(withdrawAmounts, 3)
		toAddress, _ := common.AddressFromBase58("AMUGPqbVJ3TG6pe7xRpxxaeh4ai4fu9ahc")
		toAddresses = append(toAddresses, toAddress)
		tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
		assetAddress = append(assetAddress, tokenAddress)
	}
	tx, err := testOntSdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contractAddress, []interface{}{"updateState", []interface{}{
		"0000000000000000000000000000000000000000000000000000000000000000", 6, "1.0.0",
		depositids, []interface{}{}, []interface{}{}, []interface{}{}}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	account_operator, err := newOntologyAccount1(testOntSdk)
	if err != nil {
		fmt.Printf("newOntologyAccount failed!")
	}
	testOntSdk.SetPayer(tx, account_operator.Address)
	err = testOntSdk.SignToTransaction(tx, account_operator)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := testOntSdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func ontologyDeposit() {
	testOntSdk := ontology_go_sdk.NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress("http://polaris5.ont.io:20336")

	contractAddress, _ := common.AddressFromHexString("d1b62355d7c88a76fefee8f4ce14efb477992a3c")
	account_user, err := newOntologyAccount2(testOntSdk)
	if err != nil {
		fmt.Printf("ontology account err: %s", err.Error())
		return
	}
	tokenAddress, _ := hex.DecodeString("0000000000000000000000000000000000000002")
	tx, err := testOntSdk.NeoVM.NewNeoVMInvokeTransaction(500, 40000, contractAddress, []interface{}{"deposit", []interface{}{
		account_user.Address, 3000000000, tokenAddress}})
	if err != nil {
		fmt.Printf("new transaction failed!")
	}
	testOntSdk.SetPayer(tx, account_user.Address)
	err = testOntSdk.SignToTransaction(tx, account_user)
	if err != nil {
		fmt.Printf("SignToTransaction failed!")
	}
	txHash, err := testOntSdk.SendTransaction(tx)
	if err != nil {
		fmt.Printf("SignToTransaction failed! err: %s", err.Error())
	}
	fmt.Printf("hash: %s", txHash.ToHexString())
}

func main() {
	cmd := 101
	switch cmd {
	case 100:
		layer2Balance()
	case 101:
		layer2Deposit()
	case 102:
		layer2Withdraw()
	case 103:
		getLayer2State()
	case 200:
		ontologyDeposit()
	case 201:
		commitLayer2State2Ontology()
	case 300:
		readLayer2Account()
	case 301:
		readOntologyAccount()
	}
}
