package test_case

import (
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/compound/ftoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"math/big"
)

type TestRunner struct {
	Comptroller      *comptroller.Comptroller
	Market           *ftoken.FlashToken
	OntSDk           *ontSDK.OntologySdk
	Account 		 *ontSDK.Account
	TestConfig 		 *config.Config
}

type ClaimWingAtMarket struct {
	DistributedBorrowerWing  *big.Int
	DistributedSupplierWing  *big.Int
	DistributedGuaranteeWing *big.Int
	Timestamp                uint32
}
type ClaimStates struct {
	DistributedType        string
	DistributedAddr        string
	DistributedToAddr      string
	DistributedAmount      *big.Int
	DistributedTotalAmount *big.Int
}
var sdk = ontSDK.NewOntologySdk()

func NewTestRunner(cfgFile string) (*TestRunner, error) {
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	wallet, err := ontSDK.OpenWallet(cfg.Wallet)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[1])
	sdk.SetDefaultClient(rpcClient)
	comp,err := comptroller.NewComptroller(sdk,cfg.Comptroller,account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil {
		log.Errorf("NewComptroller  err:%v",err)
	}
	market ,err := ftoken.NewFlashToken(sdk,cfg.FETH,account,cfg.GasPrice,cfg.GasLimit)
	return &TestRunner{Comptroller:comp,Market:market,OntSDk:sdk,Account:account,TestConfig:cfg},nil
}

func NewTestRunner2(cfg *config.Config,account *ontSDK.Account,sdk *ontSDK.OntologySdk, makretAddr string) (*TestRunner, error) {
	comp,err := comptroller.NewComptroller(sdk,cfg.Comptroller,account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil {
		log.Errorf("NewComptroller  err:%v",err)
	}
	market ,err := ftoken.NewFlashToken(sdk,makretAddr,account,cfg.GasPrice,cfg.GasLimit)
	return &TestRunner{Comptroller:comp,Market:market,OntSDk:sdk,Account:account,TestConfig:cfg},nil
}


func NewMarkets(cfg *config.Config, account *ontSDK.Account, sdk *ontSDK.OntologySdk, makretAddr string) (*ftoken.FlashToken, error) {
	comp,err := comptroller.NewComptroller(sdk,cfg.Comptroller,account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil {
		log.Errorf("NewComptroller  err:%v",err)
	}
	market ,err := ftoken.NewFlashToken2(sdk,makretAddr,account,cfg,comp)
	if err !=nil {
		log.Errorf("NewFlashToken  err:%v",err)
	}
	market.TestConfig = cfg

	return market, nil
}

