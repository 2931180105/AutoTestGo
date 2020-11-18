package model

import (
	"github.com/jinzhu/gorm"
)

type WingDisResultBorrow struct {
	gorm.Model
	MarketName string `json:"market_name"`
	UserAddr string `json:"user_addr"`
	MarketAddr string `json:"market_addr"`
	TotalValidBorrow string `json:"total_valid_borrow"`
	UserValidBorrow string `json:"user_valid_borrow"`
	WingSpeed string `json:"wing_speed"`
	ExpRes string `json:"exp_res"`
	RelRes string `json:"rel_res"`
	StartTime uint32 `json:"start_time"`
	EndTime uint32 `json:"end_time"`
	ErrRate string `json:"err_rate"`
}

type WingDisResultSupply struct {
	gorm.Model
	MarketName string `json:"market_name"`
	UserAddr string `json:"user_addr"`
	MarketAddr string `json:"market_addr"`
	TotalSupply string `json:"total_supply"`
	UserSupply string `json:"user_supply"`
	WingSpeed string `json:"wing_speed"`
	ExpRes string `json:"exp_res"`
	RelRes string `json:"rel_res"`
	StartTime uint32 `json:"start_time"`
	EndTime uint32 `json:"end_time"`
	ErrRate string `json:"err_rate"`
}

func NewWingDisResultSupply(marketName string, userAddr string, marketAddr string, totalSupply string, userSupply string, wingSpeed string, expRes string, relRes string, startTime uint32, endTime uint32, errRate string) *WingDisResultSupply {
	return &WingDisResultSupply{MarketName: marketName, UserAddr: userAddr, MarketAddr: marketAddr, TotalSupply: totalSupply, UserSupply: userSupply, WingSpeed: wingSpeed, ExpRes: expRes, RelRes: relRes, StartTime: startTime, EndTime: endTime, ErrRate: errRate}
}
type AccountInfo struct {
	ID uint32 `json:"id"`
	Base58 string `json:"base58"`
	Wif string `json:"wif"`
	BalanceOnt string `json:"balance_ont"`
	BalanceWing string `json:"balance_wing"`
	StakingAmount int32 `json:"staking_amount"`
	StakingTime int32 `json:"staking_time"`
}

type BorrowRateSupply struct {
	gorm.Model
	MarketName            string `json:"market_name"`
	UserAddr              string `json:"user_addr"`
	MarketAddr            string `json:"market_addr"`
	BorrowRatePerBlock    string `json:"borrow_rate_per_block"`
	SupplyingRatePerBlock string `json:"supplying_borrow_rate_per_block"`
	UserUnderlyingAdd       string `json:"user_underlying_add"`
	TotalBorrowBefore     string `json:"total_borrow_before"`
	TotalBorrowAfter      string `json:"total_borrow_after"`
	ReserveFactor         string `json:"reserve_factor"`
	ReservesAdd         string `json:"reserves_add"`
	ExpReservesAdd         string `json:"exp_reserves_add"`
	TotalInterestAdd        string `json:"total_interest_add"`
	ExpTotalInterestAdd        string `json:"exp_total_interest_add"`
	ExpUserUnderlyingAdd        string `json:"exp_user_underlying_add"`
	UserSupply            string `json:"user_supply"`
	TotalSupply           string `json:"total_supply"`
	StartTime             uint32 `json:"start_time"`
	EndTime               uint32 `json:"end_time"`
	ErrRateTotalInterest               string `json:"err_rate_total_interest"`
	ErrRateReserves       string `json:"err_rate_reserves_interest"`
	ErrRateUser      string `json:"err_rate_user_interest"`
}

func NewBorrowRate(marketName string, userAddr string, marketAddr string, borrowRatePerBlock string, supplyingRatePerBlock string, userUnderlyingAdd string, totalBorrowBefore string, totalBorrowAfter string, reserveFactor string, reservesAdd string, expReservesAdd string, totalInterestAdd string, expTotalInterestAdd string, expUserUnderlyingAdd string, userSupply string, totalSupply string, startTime uint32, endTime uint32, errRateTotalInterest string, errRateReserves string, errRateUser string) *BorrowRateSupply {
	return &BorrowRateSupply{MarketName: marketName, UserAddr: userAddr, MarketAddr: marketAddr, BorrowRatePerBlock: borrowRatePerBlock, SupplyingRatePerBlock: supplyingRatePerBlock, UserUnderlyingAdd: userUnderlyingAdd, TotalBorrowBefore: totalBorrowBefore, TotalBorrowAfter: totalBorrowAfter, ReserveFactor: reserveFactor, ReservesAdd: reservesAdd, ExpReservesAdd: expReservesAdd, TotalInterestAdd: totalInterestAdd, ExpTotalInterestAdd: expTotalInterestAdd, ExpUserUnderlyingAdd: expUserUnderlyingAdd, UserSupply: userSupply, TotalSupply: totalSupply, StartTime: startTime, EndTime: endTime, ErrRateTotalInterest: errRateTotalInterest, ErrRateReserves: errRateReserves, ErrRateUser: errRateUser}
}