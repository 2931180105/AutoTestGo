package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"math/big"
	"strings"
	"time"
)

/*
	1.use market name and usraddr
 */
func (this *FlashToken)TestBorrowRate2Supply(marketName, addr string) {
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
	usrAddr,_ := common.AddressFromBase58(addr)
	//before test
	borrowRatePerBlock, err := this.BorrowRatePerBlock()
	supplyingRatePerBlock, _ := this.SupplyRatePerBlock()
	if err != nil {
		log.Errorf("BorrowRatePerBlock err : %v", err)
	}
	hash0, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash0 AccrueInterest err : %v usrAddr%s", err, usrAddr)
		return
	}
	userUnderlying0, _ := this.BalanceOfUnderlying(usrAddr)
	utils.PrintSmartEventByHash_Ont(this.sdk, hash0)
	startTime := utils.GetTimeByTxhash(this.sdk, hash0)
	totalBorrow, _ := this.TotalBorrows()
	totalSupply, _ := this.TotalSupply()
	reserveFactor, _ := this.ReserveFactorMantissa()
	accountSnapshot, _ := this.AccountSnapshot(usrAddr)
	userSupply := accountSnapshot.TokenBalance
	reserves, _ := this.TotalReserves()
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	time.Sleep(time.Second * 300)
	hash1, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash1 AccrueInterest err : %v usrAddr:%s", err, usrAddr)
		return
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	userUnderlying1, _ := this.BalanceOfUnderlying(usrAddr)
	endTime := utils.GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("startTime:%v,endTime:%v,totalBorrow:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v",startTime, endTime,totalBorrow, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	//expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(10^18))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor),utils.ToIntByPrecise("1",9))
	log.Infof("reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Div(big.NewInt(0).Mul(userSupply, expSupplyInterestAdd), totalSupply)
	log.Infof("userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v",  userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}
