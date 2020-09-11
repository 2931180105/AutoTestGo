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

type CmpControllReader struct {
	Cfg     *config.Config
	Account *goSdk.Account
	GenSdk  *goSdk.OntologySdk
}

func NewCmpControllReader() *CmpControllReader {
	configPath := "/home/ubuntu/go/src/github.com/mockyz/AutoTestGo/wing-test/config_testnet.json"
	cfg, _ := config.ParseConfig(configPath)
	//account,_ := Utils.NewAccountByWif("L1nfGvz19cWXHDLeEMMC6vozhSLANCy9E2gNxh3YwHJMXReLddNw")
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk := goSdk.NewOntologySdk()
	sdk.SetDefaultClient(rpcClient)
	wallet, _ := sdk.OpenWallet(cfg.Wallet)
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	return &CmpControllReader{
		Cfg:     cfg,
		Account: account,
		GenSdk:  sdk,
	}
}

func (self *CmpControllReader) SendTx(contractAddrHex, methodName string, params []interface{}) {
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

func (self *CmpControllReader) SendPreExecuteTx(contractAddrHex, methodName string, params []interface{}) {
	contractAddr, _ := utils.AddressFromHexString(contractAddrHex)
	preExecResult, _ := self.GenSdk.WasmVM.PreExecInvokeWasmVMContract(
		contractAddr, methodName, params)

	log.Infof("execute result: %s", preExecResult.Result)
}

func (self *CmpControllReader) Mintallowed(contractAddrHex, fTokenHex, minterHex string) {
	addrx1, _ := utils.AddressFromHexString(fTokenHex)
	addrx2, _ := utils.AddressFromHexString(minterHex)
	params := []interface{}{addrx1, addrx2}
	self.SendPreExecuteTx(contractAddrHex, "mintAllowed", params)
}
func (self *CmpControllReader) Redeemallowed(contractAddrHex, fTokenHex, redeemer string, redeemTokens big.Int) {
	addrx1, _ := utils.AddressFromHexString(fTokenHex)
	addrx2, _ := utils.AddressFromHexString(redeemer)
	params := []interface{}{addrx1, addrx2, redeemTokens}
	self.SendPreExecuteTx(contractAddrHex, "redeemAllowed", params)
}
func (self *CmpControllReader) Borrowallowed(contractAddrHex, fTokenHex,
	borrowerhex string, borrowAmount big.Int) {
	addrx1, _ := utils.AddressFromHexString(fTokenHex)
	addrx2, _ := utils.AddressFromHexString(borrowerhex)
	params := []interface{}{addrx1, addrx2, borrowAmount}
	self.SendPreExecuteTx(contractAddrHex, "borrowAllowed", params)
}
func (self *CmpControllReader) Repayborrowallowed(contractAddrHex, fToken, payer, borrower string, repayAmount big.Int) {
	addrx1, _ := utils.AddressFromHexString(fToken)
	addrx2, _ := utils.AddressFromHexString(payer)
	addrx3, _ := utils.AddressFromHexString(borrower)
	params := []interface{}{addrx1, addrx2, addrx3, repayAmount}
	self.SendPreExecuteTx(contractAddrHex, "repayBorrowAllowed", params)
}
func (self *CmpControllReader) Liquidateborrowallowed(contractAddrHex, fTokenBorrowed,
	fTokenCollateral, liquidator, borrower string, repayAmount big.Int) {
	addrx1, _ := utils.AddressFromHexString(fTokenBorrowed)
	addrx2, _ := utils.AddressFromHexString(fTokenCollateral)
	addrx3, _ := utils.AddressFromHexString(liquidator)
	addrx4, _ := utils.AddressFromHexString(borrower)
	params := []interface{}{addrx1, addrx2, addrx3, addrx4, repayAmount}
	self.SendPreExecuteTx(contractAddrHex, "liquidateBorrowAllowed", params)
}
func (self *CmpControllReader) SeizeAllowed(contractAddrHex, fTokenBorrowed,
	fTokenCollateral, liquidator, borrower string, seizeTokens big.Int) {
	addrx1, _ := utils.AddressFromHexString(fTokenBorrowed)
	addrx2, _ := utils.AddressFromHexString(fTokenCollateral)
	addrx3, _ := utils.AddressFromHexString(liquidator)
	addrx4, _ := utils.AddressFromHexString(borrower)
	params := []interface{}{addrx1, addrx2, addrx3, addrx4, seizeTokens}
	self.SendPreExecuteTx(contractAddrHex, "seizeAllowed", params)
}
func (self *CmpControllReader) Transferallowed(contractAddrHex, fToken, src, dst string, transferTokens big.Int) {
	addrx1, _ := utils.AddressFromHexString(fToken)
	addrx2, _ := utils.AddressFromHexString(src)
	addrx3, _ := utils.AddressFromHexString(dst)
	params := []interface{}{addrx1, addrx2, addrx3, transferTokens}
	self.SendPreExecuteTx(contractAddrHex, "transferAllowed", params)
}
func (self *CmpControllReader) Assetsin(contractAddrHex, account string) {
	addrx, _ := utils.AddressFromHexString(account)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "assetsIn", params)
}
func (self *CmpControllReader) Checkmembership(contractAddrHex, account, fToken string) {
	addrx1, _ := utils.AddressFromHexString(account)
	addrx2, _ := utils.AddressFromHexString(fToken)
	params := []interface{}{addrx1, addrx2}
	self.SendPreExecuteTx(contractAddrHex, "checkMembership", params)
}
func (self *CmpControllReader) Mintverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "mintVerify", params)
}
func (self *CmpControllReader) Redeemverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "redeemVerify", params)
}
func (self *CmpControllReader) Borrowverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "borrowVerify", params)
}
func (self *CmpControllReader) Repayborrowverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "repayBorrowVerify", params)
}
func (self *CmpControllReader) Liquidateborrowverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "liquidateBorrowVerify", params)
}
func (self *CmpControllReader) Seizeverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "seizeVerify", params)
}
func (self *CmpControllReader) Transferverify(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "transferVerify", params)
}
func (self *CmpControllReader) Getaccountliquidity(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "getAccountLiquidity", params)
}
func (self *CmpControllReader) Gethypotheticalaccountliquidity(contractAddrHex, addrHex, oTokenModify, redeemTokens string, borrowAmount big.Int) {
	addrx1, _ := utils.AddressFromHexString(addrHex)
	addrx2, _ := utils.AddressFromHexString(oTokenModify)
	addrx3, _ := utils.AddressFromHexString(redeemTokens)
	params := []interface{}{addrx1, addrx2, addrx3, borrowAmount}
	self.SendPreExecuteTx(contractAddrHex, "getHypotheticalAccountLiquidity", params)
}
func (self *CmpControllReader) Liquidatecalculateseizetokens(contractAddrHex, fTokenCollateral, fTokenBorrowed string,
	actualRepayAmount big.Int) {
	addrx1, _ := utils.AddressFromHexString(fTokenCollateral)
	addrx2, _ := utils.AddressFromHexString(fTokenBorrowed)
	params := []interface{}{addrx1, addrx2, actualRepayAmount}
	self.SendPreExecuteTx(contractAddrHex, "liquidateCalculateSeizeTokens", params)
}
func (self *CmpControllReader) Admin(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "admin", params)
}
func (self *CmpControllReader) Pendingadmin(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "pendingAdmin", params)
}
func (self *CmpControllReader) Oracle(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "oracle", params)
}

func (self *CmpControllReader) GlobalParam(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "globalParam", params)
}
func (self *CmpControllReader) Closefactormantissa(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "closeFactorMantissa", params)
}
func (self *CmpControllReader) Insurancerepayfactormantissa(contractAddrHex, addrHex string) {
	addrx, _ := utils.AddressFromHexString(addrHex)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "insuranceRepayFactorMantissa", params)
}
func (self *CmpControllReader) Couldrepaybyinsurance(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "couldRepayByInsurance", params)
}
func (self *CmpControllReader) Liquidationincentivemantissa(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "liquidationIncentiveMantissa", params)
}
func (self *CmpControllReader) Maxassets(contractAddrHex string, maxAssets big.Int) {
	params := []interface{}{maxAssets}
	self.SendPreExecuteTx(contractAddrHex, "maxAssets", params)
}
func (self *CmpControllReader) Marketmeta(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "marketMeta", params)
}
func (self *CmpControllReader) Pauseguardian(contractAddrHex, guardian string) {
	addrx, _ := utils.AddressFromHexString(guardian)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "pauseGuardian", params)
}
func (self *CmpControllReader) Transferguardianpaused(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "transferGuardianPaused", params)
}
func (self *CmpControllReader) Seizeguardianpaused(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "seizeGuardianPaused", params)
}
func (self *CmpControllReader) Mintguardianpaused(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "mintGuardianPaused", params)
}
func (self *CmpControllReader) Borrowguardianpaused(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "borrowGuardianPaused", params)
}
func (self *CmpControllReader) Allmarkets(contractAddrHex string, markets []string) {
	var marketsAddresses = []common.Address{}
	for i, market := range markets {
		fmt.Println(i)
		addrx, _ := utils.AddressFromHexString(market)
		marketsAddresses = append(marketsAddresses, addrx)
	}
	params := []interface{}{marketsAddresses}
	self.SendPreExecuteTx(contractAddrHex, "allMarkets", params)
}
func (self *CmpControllReader) Ismarketexisted(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "isMarketExisted", params)
}
func (self *CmpControllReader) Wingdistributednum(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "wingDistributedNum", params)
}
func (self *CmpControllReader) Wingaddr(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "wingAddr", params)
}
func (self *CmpControllReader) Wingrate(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "wingRate", params)
}
func (self *CmpControllReader) Wingspeeds(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "wingSpeeds", params)
}
func (self *CmpControllReader) Wingsblportion(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "wingSBLPortion", params)
}
func (self *CmpControllReader) Wingsupplystate(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "wingSupplyState", params)
}
func (self *CmpControllReader) Wingborrowstate(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "wingBorrowState", params)
}
func (self *CmpControllReader) Winginsurancestate(contractAddrHex, market string) {
	addrx, _ := utils.AddressFromHexString(market)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "wingInsuranceState", params)
}
func (self *CmpControllReader) Wingsupplierindex(contractAddrHex, market, account string) {
	addrx1, _ := utils.AddressFromHexString(market)
	addrx2, _ := utils.AddressFromHexString(account)
	params := []interface{}{addrx1, addrx2}
	self.SendPreExecuteTx(contractAddrHex, "wingSupplierIndex", params)
}
func (self *CmpControllReader) Wingborrowerindex(contractAddrHex, market, account string) {
	addrx1, _ := utils.AddressFromHexString(market)
	addrx2, _ := utils.AddressFromHexString(account)
	params := []interface{}{addrx1, addrx2}
	self.SendPreExecuteTx(contractAddrHex, "wingBorrowerIndex", params)
}
func (self *CmpControllReader) Winginsuranceindex(contractAddrHex, market, account string) {
	addrx1, _ := utils.AddressFromHexString(market)
	addrx2, _ := utils.AddressFromHexString(account)
	params := []interface{}{addrx1, addrx2}
	self.SendPreExecuteTx(contractAddrHex, "wingInsuranceIndex", params)
}
func (self *CmpControllReader) Wingaccrued(contractAddrHex, account string) {
	addrx, _ := utils.AddressFromHexString(account)
	params := []interface{}{addrx}
	self.SendPreExecuteTx(contractAddrHex, "wingAccrued", params)
}
func (self *CmpControllReader) Iscomptroller(contractAddrHex string) {
	params := []interface{}{}
	self.SendPreExecuteTx(contractAddrHex, "isComptroller", params)
}
