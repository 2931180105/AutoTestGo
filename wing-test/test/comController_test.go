package test

import (
	"github.com/mockyz/AutoTestGo/common/log"
	"github.com/ontio/ontology-go-sdk/utils"
	"io/ioutil"
	//OntCommon "github.com/ontio/ontology/common"
	"testing"
)

func TestGetBase58(t *testing.T) {
	bytes, err := ioutil.ReadFile("")
	if err != nil {
		log.Fatal(err)
	}
	contractCodeStr := string(bytes)
	Contract, err := utils.GetContractAddress(contractCodeStr)
	if err != nil {
		log.Error(err)
	}
	log.Infof("contract address : %s", Contract.ToBase58())
	log.Infof("contract address : %s", Contract.ToHexString())

	addr, _ := utils.AddressFromHexString("d034792f80deeacd983dc257d29784ea71a1d5ec")
	log.Infof("base58: %s", addr.ToBase58())
}
