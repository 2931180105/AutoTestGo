package utils

import (
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology/common/log"
	"testing"
)

func TestRegId(t *testing.T) {
	pwd := []byte("123456")
	var sdk = goSdk.NewOntologySdk()
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress("http://127.0.0.1:20336")
	sdk.SetDefaultClient(rpcClient)
	wallet, err := sdk.OpenWallet("../wallet.dat")
	if err != nil {
		log.Errorf("parse wallet err: %s", err)
	}
	account, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		log.Errorf("get account err: %s", err)
	}
	testId, _ := wallet.NewDefaultSettingIdentity(pwd)
	controller, _ := testId.NewDefaultSettingController("1", pwd)
	txHash, err := sdk.Native.OntId.RegIDWithPublicKey(0, 20000000, account, account, testId.ID, controller)
	if err != nil {
		log.Errorf(" RegIDWithPublicKey err: %s", err)
	}
	log.Infof("txhash %s", txHash)
}
