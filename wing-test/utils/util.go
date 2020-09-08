package utils

import (
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	sdkcommon "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
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

func signTx(sdk *ontology_go_sdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer ontology_go_sdk.Signer) error {
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
func getContractAddr(addr string) OntCommon.Address {
	TokenBytes, _ := OntCommon.HexToBytes(addr)
	ContractAddr, _ := utils.AddressParseFromBytes(OntCommon.ToArrayReverse(TokenBytes))
	return ContractAddr
}
