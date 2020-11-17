package dbHelper

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"math/big"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 建立数据库连接
func SetupConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(172.168.3.219:3306)/wing-test?charset=utf8")
	errorHandler(err)
	return db
}

// 创建表
func CreateTable(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Errorf("create table error : %s", err)
	}
	errorHandler(err)
}


// 插入数据
func Insert(db *sql.DB, base58 string, wif string, balance_ont int, balance_wing int, staking_amount int) {
	_, err := db.Exec(INSERT_DATA, base58, wif, balance_ont, balance_wing, staking_amount)
	if err != nil {
		log.Errorf("Insert data error : %s", err)
	}
}

// 插入数据
func InsertWingDisRes(db *sql.DB, market_name,user_addr,market_addr string,total_valid_borrow,user_valid_borrow,wing_speed,exp_res,rel_res *big.Int,start_time,end_time uint32 ,err_rate *big.Float ) {
	errRate ,_:=err_rate.Float64()
	_, err := db.Exec(INSERT_WING_DIS_RES,market_name,user_addr,market_addr,total_valid_borrow.Int64(),user_valid_borrow.Int64(),wing_speed.Int64(),exp_res.Int64(),rel_res.Int64(),start_time,end_time,errRate ,time.Now())
	if err != nil {
		log.Errorf("Insert data error : %s", err)
	}
}
// 插入数据
func InsertWingDisRes4Supply(db *sql.DB, market_name,user_addr,market_addr string,total_valid_borrow,user_valid_borrow,wing_speed,exp_res,rel_res *big.Int,start_time,end_time uint32 ,err_rate *big.Float ) {
	errRate ,_:=err_rate.Float64()
	_, err := db.Exec(INSERT_WING_DIS_RES2,market_name,user_addr,market_addr,total_valid_borrow.Int64(),user_valid_borrow.Int64(),wing_speed.Int64(),exp_res.Int64(),rel_res.Int64(),start_time,end_time,errRate ,time.Now())
	if err != nil {
		log.Errorf("Insert data error : %s", err)
	}
}
// 修改数据
func Update(db *sql.DB, wing_balance uint64, base58 string) (sql.Result, error) {
	return db.Exec(UPDATE_DATA, wing_balance, base58)
}


// 修改数据
func UpdateStakingBalance(db *sql.DB, wing_balance uint64, base58 string) (sql.Result, error) {
	return db.Exec(UPDATE_Staking, wing_balance, base58)
}


// 删除记录
func Delete(db *sql.DB) {
	db.Exec(DELETE_DATA)
}

var DELETE_TABLE = `DROP TABLE student;`

// 删除表
func DeleteTable(db *sql.DB) {
	db.Exec(DELETE_TABLE)
}

// 查询数据
func Query(db *sql.DB, start, end int) *sql.Rows {
	rows, err := db.Query(QUERY_DATA, start, end)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

// 查询数据
func QueryByAddr(db *sql.DB, addr string) *sql.Rows {
	rows, err := db.Query(QUERY_Addr, addr)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func QueryAccountFromDb(start, end int) []*goSdk.Account {
	db := SetupConnect()
	rows := Query(db, start, end)
	accounts := make([]*goSdk.Account, 0)
	for rows.Next() {
		var id int
		var base58 string
		var wif string
		var balance_ont int
		var balance_wing int
		var stakeing_amount int
		var stakeing_time int
		if err := rows.Scan(&id, &base58, &wif, &balance_ont, &balance_wing, &stakeing_amount, &stakeing_time); err != nil {
			log.Infof("error ", err)
		}
		pkey, _ := hex.DecodeString(wif)
		pri, _ := keypair.DeserializePrivateKey(pkey)
		acct := goSdk.Account{
			PrivateKey: pri,
			PublicKey:  pri.Public(),
			Address:    types.AddressFromPubKey(pri.Public()),
			SigScheme:  signature.SHA256withECDSA,
		}
		//log.Infof("%s", acct.SigScheme.Name())
		//log.Infof("base58 :%s, tobase58:%s", base58, acct.Address.ToBase58())
		accounts = append(accounts, &acct)

	}
	db.Close()

	return accounts
}

func QueryAccountFromDbByBase58(address string) []*goSdk.Account {
	db := SetupConnect()
	rows := QueryByAddr(db, address)
	accounts := make([]*goSdk.Account, 0)
	for rows.Next() {
		var id int
		var base58 string
		var wif string
		var balance_ont int
		var balance_wing int
		var stakeing_amount int
		var stakeing_time int
		if err := rows.Scan(&id, &base58, &wif, &balance_ont, &balance_wing, &stakeing_amount, &stakeing_time); err != nil {
			log.Infof("error ", err)
		}
		pkey, _ := hex.DecodeString(wif)
		pri, _ := keypair.DeserializePrivateKey(pkey)
		acct := goSdk.Account{
			PrivateKey: pri,
			PublicKey:  pri.Public(),
			Address:    types.AddressFromPubKey(pri.Public()),
			SigScheme:  signature.SHA256withECDSA,
		}
		//log.Infof("%s", acct.SigScheme.Name())
		//log.Infof("base58 :%s, tobase58:%s", base58, acct.Address.ToBase58())
		accounts = append(accounts, &acct)

	}
	db.Close()

	return accounts
}
