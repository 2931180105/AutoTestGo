package Interest

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	flashTools "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	"math/big"
)

type InterestReader struct {
	Cfg     *config.Config
	Account *goSdk.Account
	GenSdk  *goSdk.OntologySdk
}

func NewInterestReader() *InterestReader {
	configPath := "wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk := goSdk.NewOntologySdk()
	sdk.SetDefaultClient(rpcClient)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	return &InterestReader{
		Cfg:     cfg,
		Account: account,
		GenSdk:  sdk,
	}
}

func (self *InterestReader) SendTx(contractAddrHex, methodName string, params []interface{}) {
	contractAddr, _ := utils.AddressFromHexString(contractAddrHex)
	mutTx, err := self.GenSdk.WasmVM.NewInvokeWasmVmTransaction(self.Cfg.GasPrice, self.Cfg.GasLimit,
		contractAddr, methodName, params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := flashTools.SignTx(self.GenSdk, mutTx, self.Cfg.StartNonce, self.Account); err != nil {
		log.Error(err)
	}
	hash, err2 := self.GenSdk.SendTransaction(mutTx)
	if err2 != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	log.Infof("txhash: %s", hash.ToHexString())
}

func (self *InterestReader) SendPreExecuteTx(contractAddrHex, methodName string, params []interface{}) {
	contractAddr, _ := utils.AddressFromHexString(contractAddrHex)
	preExecResult, _ := self.GenSdk.WasmVM.PreExecInvokeWasmVMContract(
		contractAddr, methodName, params)

	log.Infof("execute result: %s", preExecResult.Result.ToString())
}

func (self *InterestReader) IsInterestRateModel(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "isInterestRateModel", params)
}
func (self *InterestReader) UtilizationRate(contractAddrHex string, cash, borrows, reserves big.Int) {
	params := []interface{}{cash, borrows, reserves}
	self.SendPreExecuteTx(contractAddrHex, "utilizationRate", params)
}
func (self *InterestReader) BorrowRate(contractAddrHex string, cash, borrows, reserves big.Int) {
	params := []interface{}{cash, borrows, reserves}
	self.SendPreExecuteTx(contractAddrHex, "borrowRate", params)
}
func (self *InterestReader) SupplyRate(contractAddrHex string, cash, borrows, reserves, reservesFactor,
	insuranceFactor big.Int) {
	params := []interface{}{cash, borrows, reserves, reservesFactor, insuranceFactor}
	self.SendPreExecuteTx(contractAddrHex, "supplyRate", params)
}
func (self *InterestReader) BaseRatePerBlock(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "baseRatePerBlock", params)
}
func (self *InterestReader) MultiplierPerBlock(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "multiplierPerBlock", params)
}
