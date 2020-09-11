package utils

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/payload"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/states"
	"time"
)

//init
func InterestRateInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	interestRateAddr, _ := utils.AddressFromHexString(cfg.InterestRate)
	params := []interface{}{200000000, 3000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, interestRateAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("InterestRateInit, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func ComptrollerInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	globalParamAddress, _ := utils.AddressFromHexString(cfg.GlobalParam)
	params := []interface{}{account.Address, globalParamAddress}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("ComptrollerInit, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetPriceOracle(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	oracleAddress, _ := utils.AddressFromHexString(cfg.Oracle)
	params := []interface{}{oracleAddress}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setPriceOracle", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetPriceOracle, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetWingAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	wingAddress, _ := utils.AddressFromHexString(cfg.GovToken)
	params := []interface{}{wingAddress}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setWingAddr", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetWingAddr, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetCloseFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{5000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setCloseFactor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetCloseFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetInsuranceRepayFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{5000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setInsuranceRepayFactor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetInsuranceRepayFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetMaxAssets(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{10}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setMaxAssets", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetMaxAssets, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetLiquidationIncentive(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{11000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setLiquidationIncentive", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetLiquidationIncentive, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetWingRate(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	params := []interface{}{10000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setWingRate", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetLiquidationIncentive, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SupportMarket(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	fTokenInsuranceAddr, _ := utils.AddressFromHexString(cfg.FTokenInsurance)
	params := []interface{}{fTokenAddr, fTokenInsuranceAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_supportMarket", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetLiquidationIncentive, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetCollateralFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{fTokenAddr, 5000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, comptrollerAddr, "_setCollateralFactor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetLiquidationIncentive, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func AddWingMarkets(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)

	addresses := []common.Address{fTokenAddr}
	weights := []uint64{100}

	sink := common.NewZeroCopySink(nil)
	sink.WriteString("_addWingMarkets")
	length := uint64(len(addresses))
	sink.WriteVarUint(length)
	for _, v := range addresses {
		sink.WriteAddress(v)
	}
	sink.WriteVarUint(length)
	for _, v := range weights {
		sink.WriteI128(common.I128FromUint64(v))
	}

	contract := &states.WasmContractParam{}
	contract.Address = comptrollerAddr
	//bf := bytes.NewBuffer(nil)
	argbytes := sink.Bytes()
	contract.Args = argbytes

	invokePayload := &payload.InvokeCode{
		Code: common.SerializeToBytes(contract),
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
	err := genSdk.SignToTransaction(tx, account)
	if err != nil {
		log.Errorf("AddWingMarkets, genSdk.SignToTransaction error: %s", err)
	}

	txHash, err := genSdk.SendTransaction(tx)
	if err != nil {
		log.Errorf("AddWingMarkets, genSdk.SendTransaction error: %s", err)
	}
	log.Infof("AddWingMarkets success, txHash is: %s", txHash.ToHexString())
}

func FTokenInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	tokenAddr, _ := utils.AddressFromHexString(cfg.Token)
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	globalParamAddress, _ := utils.AddressFromHexString(cfg.GlobalParam)
	interestRateAddr, _ := utils.AddressFromHexString(cfg.InterestRate)
	params := []interface{}{account.Address, tokenAddr, "ONT", comptrollerAddr, globalParamAddress, interestRateAddr,
		100000000000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("FTokenInit, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetReserveFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{5000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenAddr, "_setReserveFactor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetReserveFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetInsuranceFactor(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{4000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenAddr, "_setInsuranceFactor", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetInsuranceFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetInsuranceAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	fTokenInsuranceAddr, _ := utils.AddressFromHexString(cfg.FTokenInsurance)
	params := []interface{}{fTokenInsuranceAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenAddr, "_setInsuranceAddr", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("SetInsuranceFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func FTokenInsuranceInit(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenInsuranceAddr, _ := utils.AddressFromHexString(cfg.FTokenInsurance)
	tokenAddr, _ := utils.AddressFromHexString(cfg.Token)
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	globalParamAddress, _ := utils.AddressFromHexString(cfg.GlobalParam)
	interestRateAddr, _ := utils.AddressFromHexString(cfg.InterestRate)
	params := []interface{}{account.Address, tokenAddr, "ONT", comptrollerAddr, globalParamAddress, interestRateAddr,
		100000000000000000}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenInsuranceAddr, "init", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("FTokenInit, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func SetMarketAddr(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenInsuranceAddr, _ := utils.AddressFromHexString(cfg.FTokenInsurance)
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{fTokenAddr}
	mutTx, err := genSdk.WasmVM.NewInvokeWasmVmTransaction(cfg.GasPrice, cfg.GasLimit, fTokenInsuranceAddr, "_setMarketAddr", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	if err := signTx(genSdk, mutTx, cfg.StartNonce, account); err != nil {
		log.Error(err)
	}
	hash, err := genSdk.SendTransaction(mutTx)
	if err != nil {
		log.Errorf("FTokenInit, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

func signTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
	if nonce != 0 {
		tx.Nonce = nonce
	}
	tx.Sigs = nil
	err := sdk.SignToTransaction(tx, signer)
	if err != nil {
		return fmt.Errorf("sign tx failed, err: %s", err)
	}
	return nil
}
