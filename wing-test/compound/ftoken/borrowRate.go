package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"math/big"
	"time"
)

func (this *FlashToken) TestBorrowRateByBlock(addr common.Address) {

	//before test
	borrowRatePerBlock, err := this.BorrowRatePerBlock()
	supplyingRatePerBlock, _ := this.SupplyRatePerBlock()
	if err != nil {
		log.Errorf("BorrowRatePerBlock err : %v", err)
	}
	hash0, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash0 AccrueInterest err : %v addr.ToHexString():%s", err, addr.ToHexString())
		return
	}
	userUnderlying0, _ := this.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash0)
	//startTime := GetTimeByTxhash(this.sdk, hash0)
	startTime, _ := this.sdk.GetBlockHeightByTxHash(hash0)

	totalBorrow, _ := this.TotalBorrows()
	totalSupply, _ := this.TotalSupply()
	reserveFactor, _ := this.ReserveFactorMantissa()
	accountSnapshot, _ := this.AccountSnapshot(addr)
	userSupply := accountSnapshot.TokenBalance
	reserves, _ := this.TotalReserves()
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	time.Sleep(time.Second * 300)
	hash1, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash1 AccrueInterest err : %v addr.ToHexString():%s", err, addr.ToHexString())
		return
	}
	userUnderlying1, _ := this.BalanceOfUnderlying(addr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	endTime, _ := this.sdk.GetBlockHeightByTxHash(hash1)
	//endTime := GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("addr.ToHexString(): %s, endTime:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v", addr.ToHexString(), endTime, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("addr.ToHexString(): %s, borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", addr.ToHexString(), borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("addr.ToHexString(): %s,totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", addr.ToHexString(), totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(1e9))
	log.Infof("addr.ToHexString(): %s,reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", addr.ToHexString(), reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Mul(big.NewInt(0).Div(userSupply, totalSupply), expSupplyInterestAdd)
	log.Infof("addr.ToHexString(): %s,userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v", addr.ToHexString(), userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}
