package comptroller

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	flashTools "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"math/big"
)

type CmpControllAdmin struct {
	Cfg     *config.Config
	Account *goSdk.Account
	GenSdk  *goSdk.OntologySdk
}

func NewCmpControllAdmin() *CmpControllAdmin {
	configPath := "/home/ubuntu/go/src/github.com/mockyz/AutoTestGo/wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk := goSdk.NewOntologySdk()
	sdk.SetDefaultClient(rpcClient)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	return &CmpControllAdmin{
		Cfg:     cfg,
		Account: account,
		GenSdk:  sdk,
	}
}

func (self *CmpControllAdmin) init(cmpContractAddress, ownerAddr, globalParamAddr string) {
	cmpContractAddr, _ := utils.AddressFromHexString(cmpContractAddress)
	GlobalAddress, _ := utils.AddressFromHexString(globalParamAddr)
	OwnerAddress, _ := utils.AddressFromHexString(ownerAddr)
	params := []interface{}{OwnerAddress, GlobalAddress}
	mutTx, err := self.GenSdk.WasmVM.NewInvokeWasmVmTransaction(self.Cfg.GasPrice, self.Cfg.GasLimit,
		cmpContractAddr, "init", params)
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

func (self *CmpControllAdmin) SetPendingAdmin(contractAddrHex, adminAddrHex string) {
	adminAddr, _ := utils.AddressFromHexString(adminAddrHex)
	params := []interface{}{adminAddr}
	self.SendTx(contractAddrHex, "_setPendingAdmin", params)
}

func (self *CmpControllAdmin) Acceptadmin(contractAddrHex, pendingAdminAddrHex string) {
	addrx, _ := utils.AddressFromHexString(pendingAdminAddrHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_acceptAdmin", params)
}
func (self *CmpControllAdmin) Setglobalparam(contractAddrHex, globalParamAddrHex string) {
	addrx, _ := utils.AddressFromHexString(globalParamAddrHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_setGlobalParam", params)
}
func (self *CmpControllAdmin) Setpriceoracle(contractAddrHex, oracleHex string) {
	addrx, _ := utils.AddressFromHexString(oracleHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_setPriceOracle", params)
}
func (self *CmpControllAdmin) Setwingaddr(contractAddrHex, wingHex string) {
	addrx, _ := utils.AddressFromHexString(wingHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_setWingAddr", params)
}
func (self *CmpControllAdmin) Setclosefactor(contractAddrHex string, factor big.Int) {
	params := []interface{}{factor}
	self.SendTx(contractAddrHex, "_setCloseFactor", params)
}
func (self *CmpControllAdmin) Setinsurancerepayfactor(contractAddrHex string,
	insuranceRepayFactor big.Int) {
	params := []interface{}{insuranceRepayFactor}
	self.SendTx(contractAddrHex, "_setInsuranceRepayFactor", params)
}
func (self *CmpControllAdmin) Setmaxassets(contractAddrHex string, newMaxAssets big.Int) {
	params := []interface{}{newMaxAssets}
	self.SendTx(contractAddrHex, "_setMaxAssets", params)
}
func (self *CmpControllAdmin) Setliquidationincentive(contractAddrHex string, factor big.Int) {
	params := []interface{}{factor}
	self.SendTx(contractAddrHex, "_setLiquidationIncentive", params)
}
func (self *CmpControllAdmin) Supportmarket(contractAddrHex, marketHex string) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_supportMarket", params)
}
func (self *CmpControllAdmin) Setpauseguardian(contractAddrHex, guardianHex string) {
	addrx, _ := utils.AddressFromHexString(guardianHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "_setPauseGuardian", params)
}
func (self *CmpControllAdmin) Setmintpaused(contractAddrHex, marketHex string, state bool) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx, state}
	self.SendTx(contractAddrHex, "_setMintPaused", params)
}
func (self *CmpControllAdmin) Setborrowpaused(contractAddrHex, marketHex string, state bool) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx, state}
	self.SendTx(contractAddrHex, "_setBorrowPaused", params)
}
func (self *CmpControllAdmin) Settransferpaused(contractAddrHex string, state bool) {
	params := []interface{}{state}
	self.SendTx(contractAddrHex, "_setTransferPaused", params)
}
func (self *CmpControllAdmin) Setseizepaused(contractAddrHex, marketHex string, state bool) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx, state}
	self.SendTx(contractAddrHex, "_setSeizePaused", params)
}
func (self *CmpControllAdmin) Setwingrate(contractAddrHex string, wing_rate big.Int) {
	params := []interface{}{wing_rate}
	self.SendTx(contractAddrHex, "_setWingRate", params)
}
func (self *CmpControllAdmin) Setwingsblportion(contractAddrHex, marketHex string,
	supplyPortion, borrowPortion, insurancePortion big.Int) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx, supplyPortion, borrowPortion, insurancePortion}
	self.SendTx(contractAddrHex, "_setWingSBLPortion", params)
}
func (self *CmpControllAdmin) Updatemarketwingweight(contractAddrHex, marketHex string, weight big.Int,
) {
	addrx, _ := utils.AddressFromHexString(marketHex)
	params := []interface{}{addrx, weight}
	self.SendTx(contractAddrHex, "_updateMarketWingWeight", params)
}
func (self *CmpControllAdmin) Addwingmarkets(contractAddrHex string,
	markets []string, weights []big.Int) {
	var marketsAddresses = []common.Address{}
	for i, market := range markets {
		fmt.Println(i)
		addrx, _ := utils.AddressFromHexString(market)
		marketsAddresses = append(marketsAddresses, addrx)
	}
	params := []interface{}{marketsAddresses, weights}
	self.SendTx(contractAddrHex, "_addWingMarkets", params)
}
func (self *CmpControllAdmin) Dropwingmarket(contractAddrHex string, markets []string, weights []big.Int) {
	var marketsAddresses = []common.Address{}
	for i, market := range markets {
		fmt.Println(i)
		addrx, _ := utils.AddressFromHexString(market)
		marketsAddresses = append(marketsAddresses, addrx)
	}
	params := []interface{}{marketsAddresses, weights}
	self.SendTx(contractAddrHex, "_dropWingMarket", params)
}
func (self *CmpControllAdmin) Repaybyinsurance(contractAddrHex, borrowerAddrHex string) {
	addrx, _ := utils.AddressFromHexString(borrowerAddrHex)
	params := []interface{}{addrx}
	self.SendTx(contractAddrHex, "repayByInsurance", params)
}

func (self *CmpControllAdmin) SendTx(contractAddrHex, methodName string, params []interface{}) {
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

func (self *CmpControllAdmin) SendPreExecuteTx(contractAddrHex, methodName string, params []interface{}) {
	contractAddr, _ := utils.AddressFromHexString(contractAddrHex)
	preExecResult, _ := self.GenSdk.WasmVM.PreExecInvokeWasmVMContract(
		contractAddr, methodName, params)

	log.Infof("execute result: %s", preExecResult.Result)
}
