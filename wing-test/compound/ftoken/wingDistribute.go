package ftoken

import (
	"github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math/big"
	"strings"
	"sync"
	"time"
)

func (this *FlashToken) WingSpeed4SuppluyTest(userAddr string) {
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.GetAddr()}
	total, err := this.TotalSupply()
	if err != nil {
		log.Errorf(" Market TotalBorrows err : %v", err)
	}
	log.Infof("totalBorrow: %v", total)
	accountSnap, err := this.AccountSnapshot(addr)
	if err != nil {
		log.Errorf("  AccountSnapshot err : %v", err)
	}
	user := accountSnap.TokenBalance
	log.Infof("AccountSnapshot , userBorrow: %v,TokenBalance:%v,ExchangeRate:%v", user, accountSnap.TokenBalance, accountSnap.ExchangeRate)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := utils.GetClaimWingEventByHash(this.sdk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.sdk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := utils.GetClaimWingEventByHash(this.sdk, txhash)
	expRsult := utils.ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 50)
	wingBalance1, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedSupplierWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("utils.CmpTestRuslt rate: %v", utils.CmpTestRuslt(expRsult, relRsult))
}
//no use
func (this *FlashToken) WingSpeed4InsuranceTest(userAddr string) {
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.GetAddr()}
	total, err := this.TotalSupply()
	if err != nil {
		log.Errorf(" Market Insurance TotalSupply err : %v", err)
	}
	log.Infof("Insurance TotalSupply: %v", total)
	user, err := this.BalanceOf(addr)
	if err != nil {
		log.Errorf("  Insurance balance of  err : %v", err)
	}
	log.Infof("accountInsuranceBlance , user: %v", user)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := utils.GetClaimWingEventByHash(this.sdk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.sdk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := utils.GetClaimWingEventByHash(this.sdk, txhash)
	wingBalance1, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	expRsult := utils.ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 10)
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedGuaranteeWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("utils.CmpTestRuslt rate: %v", utils.CmpTestRuslt(expRsult, relRsult))
}


func (this *FlashToken) WingSpeed4SuppluyTestNew(marketAddr common.Address ,userAddr string, sy *sync.WaitGroup)  {
	defer sy.Done()
	this.SetAddr(marketAddr)
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.GetAddr()}
	total, err := this.TotalSupply()
	if err != nil {
		log.Errorf(" Market TotalBorrows err : %v", err)
	}
	log.Infof("totalBorrow: %v", total)
	accountSnap, err := this.AccountSnapshot(addr)
	if err != nil {
		log.Errorf("  AccountSnapshot err : %v", err)
	}
	user := accountSnap.TokenBalance
	log.Infof("AccountSnapshot , accountSnap.TokenBalance: %v,TokenBalance:%v,ExchangeRate:%v", user, accountSnap.TokenBalance, accountSnap.ExchangeRate)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := utils.GetClaimWingEventByHash(this.sdk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.sdk.WaitForGenerateBlock(time.Minute*15, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := utils.GetClaimWingEventByHash(this.sdk, txhash)
	expRsult := utils.ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 50)
	wingBalance1, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedSupplierWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("utils.CmpTestRuslt rate: %v", utils.CmpTestRuslt(expRsult, relRsult))
	utils.InsertWingDisResToDb4Supply(marketName,userAddr,this.addr.ToHexString(),user,total,wingSpeed,expRsult,notifyRsult,claimStates0.Timestamp,claimStates1.Timestamp,utils.CmpTestRuslt(expRsult, notifyRsult))
}


func (this *FlashToken) WingSpeed4SuppluyTestByName(marketName  ,userAddr string) {
	//通过marketName 获取ftoken地址
	ftokenAddressList, err :=  this.Comptroller.AllMarkets()
	for _, ftokenAddress := range ftokenAddressList {
		this.SetAddr(ftokenAddress)
		mn, err :=this.Name()
		if err !=nil{
			return
		}
		if strings.Contains(mn, marketName) {
			break
		}
	}
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.GetAddr()}
	total, err := this.TotalSupply()
	if err != nil {
		log.Errorf(" Market TotalBorrows err : %v", err)
	}
	log.Infof("totalBorrow: %v", total)
	accountSnap, err := this.AccountSnapshot(addr)
	if err != nil {
		log.Errorf("  AccountSnapshot err : %v", err)
	}
	user := accountSnap.TokenBalance
	log.Infof("AccountSnapshot , userBorrow: %v,TokenBalance:%v,ExchangeRate:%v", user, accountSnap.TokenBalance, accountSnap.ExchangeRate)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := utils.GetClaimWingEventByHash(this.sdk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.sdk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := utils.GetClaimWingEventByHash(this.sdk, txhash)
	expRsult := utils.ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 50)
	wingBalance1, err := otoken.BalanceOfOToken2(this.sdk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedSupplierWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("utils.CmpTestRuslt rate: %v", utils.CmpTestRuslt(expRsult, relRsult))
	utils.InsertWingDisResToDb4Supply(marketName,userAddr,this.addr.ToHexString(),user,total,wingSpeed,expRsult,notifyRsult,claimStates0.Timestamp,claimStates1.Timestamp,utils.CmpTestRuslt(expRsult, notifyRsult))
}
//func (this *FlashToken) TestWingSpeeds4Supply(marketAddr common.Address, usrAddr string, sy *sync.WaitGroup) {
//	defer sy.Done()
//	this.WingSpeed4SuppluyTestNew(marketAddr, usrAddr)
//}