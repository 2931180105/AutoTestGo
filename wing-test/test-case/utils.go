package test_case

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/mockyz/AutoTestGo/wing-test/compound/comptroller"
	"github.com/mockyz/AutoTestGo/wing-test/compound/ftoken"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	if_borrow "github.com/mockyz/AutoTestGo/wing-test/if-pool/if-borrow"
	if_ctrl "github.com/mockyz/AutoTestGo/wing-test/if-pool/if-ctrl"
	"github.com/mockyz/AutoTestGo/wing-test/if-pool/iftoken"
	"github.com/mockyz/AutoTestGo/wing-test/if-pool/iitoken"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)
func NewMarkets(cfg *config.Config, account *ontSDK.Account, sdk *ontSDK.OntologySdk, makretAddr string) (*ftoken.FlashToken, error) {
	comp,err := comptroller.NewComptroller(sdk,cfg.Comptroller,account,cfg.GasPrice,cfg.GasLimit)
	if err !=nil {
		log.Errorf("NewComptroller  err:%v",err)
	}
	market ,err := ftoken.NewFlashToken2(sdk,makretAddr,account,cfg,comp)
	if err !=nil {
		log.Errorf("NewFlashToken  err:%v",err)
	}
	market.TestConfig = cfg

	return market, nil
}

type  IFPool struct{
	IfCtrl   if_ctrl.Comptroller
	IfBorrow if_borrow.IfBorrowPool
	IfToken	iftoken.IFToken
	IIToken  iitoken.IIToken

}
func NewIFPool(nodeRPCAddr, comptrollerAddr, borrowAddr string, signer *ontSDK.Account, gasPrice, gasLimit uint64) *IFPool {
	sdk := ontSDK.NewOntologySdk()
	client := sdk.NewRpcClient()
	client.SetAddress(nodeRPCAddr)
	sdk.SetDefaultClient(client)
	_, err := sdk.GetCurrentBlockHeight()
	if err != nil {
		return nil
	}
	addr, err := common.AddressFromHexString(comptrollerAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(comptrollerAddr)
		if err != nil {
			return nil
		}
	}
	IfCtrl := if_ctrl.Comptroller{
		Sdk:      sdk,
		Signer:   signer,
		Addr:     addr,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
	baddr, err := common.AddressFromHexString(borrowAddr)
	if err != nil {
		addr, err = common.AddressFromBase58(borrowAddr)
		if err != nil {
			return nil
		}
	}
	IfBorrow :=if_borrow.IfBorrowPool{
		Sdk:      sdk,
		Signer:   signer,
		Addr:     baddr,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
	return &IFPool{IfCtrl: IfCtrl, IfBorrow: IfBorrow}
}