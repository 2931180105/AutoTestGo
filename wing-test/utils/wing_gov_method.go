package utils

import (
	"encoding/hex"
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/payload"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/states"
	"io/ioutil"
	"time"
)

//gov init
func WingGovInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingProfit := getContractAddr(cfg.WingProfit)
	Oracle := getContractAddr(cfg.Oracle)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	GlobalParam := getContractAddr(cfg.GlobalParam)
	params := []interface{}{WingToken, WingProfit, Oracle, GlobalParam, cfg.SDRate}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

func Set_oracle_address(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	Oracle := getContractAddr(cfg.Oracle)
	params := []interface{}{Oracle}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_oracle_address", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_oracle_address
func Get_oracle_address(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_oracle_address", params)
	log.Infof("get_oracle_address: %s", resut.Result)
}

func Set_global_address(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	GlobalParam, _ := utils.AddressFromHexString(cfg.GlobalParam)
	params := []interface{}{GlobalParam}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_global_address", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_oracle_address
func Get_global_address(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_global_address", params)
	log.Infof("get_global_address: %s", resut.Result)
}

//GetGovTokenAddres
func GetGovTokenAddres(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_governance_token", params)
	log.Infof("GetGovTokenAddres: %s", resut.Result)
}

func SetGovTokenAddres(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingToken, _ := utils.AddressFromHexString(cfg.GovToken)

	params := []interface{}{WingToken}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_governance_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//authorize management pool
func SetAuth(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	AuthAddr, _ := utils.AddressFromBase58(cfg.AuthAddr)
	params := []interface{}{AuthAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "authorize", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//revoke
func RevokeAuth(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	AuthAddr, _ := utils.AddressFromBase58(cfg.AuthAddr)
	params := []interface{}{AuthAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "revoke", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_authorize_status, ToDo : improve pre exe
func GetAuthorizeStatus(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	AuthAddr, _ := utils.AddressFromBase58(cfg.AuthAddr)
	params := []interface{}{AuthAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "get_authorize_status", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_authorize_status", params)
	log.Infof("get_authorize_status: %s", resut.Result)
	return mutTx
}

//register_pool
func RegisterPool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{account.Address, ZeroPoolAddr, cfg.Weight}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "register_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_product_pools, TODO: need read struct
func GetProductPools(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{"11"}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_product_pools", params)
	log.Infof("get_product_pools: %s", resut.Result)

}

//update_pool_weight
func UpdatePoolWeight(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	AuthAddr, _ := utils.AddressFromBase58(cfg.Owner)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{AuthAddr, ZeroPoolAddr, cfg.Weight}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "update_pool_weight", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//query_pool_by_address
func QueryPoolByAddress(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "query_pool_by_address", params)
	log.Infof("get_product_pools: %s", resut.Result)
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "query_pool_by_address", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//update_pool_address, ToDo: need test more
func UpdatePoolAddress(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	AuthAddr, _ := utils.AddressFromBase58(cfg.Owner)
	OldPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	NewPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)

	params := []interface{}{AuthAddr, OldPoolAddr, NewPoolAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "update_pool_address", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//query_pool_count
func QueryPoolCount(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "query_pool_count", params)
	log.Infof("get_product_pools: %s", resut.Result)

	return resut
}

//update_profit_contract
func UpdateProfitContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	WingProfitAddr, _ := utils.AddressFromHexString(cfg.WingProfit)

	params := []interface{}{WingProfitAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "update_profit_contract", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//query_pool_count
func GetProfitContract(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_profit_contract", params)
	log.Infof("get_product_pools: %s", resut.Result)

	return resut
}

//unbound_token ToDo: invoke failed
func UnboundToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "unbound_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_unbound_pool
func Get_unbound_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_unbound_pool", params)
	log.Infof("Get_unbound_pool: %s", resut.Result)

	return resut
}

//unbound_to_governance ToDo: invoke failed
func Unbound_to_governance(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "unbound_to_governance", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//unbound_to_pool ToDo: invoke failed
func Unbound_to_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "unbound_to_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//query_unbound_to_pool
func Query_unbound_to_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr, 1}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "query_unbound_to_pool", params)
	log.Infof("Query_unbound_to_pool: %s", resut.Result)
	return resut
}

//query_unbound_to_pool_count ToDO : name ?
func Query_unbound_to_pool_count(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "query_unbound_to_pool_count", params)
	log.Infof("Get_unbound_pool: %s", resut.Result)
	return resut
}

//deposit TODO: invoke failed
func Deposit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	falge := true
	FromAddr := account.Address

	params := []interface{}{FromAddr, ZeroPoolAddr, OUSDTAddr, cfg.Amount, falge}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "unbound_to_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//withdraw TODO: invoke failed
func Withdraw(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	falge := true
	ToAddr := account.Address
	params := []interface{}{ToAddr, ZeroPoolAddr, OUSDTAddr, cfg.Amount, falge}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "withdraw", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_lock_asset
func Get_lock_asset(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	params := []interface{}{ZeroPoolAddr, OUSDTAddr}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_lock_asset", params)
	log.Infof("Get_lock_asset: %s", resut.Result)
	return resut
}

//black_pool
func Black_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	falge := true
	ToAddr := account.Address
	params := []interface{}{ToAddr, ZeroPoolAddr, OUSDTAddr, cfg.Amount, falge}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "black_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//white_pool
func White_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "white_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//upgrade_pool
func Upgrade_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "upgrade_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//freeze_pool
func Freeze_pool(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{ZeroPoolAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "freeze_pool", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//freeze_pool
func SetFFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{cfg.SDRate}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_f_factor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//add_support_token
func Add_support_token(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	OETHAddr, _ := utils.AddressFromHexString(cfg.OETH)
	token := NewToken("OETH", 2, OETHAddr)
	sink := OntCommon.NewZeroCopySink(nil)
	sink.WriteString("add_support_token")
	token.Serialize(sink)
	sink.WriteI128(OntCommon.I128FromUint64(20))
	contract := &states.WasmContractParam{}
	contract.Address = WingGovAddr
	argbytes := sink.Bytes()
	contract.Args = argbytes
	invokePayload := &payload.InvokeCode{
		Code: OntCommon.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		Payer:    account.Address,
		GasPrice: 2500,
		GasLimit: 300000,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	if err := signTx(genSdk, tx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return tx
}

//get_support_token
func Get_support_token(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_support_token", params)
	log.Infof("get_support_token: %s", resut.Result)
	return resut
}

//get_support_token
func Get_f_fatcor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_f_factor", params)
	log.Infof("Get_f_fatcor: %s", resut.Result)
	return resut
}

//set_exchange_rate
func Set_exchange_rate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	ONT, _ := utils.AddressFromBase58(cfg.ONT)
	token := NewToken("ONT", 1, ONT)
	sink := OntCommon.NewZeroCopySink(nil)
	sink.WriteString("set_exchange_rate")
	token.Serialize(sink)
	sink.WriteI128(OntCommon.I128FromUint64(1000000000))
	contract := &states.WasmContractParam{}
	contract.Address = WingGovAddr
	argbytes := sink.Bytes()
	contract.Args = argbytes
	invokePayload := &payload.InvokeCode{
		Code: OntCommon.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		Payer:    account.Address,
		GasPrice: 2500,
		GasLimit: 300000,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	if err := signTx(genSdk, tx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return tx
}

//get_exchange_rate
func Get_exchange_rate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	OETHAddr, _ := utils.AddressFromBase58(cfg.ONT)
	token := NewToken("ONT", 1, OETHAddr)
	sink := OntCommon.NewZeroCopySink(nil)
	sink.WriteString("get_exchange_rate")
	token.Serialize(sink)
	contract := &states.WasmContractParam{}
	contract.Address = WingGovAddr
	argbytes := sink.Bytes()
	contract.Args = argbytes
	invokePayload := &payload.InvokeCode{
		Code: OntCommon.SerializeToBytes(contract),
	}
	tx := &types.MutableTransaction{
		Payer:    account.Address,
		GasPrice: 2500,
		GasLimit: 300000,
		TxType:   types.InvokeWasm,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  invokePayload,
		Sigs:     nil,
	}
	resut, err := genSdk.PreExecTransaction(tx)
	if err != nil {
		log.Errorf("error: %s", err)
	}
	log.Infof("result: %s", resut.Result)
}

//set_exchange_rates batch TODO: invoke failed
func Set_exchange_rates(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_exchange_rates", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//migrate TODO: not finished
func WingGovMigrate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	bytes, err := ioutil.ReadFile("wing-test/contract/wing_dao_contracts_new.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	CodeStr, _ := hex.DecodeString(string(bytes))
	CodeContractAddr, err := utils.GetContractAddress(string(bytes))
	if err != nil {
		log.Error(err)
	}
	log.Infof("CodeContractAddr address : %s", CodeContractAddr.ToBase58())
	log.Infof("CodeContractAddr address : %s", CodeContractAddr.ToHexString())

	params := []interface{}{CodeStr, 3, "WING Token", "1.0.1", "Wing Team", "support@wing.finance", "Wing is a credit-based, cross-chain DeFi platform."}
	mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "migrate", params)

	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash1, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	PrintSmartEventByHash_Ont(genSdk, hash1.ToHexString())
}

//migrate TODO: not finished
func Destroy(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "destroy", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
