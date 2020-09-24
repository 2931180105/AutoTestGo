package ftoken

import (
	"fmt"
	"github.com/mockyz/AutoTestGo/common/log"
	Utils "github.com/mockyz/AutoTestGo/wing-test/utils"
	goSdk "github.com/ontio/ontology-go-sdk"
	OntCommon "github.com/ontio/ontology/common"
)

//TODO： init/wing token set gov address
func GetAllBalanceOfUnderlying(address string) {
	cfg, _, sdk := Utils.GetPrvConfig()
	BalanceOfUnderlying(cfg, address, sdk, cfg.FWING)
	BalanceOfUnderlying(cfg, address, sdk, cfg.FUSDC)
	BalanceOfUnderlying(cfg, address, sdk, cfg.FRENBTC)
	BalanceOfUnderlying(cfg, address, sdk, cfg.FONT)
	BalanceOfUnderlying(cfg, address, sdk, cfg.FETH)
	BalanceOfUnderlying(cfg, address, sdk, cfg.FETH9)
}

func AllMarkets(sdk *goSdk.OntologySdk, comptrooller OntCommon.Address) {
	params := []interface{}{}
	result, err := sdk.WasmVM.PreExecInvokeWasmVMContract(comptrooller, "allMarkets", params)
	if err != nil {
		fmt.Println("construct tx err", err)
	}
	log.Infof("AllMarkets:%s", result.Result)
}
