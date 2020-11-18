package ftoken

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper/dao"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper/model"
	"github.com/mockyz/AutoTestGo/wing-test/utils"
	"github.com/ontio/ontology/common"
	"math/big"
	"strings"
	"sync"
	"time"
)

/*
	1.use market name and usraddr
 */
func (this *FlashToken)TestBorrowRate(marketName, addr string) {
	borrowRateSupply := new(model.BorrowRateSupply)
	borrowRateSupply.UserAddr=addr
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
	userUnderlying1, _ := this.BalanceOfUnderlying(usrAddr)
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	endTime := utils.GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("startTime:%v,endTime:%v,totalBorrow:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v",startTime, endTime,totalBorrow, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := utils.Sub(reserves1, reserves)
	totalInterestAdd := utils.Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := utils.Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := utils.Div(utils.Mul(expTotalInterestAdd, reserveFactor), utils.ToIntByPrecise("1",9))
	log.Infof("reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", reservesAdd, expReservesAdd, utils.Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.Div(utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock),big.NewInt(100))
	expUserUnderlyingAdd := utils.Div(utils.Mul(userSupply, expSupplyInterestAdd), totalSupply)
	log.Infof("userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v",  userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
	errTotalAdd := utils.CmpTestRuslt(expTotalInterestAdd,totalInterestAdd)
	errUser := utils.CmpTestRuslt(expUserUnderlyingAdd,userUnderlyingAdd)
	errReserves := utils.CmpTestRuslt(expReservesAdd,reservesAdd)
	dao.SaveBorrowRateSupply(marketName,addr,this.addr.ToHexString(),borrowRatePerBlock.String(),supplyingRatePerBlock.String(),userUnderlyingAdd.String(),totalBorrow.String(),totalBorrow1.String(),reserveFactor.String(),reservesAdd.String(),expReservesAdd.String(),totalInterestAdd.String(),expTotalInterestAdd.String(),expUserUnderlyingAdd.String(),userSupply.String(),totalSupply.String(),startTime,endTime,errTotalAdd.String(),errReserves.String(),errUser.String())
}


/*
	1.use market name and usraddr
*/
func (this *FlashToken)TestBorrowRateSync(marketName, addr string, sy *sync.WaitGroup)  {
	defer sy.Done()
	borrowRateSupply := new(model.BorrowRateSupply)
	borrowRateSupply.UserAddr=addr
	//ftokenAddressList, err :=  this.Comptroller.AllMarkets()
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
	userUnderlying1, _ := this.BalanceOfUnderlying(usrAddr)
	utils.PrintSmartEventByHash_Ont(this.sdk, hash1)
	endTime := utils.GetTimeByTxhash(this.sdk, hash1)
	totalBorrow1, _ := this.TotalBorrows()
	totalSupply1, _ := this.TotalSupply()
	reserves1, _ := this.TotalReserves()
	log.Infof("startTime:%v,endTime:%v,totalBorrow:%v,totalBorrow1:%v,totalSupply1:%v,reserves1:%v,userUnderlying1:%v",startTime, endTime,totalBorrow, totalBorrow1, totalSupply1, reserves1, userUnderlying1)
	log.Infof("borrowRatePerBlock:%v,startTime:%v,totalBorrow:%v,totalSupply:%v,reserveFactor:%v,userSupply:%v,reserves:%v,userUnderlying0:%v", borrowRatePerBlock, totalBorrow, totalSupply, reserveFactor, userSupply, reserves, userUnderlying0)
	reservesAdd := utils.Sub(reserves1, reserves)
	totalInterestAdd := utils.Sub(totalBorrow1, totalBorrow)
	userUnderlyingAdd := utils.Sub(userUnderlying1, userUnderlying0)
	expTotalInterestAdd := utils.ExpInterestAdd(totalBorrow, big.NewInt(int64(endTime-startTime)), borrowRatePerBlock)
	log.Infof("totalInterestAdd:%v,expTotalInterestAdd:%v,expTotalInterestAdd - totalInterestAdd : %v", totalInterestAdd, expTotalInterestAdd, big.NewInt(0).Sub(expTotalInterestAdd, totalInterestAdd))
	expReservesAdd := utils.Div(utils.Mul(expTotalInterestAdd, reserveFactor), utils.ToIntByPrecise("1",9))
	log.Infof("reservesAdd:%v,expReservesAdd:%v,totalInterestAdd - expInterestAdd: %v", reservesAdd, expReservesAdd, utils.Sub(expReservesAdd, reservesAdd))
	expSupplyInterestAdd := utils.Div(utils.ExpInterestAdd(totalSupply, big.NewInt(int64(endTime-startTime)), supplyingRatePerBlock),big.NewInt(100))
	expUserUnderlyingAdd := utils.Div(utils.Mul(userSupply, expSupplyInterestAdd), totalSupply)
	log.Infof("userUnderlying0:%v,expSupplyInterestAdd: %v,userUnderlyingAdd:%v,expUserUnderlyingAdd:%v,expUserUnderlyingAdd - userUnderlyingAdd: %v",  userUnderlying0, expSupplyInterestAdd, userUnderlyingAdd, expUserUnderlyingAdd, big.NewInt(0).Sub(userUnderlyingAdd, expUserUnderlyingAdd))
	errTotalAdd := utils.CmpTestRuslt(expTotalInterestAdd,totalInterestAdd)
	errUser := utils.CmpTestRuslt(expUserUnderlyingAdd,userUnderlyingAdd)
	errReserves := utils.CmpTestRuslt(expReservesAdd,reservesAdd)
	dao.SaveBorrowRateSupply(marketName,addr,this.addr.ToHexString(),borrowRatePerBlock.String(),supplyingRatePerBlock.String(),userUnderlyingAdd.String(),totalBorrow.String(),totalBorrow1.String(),reserveFactor.String(),reservesAdd.String(),expReservesAdd.String(),totalInterestAdd.String(),expTotalInterestAdd.String(),expUserUnderlyingAdd.String(),userSupply.String(),totalSupply.String(),startTime,endTime,errTotalAdd.String(),errReserves.String(),errUser.String())
}
