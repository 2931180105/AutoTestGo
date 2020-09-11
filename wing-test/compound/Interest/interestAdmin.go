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

type InterestAdmin struct {
	Cfg     *config.Config
	Account *goSdk.Account
	GenSdk  *goSdk.OntologySdk
}

func NewInterestAdmin() *InterestAdmin {
	configPath := "wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk := goSdk.NewOntologySdk()
	sdk.SetDefaultClient(rpcClient)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	return &InterestAdmin{
		Cfg:     cfg,
		Account: account,
		GenSdk:  sdk,
	}
}

func (self *InterestAdmin) init(ContractAddress string, baseRatePerYear, multiplierPerYear big.Int) {
	ContractAddr, _ := utils.AddressFromHexString(ContractAddress)
	params := []interface{}{baseRatePerYear, multiplierPerYear}
	mutTx, err := self.GenSdk.WasmVM.NewInvokeWasmVmTransaction(self.Cfg.GasPrice, self.Cfg.GasLimit,
		ContractAddr, "init", params)
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
