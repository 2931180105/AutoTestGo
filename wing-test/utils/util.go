package utils

import (
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	sdkcommon "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"time"
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

func GenerateAccounts(cfg *config.Config, admin *ontology_go_sdk.Account, goSdk *ontology_go_sdk.OntologySdk) []*ontology_go_sdk.Account {
	pwd := []byte("123456")
	wallet, _ := goSdk.CreateWallet("tmp2.dat")
	accts := make([]*ontology_go_sdk.Account, cfg.AccountNum)
	before_amount_ont, _ := goSdk.Native.Ont.BalanceOf(admin.Address)
	before_amount_ong, _ := goSdk.Native.Ong.BalanceOf(admin.Address)
	for i := 0; i < cfg.AccountNum; i++ {
		acct, _ := wallet.NewDefaultSettingAccount(pwd)
		txhash, err := goSdk.Native.Ont.Transfer(cfg.GasPrice, cfg.GasLimit, admin, admin, acct.Address, cfg.Amount)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send  Ont sentnum:***%d", txhash.ToHexString(), i)
		}
		//time.Sleep(time.Second * 3)
		//PrintSmartEventByHash_Ont(goSdk,txHash.ToHexString())
		txhash, err = goSdk.Native.Ong.Transfer(cfg.GasPrice, cfg.GasLimit, admin, admin, acct.Address, cfg.Amount*100000000)
		if err != nil {
			log.Errorf("send ONG tx failed, err: %s********", err)
		}
		accts[i] = acct
	}
	wallet.Save()
	time.Sleep(time.Second * 6)
	after_amount_ont, _ := goSdk.Native.Ont.BalanceOf(admin.Address)
	after_amount_ong, _ := goSdk.Native.Ong.BalanceOf(admin.Address)
	log.Infof("before_amount_ont: %d , after_amount_ont  : %d", before_amount_ont, after_amount_ont)
	log.Infof("before_amount_ong: %d , after_amount_ong  : %d", before_amount_ong, after_amount_ong)
	log.Infof("balance change of ONT : %d", before_amount_ont-after_amount_ont)
	log.Infof("balance change of ONG : %d", before_amount_ong-after_amount_ong)
	return accts
}
