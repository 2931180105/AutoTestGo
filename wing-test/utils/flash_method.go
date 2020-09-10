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
	params := []interface{}{300000000, 3000000000}
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
		log.Errorf("SetWingRate, send tx failed, err: %s********", err)
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
		log.Errorf("SupportMarket, send tx failed, err: %s********", err)
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
		log.Errorf("SetCollateralFactor, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}

type MarketMeta struct {
	Addr1 common.Address
	Addr2 common.Address
	Bool1 bool
	Bool2 bool
	U1    common.I128
	U2    common.I128
}

func (this *MarketMeta) Deserialization(source *common.ZeroCopySource) error {
	address1, eof := source.NextAddress()
	if eof {
		return fmt.Errorf("address1 deserialization error eof")
	}
	address2, eof := source.NextAddress()
	if eof {
		return fmt.Errorf("address2 deserialization error eof")
	}
	bool1, irregular, eof := source.NextBool()
	if eof || irregular {
		return fmt.Errorf("bool1 deserialization error eof")
	}
	bool2, irregular, eof := source.NextBool()
	if eof || irregular {
		return fmt.Errorf("bool2 deserialization error eof")
	}
	u1, eof := source.NextI128()
	if eof {
		return fmt.Errorf("u1 deserialization error eof")
	}
	u2, eof := source.NextI128()
	if eof {
		return fmt.Errorf("u2 deserialization error eof")
	}
	this.Addr1 = address1
	this.Addr2 = address2
	this.Bool1 = bool1
	this.Bool2 = bool2
	this.U1 = u1
	this.U2 = u2
	return nil
}

func GetMarketMeta(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	comptrollerAddr, _ := utils.AddressFromHexString(cfg.Comptroller)
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{fTokenAddr}
	r, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(comptrollerAddr, "marketMeta", params)
	if err != nil {
		fmt.Println("genSdk.WasmVM.PreExecInvokeWasmVMContract err", err)
	}
	result, err := r.Result.ToByteArray()
	if err != nil {
		fmt.Println("r.Result.ToByteArray err", err)
	}
	marketMeta := new(MarketMeta)
	source := common.NewZeroCopySource(result)
	err = marketMeta.Deserialization(source)
	if err != nil {
		fmt.Println("marketMeta.Deserialization err", err)
	}
	fmt.Println(marketMeta)
}

func GetOracle(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	oracleAddr, _ := utils.AddressFromHexString(cfg.Oracle)
	params := []interface{}{"USDT"}
	r, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(oracleAddr, "getUnderlyingPrice", params)
	if err != nil {
		fmt.Println("genSdk.WasmVM.PreExecInvokeWasmVMContract err", err)
	}
	result, err := r.Result.ToByteArray()
	if err != nil {
		fmt.Println("r.Result.ToByteArray err", err)
	}
	source := common.NewZeroCopySource(result)
	price, eof := source.NextI128()
	if eof {
		fmt.Println("source.NextI128 err", err)
	}
	fmt.Println(price.ToBigInt().Uint64())
}

func GetUnderlyingName(cfg *config.Config, account *goSdk.Account, genSdk *goSdk.OntologySdk) {
	fTokenAddr, _ := utils.AddressFromHexString(cfg.FToken)
	params := []interface{}{}
	r, err := genSdk.WasmVM.PreExecInvokeWasmVMContract(fTokenAddr, "underlyingName", params)
	if err != nil {
		fmt.Println("genSdk.WasmVM.PreExecInvokeWasmVMContract err", err)
	}
	result, err := r.Result.ToString()
	if err != nil {
		fmt.Println("r.Result.ToByteArray err", err)
	}
	fmt.Println(result)
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
	params := []interface{}{account.Address, tokenAddr, cfg.TokenName, comptrollerAddr, globalParamAddress, interestRateAddr,
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
		log.Errorf("SetInsuranceAddr, send tx failed, err: %s********", err)
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
		log.Errorf("FTokenInsuranceInit, send tx failed, err: %s********", err)
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
		log.Errorf("SetMarketAddr, send tx failed, err: %s********", err)
	} else {
		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
	}
}
