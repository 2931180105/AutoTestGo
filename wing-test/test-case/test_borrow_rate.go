package test_case

import (
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"math/big"
	"sync"
	"time"
)

//use time to per rate
func TestBorrowRateNew(marketName string, sy *sync.WaitGroup) {
	defer sy.Done()
	testRunner, err := NewTestRunner("../config.json")
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	addr, _ := common.AddressFromBase58("AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p")
	//before test
	borrowRatePerBlock, err := testRunner.Market.BorrowRatePerBlock()
	supplyingRatePerBlock, _ := testRunner.Market.SupplyRatePerBlock()
	if err != nil {
		log.Errorf("BorrowRatePerBlock err : %v", err)
	}
	hash0, err := testRunner.Market.AccrueInterest()
	if err != nil {
		log.Errorf("  hash0 AccrueInterest err : %v marketName:%s", err, marketName)
		return
	}
	userUnderlying0, _ := testRunner.Market.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(testRunner.OntSDk, hash0)
	startTime := GetTimeByTxhash(testRunner.OntSDk, hash0)
	totalBorrow, _ := testRunner.Market.TotalBorrows()
	totalSupply, _ := testRunner.Market.TotalSupply()
	reserveFactor, _ := testRunner.Market.ReserveFactorMantissa()
	accountSnapshot, _ := testRunner.Market.AccountSnapshot(addr)
	userSupply := accountSnapshot.TokenBalance
	reserves, _ := testRunner.Market.TotalReserves()
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	time.Sleep(time.Second * 300)
	hash1, err := testRunner.Market.AccrueInterest()
	if err != nil {
		log.Errorf("  hash1 AccrueInterest err : %v marketName:%s", err, marketName)
		return
	}
	userUnderlying1, _ := testRunner.Market.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(testRunner.OntSDk, hash1)
	endTime := GetTimeByTxhash(testRunner.OntSDk, hash1)
	totalBorrow1, _ := testRunner.Market.TotalBorrows()
	totalSupply1, _ := testRunner.Market.TotalSupply()
	reserves1, _ := testRunner.Market.TotalReserves()
	log.Infof("marketName: %s, endTime:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v", marketName, endTime, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("marketName: %s, borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", marketName, borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("marketName: %s,totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", marketName, totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(1e9))
	log.Infof("marketName: %s,reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", marketName, reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Mul(big.NewInt(0).Div(userSupply, totalSupply), expSupplyInterestAdd)
	log.Infof("marketName: %s,userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v", marketName, userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}

func TestBorrowRateByBlock(marketName string) {
	testRunner, err := NewTestRunner("/home/ubuntu/go/src/github.com/mockyz/AutoTestGo/wing-test/config_prv.json")
	if err != nil {
		log.Errorf("NewTestRunner err : %v", err)
	}
	addr, _ := common.AddressFromBase58("AT9sH4s84NGJYVqNHQWN6vkgb7jQ12eR7p")
	//before test
	borrowRatePerBlock, err := testRunner.Market.BorrowRatePerBlock()
	supplyingRatePerBlock, _ := testRunner.Market.SupplyRatePerBlock()
	if err != nil {
		log.Errorf("BorrowRatePerBlock err : %v", err)
	}
	hash0, err := testRunner.Market.AccrueInterest()
	if err != nil {
		log.Errorf("  hash0 AccrueInterest err : %v marketName:%s", err, marketName)
		return
	}
	userUnderlying0, _ := testRunner.Market.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(testRunner.OntSDk, hash0)
	//startTime := GetTimeByTxhash(testRunner.OntSDk, hash0)
	startTime, _ := testRunner.OntSDk.GetBlockHeightByTxHash(hash0)

	totalBorrow, _ := testRunner.Market.TotalBorrows()
	totalSupply, _ := testRunner.Market.TotalSupply()
	reserveFactor, _ := testRunner.Market.ReserveFactorMantissa()
	accountSnapshot, _ := testRunner.Market.AccountSnapshot(addr)
	userSupply := accountSnapshot.TokenBalance
	reserves, _ := testRunner.Market.TotalReserves()
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	time.Sleep(time.Second * 300)
	hash1, err := testRunner.Market.AccrueInterest()
	if err != nil {
		log.Errorf("  hash1 AccrueInterest err : %v marketName:%s", err, marketName)
		return
	}
	userUnderlying1, _ := testRunner.Market.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(testRunner.OntSDk, hash1)
	endTime, _ := testRunner.OntSDk.GetBlockHeightByTxHash(hash1)
	//endTime := GetTimeByTxhash(testRunner.OntSDk, hash1)
	totalBorrow1, _ := testRunner.Market.TotalBorrows()
	totalSupply1, _ := testRunner.Market.TotalSupply()
	reserves1, _ := testRunner.Market.TotalReserves()
	log.Infof("marketName: %s, endTime:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v", marketName, endTime, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("marketName: %s, borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", marketName, borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("marketName: %s,totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", marketName, totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(1e9))
	log.Infof("marketName: %s,reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", marketName, reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Mul(big.NewInt(0).Div(userSupply, totalSupply), expSupplyInterestAdd)
	log.Infof("marketName: %s,userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v", marketName, userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}
