package wingGov

import (
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

//
import (
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
)

func DeployContractProfit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, wasmFilePath string) common.Uint256 {
	bytes, err := ioutil.ReadFile(wasmFilePath)
	if err != nil {
		log.Fatal(err)
	}
	profitCodeStr := string(bytes)
	profitContract, err := utils.GetContractAddress(profitCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("profitContract address : %s", profitContract.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, profitCodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}

func DeployContractt(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, wasmFilePath string) string {
	bytes, err := ioutil.ReadFile(wasmFilePath)
	if err != nil {
		log.Fatal(err)
	}
	contractCodeStr := string(bytes)
	Contract, err := utils.GetContractAddress(contractCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("contract address : %s", Contract.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, contractCodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	log.Infof("DeployContractt reslut : %s", result.ToHexString())
	Utils.PrintSmartEventByHash_Ont(genSdk, result.ToHexString())
	return Contract.ToHexString()
}

func DeployContractOracle(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/oracle.wasm.str")
	if err != nil {
		log.Fatal(err)
	}
	CodeStr := string(bytes)
	CodeContractAddr, err := utils.GetContractAddress(CodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Oracle address : %s", CodeContractAddr.ToHexString())
	result, err := genSdk.WasmVM.DeployWasmVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, CodeStr, "name", "version", "author", "email", "desc")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}

func DeployContractWingToken(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) common.Uint256 {
	bytes, err := ioutil.ReadFile("wing-test/contract/testnet/wing.avm")
	if err != nil {
		log.Fatal(err)
	}
	CodeStr := string(bytes)
	CodeContractAddr, err := utils.GetContractAddress(CodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Wing address : %s", CodeContractAddr.ToHexString())
	result, err := genSdk.NeoVM.DeployNeoVMSmartContract(cfg.GasPrice, cfg.GasLimit, account, true, CodeStr, "WING Token", "1.0.1", "Wing Team", "support@wing.finance", "Wing is a credit-based, cross-chain DeFi platform.")
	if err != nil {
		log.Errorf("deployContreac  failed: %s", err)
	}
	return result
}
func BatchStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, accts []*goSdk.Account) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	for i := 0; i < len(accts); i++ {
		acct := accts[i]
		params := []interface{}{acct.Address, cfg.StakeOnt}
		mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "staking", params)
		if err := signTx(genSdk, mutTx, cfg.StartNonce, acct); err != nil {
			log.Error(err)
		}
		time.Sleep(time.Second)
		txhash, err := genSdk.SendTransaction(mutTx)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send staking tx %s****sentnum:***%d", txhash.ToHexString(), i)
		}
	}
}

func BatchUnStaking(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk, accts []*goSdk.Account) {
	ZeroPoolAddr, _ := utils.AddressFromHexString(cfg.ZeroPool)
	//accts := GenerateAccounts(cfg, account, genSdk)
	for i := 0; i < len(accts); i++ {
		acct := accts[i]
		params := []interface{}{acct.Address, cfg.StakeOnt}
		mutTx, _ := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, ZeroPoolAddr, "unstaking", params)
		if err := signTx(genSdk, mutTx, cfg.StartNonce, acct); err != nil {
			log.Error(err)
		}
		txhash, err := genSdk.SendTransaction(mutTx)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send unstaking tx %s****sentnum:***%d", txhash.ToHexString(), i)
		}
		ZeroPoolWithDraw(cfg, acct, genSdk)
		time.Sleep(time.Second * 3)
		resut2, _ := genSdk.NeoVM.PreExecInvokeNeoVMContract(ZeroPoolAddr, []interface{}{"balanceOf", []interface{}{acct.Address}})
		log.Infof("Address %s , WIng Token BalanceOf : %s", acct.Address.ToBase58(), resut2)
	}
}
func AddAllSupportToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk) {

	hash1, err := sdk.SendTransaction(Add_support_token(cfg, account, sdk, "ONTd", cfg.ONTD))
	if err != nil {
		log.Errorf("send ONTd tx failed, err: %s********", err)
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "USDC", cfg.OUSDC))
	if err != nil {
		log.Errorf("send USDC tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "WBTC", cfg.OWBTC))
	if err != nil {
		log.Errorf("send  WBTC failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "renBTC", cfg.RENBTC))
	if err != nil {
		log.Errorf("send  renBTC failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "WING", cfg.GovToken))
	if err != nil {
		log.Errorf("send  WING failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Add_support_token(cfg, account, sdk, "ETH", cfg.OETH))
	if err != nil {
		log.Errorf("send  ETH failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

}

func UpdateAllSuuportToken(cfg *config.Config, account *goSdk.Account, sdk *goSdk.OntologySdk) {
	hash1, err := sdk.SendTransaction(Update_support_token(cfg, account, sdk, "DAI", cfg.ODAI))
	if err != nil {
		log.Errorf("send DAI tx failed, err: %s********", err)
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Update_support_token(cfg, account, sdk, "ETH", cfg.OETH))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())
	hash1, err = sdk.SendTransaction(Update_support_token(cfg, account, sdk, "BTC", cfg.OWBTC))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	hash1, err = sdk.SendTransaction(Update_support_token(cfg, account, sdk, "USDT", cfg.OUSDT))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	hash1, err = sdk.SendTransaction(Update_support_token(cfg, account, sdk, "ONTd", cfg.ONTD))
	if err != nil {
		log.Errorf("send  tx failed, err: %s********", err)
		return
	}
	time.Sleep(time.Second * 3)
	Utils.PrintSmartEventByHash_Ont(sdk, hash1.ToHexString())

}
