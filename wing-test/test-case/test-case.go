package test_case

import (
	"github.com/mockyz/AutoTestGo/wing-test/compound/otoken"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math/big"
	"sync"
	"time"
)

func (this *TestRunner) WingSpeed4BorrowTest(userAddr string) {
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.Market.GetAddr()}
	totalBorrow, err := this.Market.TotalBorrows()
	if err != nil {
		log.Errorf(" Market TotalBorrows err : %v", err)
	}
	log.Infof("totalBorrow: %v", totalBorrow)
	accountSnap, err := this.Market.AccountSnapshot(addr)
	if err != nil {
		log.Errorf("  AccountSnapshot err : %v", err)
	}
	uesrBorrow := accountSnap.BorrowBalance
	log.Infof("AccountSnapshot , userBorrow: %v,TokenBalance:%v,ExchangeRate:%v", uesrBorrow, accountSnap.TokenBalance, accountSnap.ExchangeRate)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := GetClaimWingEventByHash(this.OntSDk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.Market.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Market.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.OntSDk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := GetClaimWingEventByHash(this.OntSDk, txhash)
	expRsult := ExpTestRuslt(wingSpeed, uesrBorrow, totalBorrow, claimStates0.Timestamp, claimStates1.Timestamp, 45)
	wingBalance1, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedBorrowerWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(notifyRsult)))
	log.Infof("CmpTestRuslt rate: %v", CmpTestRuslt(expRsult, notifyRsult))
}

func (this *TestRunner) WingSpeed4SuppluyTest(userAddr string) {
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddrs := []common.Address{this.Market.GetAddr()}
	total, err := this.Market.TotalSupply()
	if err != nil {
		log.Errorf(" Market TotalBorrows err : %v", err)
	}
	log.Infof("totalBorrow: %v", total)
	accountSnap, err := this.Market.AccountSnapshot(addr)
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
	claimStates0 := GetClaimWingEventByHash(this.OntSDk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(this.Market.GetAddr())
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Market.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.OntSDk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := GetClaimWingEventByHash(this.OntSDk, txhash)
	expRsult := ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 45)
	wingBalance1, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedSupplierWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("CmpTestRuslt rate: %v", CmpTestRuslt(expRsult, relRsult))
}

func (this *TestRunner) WingSpeed4InsuranceTest(userAddr string) {
	addr, _ := common.AddressFromBase58(userAddr)
	marketAddr, _ := common.AddressFromHexString("5daf48bd1aac8bc24e2b98ba426578e76f9abfcc")
	marketAddrs := []common.Address{marketAddr}
	total, err := this.Market.TotalSupply()
	if err != nil {
		log.Errorf(" Market Insurance TotalSupply err : %v", err)
	}
	log.Infof("Insurance TotalSupply: %v", total)
	user, err := this.Market.BalanceOf(addr)
	if err != nil {
		log.Errorf("  Insurance balance of  err : %v", err)
	}
	log.Infof("accountInsuranceBlance , user: %v", user)
	txhash0, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash0, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates0 := GetClaimWingEventByHash(this.OntSDk, txhash0)
	wingBalance0, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	log.Infof("wingBalance0: %v", wingBalance0)
	wingSpeed, err := this.Comptroller.WingSpeeds(marketAddr)
	if err != nil {
		log.Errorf(" WingSpeeds Comptroller err : %v", err)
	}
	marketName, _ := this.Market.Name()
	log.Infof("marketName: %s, speeds: %v", marketName, wingSpeed)
	//wait time
	this.OntSDk.WaitForGenerateBlock(time.Minute*5, 1000)
	txhash, _, remian, err := this.Comptroller.ClaimWingAtMarkets(addr, marketAddrs, false)
	log.Infof("txhash:%v,remian:%v", txhash, remian)
	if err != nil {
		log.Errorf("  ClaimWingAtMarkets err : %v", err)
	}
	claimStates1 := GetClaimWingEventByHash(this.OntSDk, txhash)
	wingBalance1, err := otoken.BalanceOfOToken2(this.OntSDk, this.TestConfig.WING, userAddr)
	if err != nil {
		log.Errorf(" wingBalance0 balance err : %v", err)
	}
	expRsult := ExpTestRuslt(wingSpeed, user, total, claimStates0.Timestamp, claimStates1.Timestamp, 10)
	log.Infof("wingBalance1: %v", wingBalance1)
	relRsult := big.NewInt(0).Sub(wingBalance1, wingBalance0)
	notifyRsult := claimStates1.DistributedGuaranteeWing
	log.Infof("relRsult: %v,expRsult: %v,notifyRsult: %v", relRsult, expRsult, notifyRsult)
	log.Infof("notifyRsult sub relRsult: %v", big.NewInt(0).Sub(notifyRsult, relRsult))
	log.Infof("errRate : %v", new(big.Float).Quo(new(big.Float).SetInt(expRsult), new(big.Float).SetInt(relRsult)))
	log.Infof("CmpTestRuslt rate: %v", CmpTestRuslt(expRsult, relRsult))
}

func TestWingSpeeds(marketName, usrAddr string, sy *sync.WaitGroup) {
	defer sy.Done()
	tr, err := NewTestRunner("../config.json")
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	tr.WingSpeed4BorrowTest(usrAddr)
	tr.WingSpeed4SuppluyTest(usrAddr)
}
