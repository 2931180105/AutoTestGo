package dbHelper

var INSERT_WING_DIS_RES = `INSERT INTO wing_dis_result_borrow VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,  ?,?);`
var INSERT_WING_DIS_RES2 = `INSERT INTO wing_dis_result_supply VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,  ?,?);`


var QUERY_DATA = `SELECT * FROM account_info limit ?,?;`

var QUERY_Addr = `SELECT * FROM account_info WHERE base58=?;`
var DELETE_DATA = `DELETE FROM student WHERE age>=30`
var UPDATE_Staking = `UPDATE account_info SET staking_amount=? WHERE base58=?;`
var UPDATE_DATA = `UPDATE account_info SET balance_wing=? WHERE base58=?;`
var INSERT_DATA = `INSERT INTO account_info(base58,wif,balance_ont,balance_wing,staking_amount) VALUES(?,?,?,?,?);`
