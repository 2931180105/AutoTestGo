package utils

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	sdkcommon "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
)

func NewAccountByWif(Wif string) (*ontology_go_sdk.Account, error) {
	privateKey, err := keypair.WIF2Key([]byte(Wif))
	if err != nil {
		log.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	log.Infof("address: %s\n", address.ToBase58())
	return &ontology_go_sdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func PrintSmartEventByHash_Ont(sdk *ontology_go_sdk.OntologySdk, txHash string) []*sdkcommon.NotifyEventInfo {
	evts, err := sdk.GetSmartContractEvent(txHash)
	if err != nil {
		fmt.Printf("GetSmartContractEvent error:%s", err)
		return nil
	}
	fmt.Printf("evts = %+v\n", evts)
	fmt.Printf("TxHash:%s\n", txHash)
	fmt.Printf("State:%d\n", evts.State)
	for _, notify := range evts.Notify {
		fmt.Printf("ContractAddress:%s\n", notify.ContractAddress)
		fmt.Printf("States:%+v\n", notify.States)
	}
	return evts.Notify
}
