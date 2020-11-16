package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"math/big"
	"strings"
	"time"
)

func (this *FlashToken) TestBorrowRateByBlock() {
	//before test
	borrowRatePerBlock, err := this.BorrowRatePerBlock()
	supplyingRatePerBlock, _ := this.SupplyRatePerBlock()
	if err != nil {
		log.Errorf("BorrowRatePerBlock err : %v", err)
	}
	hash0, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash0 AccrueInterest err : %v this.signer.Address:%s", err, this.signer.Address)
		return
	}
	userUnderlying0, _ := this.BalanceOfUnderlying(this.signer.Address)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash0)
	startTime := utils.GetTimeByTxhash(this.sdk, hash0)
	//startTime, _ := this.sdk.GetBlockHeightByTxHash(hash0)

	totalBorrow, _ := this.TotalBorrows()
	totalSupply, _ := this.TotalSupply()
	reserveFactor, _ := this.ReserveFactorMantissa()
	accountSnapshot, _ := this.AccountSnapshot(this.signer.Address)
	userSupply := accountSnapshot.TokenBalance
	reserves, _ := this.TotalReserves()
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	time.Sleep(time.Second * 300)
	hash1, err := this.AccrueInterest()
	if err != nil {
		log.Errorf("  hash1 AccrueInterest err : %v this.signer.Address:%s", err, this.signer.Address)
		return
	}
	userUnderlying1, _ := this.BalanceOfUnderlying(this.signer.Address)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	//endTime, _ := this.sdk.GetBlockHeightByTxHash(hash1)
	endTime := utils.GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("this.signer.Address: %s, endTime:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v", this.signer.Address, endTime, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("this.signer.Address: %s, borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", this.signer.Address, borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("this.signer.Address: %s,totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", this.signer.Address, totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(1e9))
	log.Infof("this.signer.Address: %s,reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", this.signer.Address, reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Div(big.NewInt(0).Mul(userSupply, expSupplyInterestAdd), totalSupply)
	log.Infof("this.signer.Address: %s,userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v", this.signer.Address, userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}


/*
	1.use market name and usraddr
 */
func (this *FlashToken) TestBorrowRate2Supply(marketName, addr string) {
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
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash0)
	startTime := utils.GetTimeByTxhash(this.sdk, hash0)
	//startTime, _ := this.sdk.GetBlockHeightByTxHash(hash0)
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
	userUnderlying1, _ := this.BalanceOfUnderlying(usrAddr)
	if err != nil {
		log.Errorf("  hash1 BalanceOfUnderlying err : %v", err)
	}
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	endTime := utils.GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("usrAddr: %s, endTime:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v", usrAddr, endTime, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("usrAddr: %s, borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", usrAddr, borrowRatePerBlock, startTime, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := big.NewInt(0).Sub(reserves1, reserves)
	totalInterestAdd := big.NewInt(0).Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := big.NewInt(0).Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("usrAddr: %s,totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", usrAddr, totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := big.NewInt(0).Div(big.NewInt(0).Mul(expTotalInterestAdd, reserveFactor), big.NewInt(1e9))
	log.Infof("usrAddr: %s,reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", usrAddr, reservesAdd, expReservesAdd, big.NewInt(0).Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock)
	expUserUnderlyingAdd := big.NewInt(0).Div(big.NewInt(0).Mul(userSupply, expSupplyInterestAdd), totalSupply)
	log.Infof("usrAddr: %s,userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v", usrAddr, userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
}
