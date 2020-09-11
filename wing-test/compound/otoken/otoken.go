package otoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	WingUtils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/core/types"
)

//var ToAddres = "ANxSSzWmFnAtqWBtq2KthP73oX4bHf9FyZ"

//init
func OTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, oToken string) *types.MutableTransaction {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	params := []interface{}{"init", []interface{}{}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

func OTokenTransfer(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddres, oToken string) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	toAddress, _ := utils.AddressFromBase58(toAddres)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 1000000000}}
	mutTx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
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

//
//func OWBTCTokenTransfer(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddrees string, oToken string) {
//	OTokenAddr, _ := utils.AddressFromHexString(oToken)
//	toAddress, _ := utils.AddressFromBase58(toAddrees)
//	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 1000000000}}
//	mutTx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
//	if err != nil {
//		fmt.Println("construct tx err", err)
//	}
//	if err := WingUtils.SignTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
//		log.Error(err)
//	}
//	hash1, err := sdk.SendTransaction(mutTx)
//	if err != nil {
//		log.Errorf("send  tx failed, err: %s********", err)
//		return
//	} else {
//		log.Infof("txhash: %s", hash1.ToHexString())
//	}
//
//}

func WingTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}

}
func OTokenDelegateToProxy(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, oToken string) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	params := []interface{}{"delegateToProxy", []interface{}{account.Address, 1000000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
}

//
//func OUSDTTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string, oToken string) {
//	OTokenAddr, _ := utils.AddressFromHexString(oToken)
//	toAddress, _ := utils.AddressFromBase58(toAddrees)
//	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 100000000000}}
//	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
//	if err != nil {
//		fmt.Println("construct tx err", err)
//	}
//	if err := WingUtils.SignTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
//		log.Error(err)
//	}
//	hash1, err := genSdk.SendTransaction(mutTx)
//	if err != nil {
//		log.Errorf("send  tx failed, err: %s********", err)
//	} else {
//		log.Infof("txhash: %s", hash1.ToHexString())
//	}
//}

func ApproveOToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string, oToken string) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	toAddress, _ := utils.AddressFromHexString(toAddrees)
	params := []interface{}{"approve", []interface{}{account.Address, toAddress, 100000000000}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
}

func AllowanceToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, owner, spender string, oToken string) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	ownerAddr, _ := utils.AddressFromBase58(owner)
	spenderAddr, _ := utils.AddressFromBase58(spender)
	params := []interface{}{"allowance", []interface{}{ownerAddr, spenderAddr}}
	result, err := genSdk.NeoVM.PreExecInvokeNeoVMContract(OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("owner address: %s, spender address: %s, allowance tokenAddr: %s, amount:%s", owner, spender, oToken, result.Result)
}

func BalanceOfOToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees, oToken string) {
	BalanceAddr, _ := utils.AddressFromBase58(toAddrees)
	TokenAddr, _ := utils.AddressFromHexString(oToken)
	params := []interface{}{"balanceOf", []interface{}{BalanceAddr}}
	result, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(TokenAddr, params)
	log.Infof("result: %s", result.Result)

}

func TransferAllTestToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddrees string) {
	WingTokenTransfer(cfg, account, sdk, toAddrees)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.OUSDT)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.OWBTC)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.OETH)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.ODAI)
}

func DelegateToProxyAllTestToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk) {
	OTokenDelegateToProxy(cfg, account, sdk, cfg.OUSDT)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.OWBTC)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.OETH)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.ODAI)
}
