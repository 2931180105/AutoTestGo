package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

func GetGovTokenAddres(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	//WingToken, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "get_governance_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
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
	AuthAddr, _ := utils.AddressFromBase58(cfg.Owner)
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	params := []interface{}{AuthAddr, ZeroPoolAddr, cfg.Weight}
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
func GetProductPools(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{"11"}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_product_pools", params)
	log.Infof("get_product_pools: %s", resut.Result)
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "get_product_pools", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
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

//add_support_token TODO: invoke failed
func Add_support_token(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	params := []interface{}{OUSDTAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "add_support_token", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_support_token
func Get_support_token(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_support_token", params)
	log.Infof("get_support_token: %s", resut.Result)
	return resut
}

//set_exchange_rate TODO: invoke failed
func Set_exchange_rate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	params := []interface{}{OUSDTAddr, cfg.ExchangeRate}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "set_exchange_rate", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}

//get_exchange_rate
func Get_exchange_rate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *common.PreExecResult {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	params := []interface{}{}
	resut, _ := genSdk.WasmVM.PreExecInvokeWasmVMContract(WingGovAddr, "get_exchange_rate", params)
	log.Infof("get_exchange_rate: %s", resut.Result)
	return resut
}

//migrate TODO: not finished
func migrate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) *types.MutableTransaction {
	WingGovAddr, _ := utils.AddressFromHexString(cfg.WingGov)
	OUSDTAddr, _ := utils.AddressFromHexString(cfg.OUSDT)
	params := []interface{}{OUSDTAddr, cfg.ExchangeRate}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, WingGovAddr, "migrate", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	return mutTx
}
