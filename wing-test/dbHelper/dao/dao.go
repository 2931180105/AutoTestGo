package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper/model"
	"math/big"
)

var db *gorm.DB

type Config struct {
	Ip   string
	Port uint16
	User string
	Pwd  string
	Db   string
}

func NewDao(c *Config) {
	if db != nil {
		return
	}
	var err error
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?", c.User, c.Pwd, c.Ip, c.Port, c.Db)
	if db, err = gorm.Open("mysql", url); err != nil {
		panic("failed to connect database")
	}

}

type Stat struct {
	Ip   string
	Port uint16
	Hash string
	Send uint64
	Recv uint64
}

func InsertStat(ip string, port uint16, hash string, send, recv uint64) error {
	return db.Create(&Stat{Ip: ip, Port: port, Hash: hash, Send: send, Recv: recv}).Error
}

type Subnet struct {
	Pubkey string
}

func CheckSubnet(id string) bool {
	data := new(Subnet)
	affected := db.Where("pubkey = ?", id).First(data).RowsAffected
	return affected > 0
}

func GetAllSubnet() []string {
	list := []string{}
	ids := []*Subnet{}
	db.Find(&ids)
	for _, v := range ids {
		list = append(list, v.Pubkey)
	}
	return list
}

func init() {
	cfg := &Config{
		Ip:   "172.168.3.219",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "wing-test",
	}
	NewDao(cfg)
}

func SaveWingDisResultBorrow(market_name,user_addr,market_addr string,total_valid_borrow,user_valid_borrow,wing_speed,exp_res,rel_res *big.Int,start_time,end_time uint32 ,err_rate *big.Float ) {
	data := &model.WingDisResultBorrow{MarketName:market_name,UserAddr:user_addr,MarketAddr:market_addr,TotalValidBorrow:total_valid_borrow.String(),UserValidBorrow:user_valid_borrow.String(),WingSpeed:wing_speed.String(),ExpRes:exp_res.String(),RelRes:rel_res.String(),StartTime:start_time,EndTime:end_time,ErrRate:err_rate.String()}
	db.Create(data)
}
func SaveWingDisResultSupply(market_name,user_addr,market_addr string,total_valid_borrow,user_valid_borrow,wing_speed,exp_res,rel_res *big.Int,start_time,end_time uint32 ,err_rate *big.Float ) {
	data := &model.WingDisResultSupply{MarketName:market_name,UserAddr:user_addr,MarketAddr:market_addr,TotalSupply:total_valid_borrow.String(),UserSupply:user_valid_borrow.String(),WingSpeed:wing_speed.String(),ExpRes:exp_res.String(),RelRes:rel_res.String(),StartTime:start_time,EndTime:end_time,ErrRate:err_rate.String()}
	db.Create(data)
}

func SaveBorrowRateSupply(marketName string, userAddr string, marketAddr string, borrowRatePerBlock string, supplyingRatePerBlock string, userUnderlyingAdd string, totalBorrowBefore string, totalBorrowAfter string, reserveFactor string, reservesAdd string, expReservesAdd string, totalInterestAdd string, expTotalInterestAdd string, expUserUnderlyingAdd string, userSupply string, totalSupply string, startTime uint32, endTime uint32, errRateTotalInterest string, errRateReserves string, errRateUser string) {
	data := &model.BorrowRateSupply{MarketName: marketName, UserAddr: userAddr, MarketAddr: marketAddr, BorrowRatePerBlock: borrowRatePerBlock, SupplyingRatePerBlock: supplyingRatePerBlock, UserUnderlyingAdd: userUnderlyingAdd, TotalBorrowBefore: totalBorrowBefore, TotalBorrowAfter: totalBorrowAfter, ReserveFactor: reserveFactor, ReservesAdd: reservesAdd, ExpReservesAdd: expReservesAdd, TotalInterestAdd: totalInterestAdd, ExpTotalInterestAdd: expTotalInterestAdd, ExpUserUnderlyingAdd: expUserUnderlyingAdd, UserSupply: userSupply, TotalSupply: totalSupply, StartTime: startTime, EndTime: endTime, ErrRateTotalInterest: errRateTotalInterest, ErrRateReserves: errRateReserves, ErrRateUser: errRateUser}
	db.Create(data)
}



