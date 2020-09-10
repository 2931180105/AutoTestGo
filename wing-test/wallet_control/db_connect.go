package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mockyz/AutoTestGo/common/log"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 建立数据库连接
func setupConnect() *sql.DB {
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

var INSERT_DATA = `INSERT INTO account_info(base58,wif,balance_ont,balance_wing,staking_amount) VALUES(?,?,?,?,?);`

// 插入数据
func Insert(db *sql.DB, base58 string, wif string, balance_ont int, balance_wing int, staking_amount int) {
	_, err := db.Exec(INSERT_DATA, base58, wif, balance_ont, balance_wing, staking_amount)
	if err != nil {
		log.Errorf("Insert data error : %s", err)
	}
}

var UPDATE_DATA = `UPDATE student SET age=28 WHERE sname="唐僧";`

// 修改数据
func Update(db *sql.DB) {
	db.Exec(UPDATE_DATA)

}

var DELETE_DATA = `DELETE FROM student WHERE age>=30`

// 删除记录
func Delete(db *sql.DB) {
	db.Exec(DELETE_DATA)
}

var DELETE_TABLE = `DROP TABLE student;`

// 删除表
func DeleteTable(db *sql.DB) {
	db.Exec(DELETE_TABLE)
}

var QUERY_DATA = `SELECT * FROM account_info limit ?,? ;`

// 查询数据
func Query(db *sql.DB, start, end int) {
	rows, err := db.Query(QUERY_DATA, start, end)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var base58 string
		var wif string
		var balance_ont int
		var balance_wing int
		var stakeing_amount int
		if err := rows.Scan(&id, &base58, &wif, &balance_ont, &balance_wing, &stakeing_amount); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s is %d\n", base58, balance_ont)
	}
}

func main() {
	// 建立数据连接
	db := setupConnect()
	// 创建数据库表
	//CreateTable(db, CREATE_TABLE)
	// 插入数据
	//Insert(db, "base58","wif",1,2,3)
	//pwd := []byte("123456")
	//wallet, _ := ontology_go_sdk.NewOntologySdk().CreateWallet("tmp2.dat")
	//account, err := wallet.NewDefaultSettingAccount(pwd)
	//if err != nil{
	//	log.Infof(" new account error : %s" , err)
	//}
	//base58 := account.Address.ToBase58()
	//hexWif := keypair.SerializePrivateKey(account.PrivateKey)
	//log.Infof("wif :%s",hex.EncodeToString(hexWif))
	//Insert(db, base58,hex.EncodeToString(hexWif),1,2,3)
	//wif,_ :=hex.DecodeString(hex.EncodeToString(hexWif))
	//account,_ :=ontology_go_sdk.NewAccountFromPrivateKey()
	Query(db, 1, 100)

	// 查询数据
	// 删除数据
	//Delete(db)
	//// 插入数据
	//Insert(db)
	//// 修改数据
	//Update(db)
	//// 查询数据
	//Query(db)
	//// 删除表
	//DeleteTable(db)
	//wallet, _ := ontology_go_sdk.NewOntologySdk().CreateWallet("tmp2.dat")
	//// 关闭数据库连接
	//wait := new(sync.WaitGroup)
	//for i := uint(0); i < 3; i++ {
	//	go func(nonce uint32, routineIndex uint) {
	//		wait.Add(1)
	//		defer wait.Done()
	//		for j :=0; j<10 ;j++  {
	//			pwd := []byte("123456")
	//			account, err := wallet.NewDefaultSettingAccount(pwd)
	//			if err != nil{
	//				log.Infof(" new account error : %s" , err)
	//			}
	//			base58 := account.Address.ToBase58()
	//			hexWif := keypair.SerializePrivateKey(account.PrivateKey)
	//			log.Infof("wif :%s",hex.EncodeToString(hexWif))
	//			Insert(db, base58,hex.EncodeToString(hexWif),1,2,3)
	//		}
	//	}(100,i)
	//}
	//wait.Wait()
	db.Close()
}
