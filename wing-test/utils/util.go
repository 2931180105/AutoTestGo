package utils

import (
	"encoding/hex"
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
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

func SignTx(sdk *ontology_go_sdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer ontology_go_sdk.Signer) error {
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
	//accts := make([]*ontology_go_sdk.Account, cfg.AccountNum)
	before_amount_ont, _ := goSdk.Native.Ont.BalanceOf(admin.Address)
	before_amount_ong, _ := goSdk.Native.Ong.BalanceOf(admin.Address)
	accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)
	for i := 0; i < cfg.AccountNum; i++ {
		acct := accounts[i]
		txhash, err := goSdk.Native.Ont.Transfer(cfg.GasPrice, cfg.GasLimit, admin, admin, acct.Address, cfg.Amount)
		if err != nil {
			log.Errorf("send tx failed, err: %s********", err)
		} else {
			log.Infof("send  Ont sentnum:***%d", txhash.ToHexString(), i)
		}
		txhash, err = goSdk.Native.Ong.Transfer(cfg.GasPrice, cfg.GasLimit, admin, admin, acct.Address, cfg.Amount*100000000)
		if err != nil {
			log.Errorf("send ONG tx failed, err: %s********", err)
		}
	}
	time.Sleep(time.Second * 6)
	after_amount_ont, _ := goSdk.Native.Ont.BalanceOf(admin.Address)
	after_amount_ong, _ := goSdk.Native.Ong.BalanceOf(admin.Address)
	log.Infof("before_amount_ont: %d , after_amount_ont  : %d", before_amount_ont, after_amount_ont)
	log.Infof("before_amount_ong: %d , after_amount_ong  : %d", before_amount_ong, after_amount_ong)
	log.Infof("balance change of ONT : %d", before_amount_ont-after_amount_ont)
	log.Infof("balance change of ONG : %d", before_amount_ong-after_amount_ong)
	return accounts
}
func NewAccountToDb(wallet *ontology_go_sdk.Wallet) {
	db := DbHelp.SetupConnect()
	pwd := []byte("123456")
	account, err := wallet.NewDefaultSettingAccount(pwd)
	if err != nil {
		log.Infof(" new account error : %s", err)
	}
	log.Infof("wallet account: %s", wallet.GetAccountCount())
	base58 := account.Address.ToBase58()
	hexWif := keypair.SerializePrivateKey(account.PrivateKey)
	hexWifStr := hex.EncodeToString(hexWif)
	DbHelp.Insert(db, base58, hexWifStr, 1, 2, 3)
	db.Close()
}
