package utils

import (
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
)

//txdata,err :=sdk.GetTransaction("1b72fa611f4dfe65efa2797051cd7b7143da2c0737cb6b30919360f03e20d91a")
//log.Infof("txdata is %x ", txdata)
//Utils "github.com/mockyz/AutoTestGo/ont-test/utils"

func BalanceOfONG(sdk *goSdk.OntologySdk, address common.Address) {
	balance, err := sdk.Native.Ong.BalanceOf(address)
	if err != nil {
		log.Errorf("get account err: %s", err)
		return
	}
	log.Infof("account %s is balance of %d", address.ToBase58(), balance)
}

func BalanceOfONT(sdk *goSdk.OntologySdk, address common.Address) {
	balance, err := sdk.Native.Ont.BalanceOf(address)

	if err != nil {
		log.Errorf("get account err: %s", err)
		return
	}
	log.Infof("account %s ONG  balance of %d", address.ToBase58(), balance)
}
