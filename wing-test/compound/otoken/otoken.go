package compound

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

var ToAddres = "ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ"

//init
func OTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OWBTC)
	params := []interface{}{"init", []interface{}{}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//ATqpnrgVjzmkeHEqPiErnsxTEgi5goor2e
func OTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OWBTC)
	toAddress, _ := utils.AddressFromBase58(ToAddres)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 1000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
func OWBTCTokenTransfer(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OWBTC)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 1000000000}}
	mutTx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := sdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}

}

//ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ

func WingTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}

}
func OETHTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OETH)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(OTokenTransfer(cfg, account, genSdk))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
}

func OUSDTTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
}

func ApproveOToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.OETH)
	toAddress, _ := utils.AddressFromHexString(toAddrees)
	params := []interface{}{"approve", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
}

func BalanceOfOToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	BalanceAddr, _ := utils.AddressFromBase58(toAddrees)
	TokenAddr, _ := utils.AddressFromHexString(cfg.OETH)
	params := []interface{}{"balanceOf", []interface{}{BalanceAddr}}
	result, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(TokenAddr, params)
	log.Infof("result: %s", result.Result)

}

func TransferAllTestToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddrees string) {
	WingTokenTransfer(cfg, account, sdk, toAddrees)
	OUSDTTokenTransfer(cfg, account, sdk, toAddrees)
	OWBTCTokenTransfer(cfg, account, sdk, toAddrees)
	OETHTokenTransfer(cfg, account, sdk, toAddrees)
}
