package wingGov

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"time"
)

//ZeroPoolWithDraw
func ZeroPoolWithDraw(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "withdraw_wing", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err2 := genSdk.SendTransaction(mutTx)
	if err2 != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	Utils.PrintSmartEventByHash_Ont(genSdk, hash.ToHexString())
}

//TODO : check method
func ZeroPoolGetUndlying(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.ResultItem {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(ZeroPoolAddr, "get_unbound_wing", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("get_unbound_wing :%s", result.Result)
	return result.Result
}

//TODO : GetStaking
func ZeroPoolGetStakingBalance(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.ResultItem {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address}
	result, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(ZeroPoolAddr, "get_staking_balance", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("get_unbound_wing :%s", result.Result)
	return result.Result
}

func ZeroPoolInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, zeroPoolAddr string) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(zeroPoolAddr)
	Global, _ := utils.AddressFromHexString(cfg.GlobalParam)
	GovToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{Global, GovToken}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err2 := genSdk.SendTransaction(mutTx)
	if err2 != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	log.Infof("txhash: %s", hash.ToHexString())
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(genSdk, hash.ToHexString())
}
func ZeroPoolStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, StakeOnt int) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address, StakeOnt}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "staking", params)
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	txhash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send staking tx: %s", txhash.ToHexString())
		time.Sleep(time.Second * 3)
		Utils.PrintSmartEventByHash_Ont(genSdk, txhash.ToHexString())
	}
}
func ZeroPoolStakingByAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, StakeOnt int, zeroPooAddr string) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(zeroPooAddr)
	params := []interface{}{account.Address, StakeOnt}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "staking", params)
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	txhash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send staking tx: %s", txhash.ToHexString())
		time.Sleep(time.Second * 3)
		Utils.PrintSmartEventByHash_Ont(genSdk, txhash.ToHexString())
	}
}
func ZeroPoolUnStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, StakeOnt int) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address, StakeOnt}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "unstaking", params)
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	time.Sleep(time.Second)
	txhash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send staking tx %s****", txhash.ToHexString())
		Utils.PrintSmartEventByHash_Ont(genSdk, txhash.ToHexString())
	}
}

func ZeroPoolUnStakingByAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, StakeOnt int, zeroPoolAddr string) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(zeroPoolAddr)
	params := []interface{}{account.Address, StakeOnt}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "unstaking", params)
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	time.Sleep(time.Second)
	txhash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send staking tx %s****", txhash.ToHexString())
		Utils.PrintSmartEventByHash_Ont(genSdk, txhash.ToHexString())
	}
}
func ZeroPooWithdrawWing(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{account.Address, cfg.StakeOnt}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "withdraw_wing", params)
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	time.Sleep(time.Second)
	txhash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send tx failed, err: %s********", err)
	} else {
		log.Infof("send staking tx %s****", txhash.ToHexString())
		Utils.PrintSmartEventByHash_Ont(genSdk, txhash.ToHexString())
	}
}

func MigrateZeroPool(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, codePath string) string {
	//time.Sleep(time.Second*3)
	newContractString := ContractMigrate(cfg, account, sdk, cfg.ZeroPool, codePath)
	log.Infof("new zero pool2 : %s", newContractString)
	time.Sleep(time.Second * 3)
	//	update pool address
	hash, err := sdk.SendTransaction(UpdatePoolAddress(cfg, account, sdk, newContractString, cfg.ZeroPool))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return ""
	}
	log.Infof("UpdatePoolAddress hash : %s", hash.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
	return newContractString
}

//efd78c612b66c690a59721b7bdd1c0e090c52ec4

func MigrateComptroller(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, oldCompaddr, newCompAddr string) {

	hash, err := sdk.SendTransaction(UpdatePoolAddress(cfg, account, sdk, newCompAddr, oldCompaddr))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	log.Infof("UpdatePoolAddress hash : %s", hash.ToHexString())
	Utils.PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
}
