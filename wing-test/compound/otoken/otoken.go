package otoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	WingUtils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"math/big"
	"time"
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

//DAI ETH 18
func OTokenTransfer(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddres, oToken string, precise uint64) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	toAddress, _ := utils.AddressFromBase58(toAddres)
	amount := WingUtils.ToIntByPrecise("100", precise)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, amount}}
	mutTx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, OTokenAddr, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTx(sdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := sdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********， token : %s", err, oToken)
		return
	} else {
		log.Infof("txhash: %s", hash1.ToHexString())
	}
	//time.Sleep(time.Second * 3)
	//WingUtils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
}

func WingTokenTransfer(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees string) {
	OTokenAddr, _ := utils.AddressFromHexString(cfg.GovToken)
	toAddress, _ := utils.AddressFromBase58(toAddrees)
	params := []interface{}{"transfer", []interface{}{account.Address, toAddress, 200000000000000}}
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
func OTokenDelegateToProxy(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, oToken string, precise uint64) {
	OTokenAddr, _ := utils.AddressFromHexString(oToken)
	amount := WingUtils.ToIntByPrecise("1000000", precise)
	params := []interface{}{"delegateToProxy", []interface{}{account.Address, amount}}
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
	time.Sleep(time.Second * 3)
	WingUtils.PrintSmartEventByHash_Ont(genSdk, hash1.ToHexString())
}

func ApproveOToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, toAddrees OntCommon.Address, oToken OntCommon.Address, amount *big.Int) {
	params := []interface{}{"approve", []interface{}{account.Address, toAddrees, amount}}
	mutTx, err := genSdk.NeoVM.NewNeoVMInvokeTransaction(cfg.GasPrice, cfg.GasLimit, oToken, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := WingUtils.SignTxAndSendTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
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

func BalanceOfOToken(goSdk *goSdk.OntologySdk, toAddrees, oToken string) {
	BalanceAddr, _ := utils.AddressFromBase58(toAddrees)
	TokenAddr, _ := utils.AddressFromHexString(oToken)
	params := []interface{}{"balanceOf", []interface{}{BalanceAddr}}
	result, err := goSdk.NeoVM.PreExecInvokeNeoVMContract(TokenAddr, params)
	if err != nil {
		log.Errorf("BalanceOfOToken:%v", err)
	}
	balance, err := result.Result.ToInteger()
	if err != nil {
		log.Errorf("BalanceOfOToken result.Result.ToInteger() :%v", err)
	}
	log.Infof("result: %v", balance)
}
func BalanceOfOToken2(goSdk *goSdk.OntologySdk, toAddrees, oToken string) (*big.Int, error ){
	BalanceAddr, _ := utils.AddressFromBase58(toAddrees)
	TokenAddr, _ := utils.AddressFromHexString(oToken)
	params := []interface{}{"balanceOf", []interface{}{BalanceAddr}}
	result, err := goSdk.NeoVM.PreExecInvokeNeoVMContract(TokenAddr, params)
	if err != nil {
		log.Errorf("BalanceOfOToken:%v", err)
	}
	balance, err := result.Result.ToInteger()
	if err != nil {
		log.Errorf("BalanceOfOToken result.Result.ToInteger() :%v", err)
	}
	log.Infof("result: %v", balance)
	return balance,nil
}

func TransferAllTestToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk, toAddrees string) {
	//WingTokenTransfer(cfg, account, sdk, toAddrees)
	//ToAddres, _ := utils.AddressFromBase58(toAddrees)
	//_, _ = sdk.Native.Ont.Transfer(cfg.GasPrice, cfg.GasLimit, account, account, ToAddres, 1)
	//_, _ = sdk.Native.Ong.Transfer(cfg.GasPrice, cfg.GasLimit, account, account, ToAddres, 100000000000)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.RENBTC, 8)
	log.Infof("toaddress: %s", toAddrees)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.ETH, 18)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.DAI, 18)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.OKB, 18)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.UNI, 18)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.SUSD, 18)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.GovToken, 9)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.WBTC, 8)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.ONTd, 9)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.USDC, 6)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.NEO, 8)
	OTokenTransfer(cfg, account, sdk, toAddrees, cfg.USDT, 6)

}
func BalanceOfAllToken(cfg *config.Config, goSdk *goSdk.OntologySdk, toAddrees string) {
	BalanceOfOToken(goSdk, toAddrees, cfg.WBTC)
	BalanceOfOToken(goSdk, toAddrees, cfg.ONTd)
	BalanceOfOToken(goSdk, toAddrees, cfg.RENBTC)
	BalanceOfOToken(goSdk, toAddrees, cfg.USDC)
	BalanceOfOToken(goSdk, toAddrees, cfg.WING)
	BalanceOfOToken(goSdk, toAddrees, cfg.ETH)
	BalanceOfOToken(goSdk, toAddrees, cfg.DAI)
	BalanceOfOToken(goSdk, toAddrees, cfg.USDT)
	BalanceOfOToken(goSdk, toAddrees, cfg.SUSD)
	BalanceOfOToken(goSdk, toAddrees, cfg.NEO)
	BalanceOfOToken(goSdk, toAddrees, cfg.OKB)
	BalanceOfOToken(goSdk, toAddrees, cfg.UNI)

}

func DelegateToProxyAllTestToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk) {
	//OTokenDelegateToProxy(cfg, account, sdk, cfg.ODAI, 18)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.WBTC, 8)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.ETH, 18)
	//OTokenDelegateToProxy(cfg, account, sdk, cfg.ETH9, 9)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.RENBTC, 8)
	//OTokenDelegateToProxy(cfg, account, sdk, cfg.ONTD, 9)
	OTokenDelegateToProxy(cfg, account, sdk, cfg.USDC, 6)
}

func GenerateAccountsToken(cfg *config.Config, admin *goSdk.Account, goSdk *goSdk.OntologySdk) {
	accounts := DbHelp.QueryAccountFromDb(900, cfg.AccountNum)
	for i := 0; i < cfg.AccountNum; i++ {
		TransferAllTestToken(cfg, admin, goSdk, accounts[i].Address.ToBase58())
		time.Sleep(time.Second/2)
	}
}
