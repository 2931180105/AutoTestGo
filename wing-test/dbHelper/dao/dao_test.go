package dao

import (
	"github.com/mockyz/AutoTestGo/wing-test/dbHelper/model"
	"testing"
)

func TestInsertStat(t *testing.T) {
	cfg := &Config{
		Ip:   "172.168.3.219",
		Port: 3306,
		User: "root",
		Pwd:  "123456",
		Db:   "wing-test",
	}
	NewDao(cfg)
	//db.Table("want_name").AutoMigrate(&model.WingDisResultSupplyTmp{})
	//db.CreateTable(&model.WingDisResultBorrow{})
	//db.CreateTable(&model.WingDisResultSupply{})
	db.CreateTable(&model.BorrowRateSupply{})

	//db.Table("nft_transfer_events").AutoMigrate(&event.NFTTransferEventModel{})
	//err := InsertStat("127.0.0.1", 30001, "0xjflaksdjfoi23rnlasdf", 10, 10)
	//if err != nil {
	//	t.Fatal(err)
	//}
}
