package utils

import (
	"encoding/hex"
	"fmt"
	config "github.com/mockyz/AutoTestGo/wing-test/config_ont"
	DbHelp "github.com/mockyz/AutoTestGo/wing-test/dbHelper"
	"github.com/ontio/ontology-crypto/keypair"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/client"
	"github.com/ontio/ontology-go-sdk/utils"
	OntCommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/types"
	"math/big"
	"time"
)

func NewAccountByWif(Wif string) (*goSdk.Account, error) {
	privateKey, err := keypair.WIF2Key([]byte(Wif))
	if err != nil {
		log.Errorf("decrypt privateKey error:%s", err)
	}
	pub := privateKey.Public()
	address := types.AddressFromPubKey(pub)
	log.Infof("address: %s\n", address.ToBase58())
	return &goSdk.Account{
		PrivateKey: privateKey,
		PublicKey:  pub,
		Address:    address,
	}, nil
}

func PrintSmartEventByHash_Ont(sdk *goSdk.OntologySdk, txHash string) {
	time.Sleep(time.Second * 3)
	//wait hash
	for j := 0; j < 50; j++ {
		time.Sleep(time.Second * 3)
		evts, err := sdk.GetSmartContractEvent(txHash)
		if err != nil || evts == nil {
			continue
		} else {
			log.Infof("evts = %s\n", evts)
			log.Infof("TxHash:%s\n", txHash)
			log.Infof("State:%d\n", evts.State)
			for _, notify := range evts.Notify {
				log.Infof("ContractAddress:%s\n", notify.ContractAddress)
				log.Infof("States:%+v\n", notify.States)
			}
			break
		}
	}

}

func SignTxAndSendTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
	if nonce != 0 {
		tx.Nonce = nonce
	}
	tx.Sigs = nil
	err := sdk.SignToTransaction(tx, signer)
	if err != nil {
		return fmt.Errorf("sign tx failed, err: %s", err)
	}
	hash, err := sdk.SendTransaction(tx)
	if err != nil {
		log.Error(err)
		return err
	}
	PrintSmartEventByHash_Ont(sdk, hash.ToHexString())
	return nil
}

func SignTx(sdk *goSdk.OntologySdk, tx *types.MutableTransaction, nonce uint32, signer goSdk.Signer) error {
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

func GetTestConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := ""
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Error(err)
	}
	wallet, err := sdk.OpenWallet("")
	if err != nil {
		log.Error(err)
	}
	account, err := wallet.GetDefaultAccount([]byte(cfg.Password))
	if err != nil {
		log.Error(err)
	}
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetMainConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := ""
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		log.Errorf("ParseConfig error:%s", err)
	}
	wallet, _ := sdk.OpenWallet("/Users[表情]yao/go[表情]c/github.com/mockyz/AutoTestGo/wing-test/WING_OTHER_OWNER.dat")
	//wallet, err := sdk.OpenWallet("/Users[表情]yao/go[表情]c/github.com/mockyz/AutoTestGo/wing-test/WING_OWNER.dat")
	if err != nil {
		log.Errorf("OpenWallet error:%s", err)
	}
	account, err := wallet.GetDefaultAccount([]byte(cfg.Password))
	if err != nil {
		log.Errorf("GetDefaultAccount error:%s", err)
	}
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GetPrvConfig() (*config.Config, *goSdk.Account, *goSdk.OntologySdk) {
	var sdk = goSdk.NewOntologySdk()
	configPath := "/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/config_testnet_02.json"
	cfg, _ := config.ParseConfig(configPath)
	wallet, _ := sdk.OpenWallet("/Users/yaoyao/go/src/github.com/mockyz/AutoTestGo/wing-test/wallet.dat")
	account, _ := wallet.GetDefaultAccount([]byte(cfg.Password))
	rpcClient := client.NewRpcClient()
	rpcClient.SetAddress(cfg.Rpc[2])
	sdk.SetDefaultClient(rpcClient)
	return cfg, account, sdk
}

func GenerateAccounts(cfg *config.Config, admin *goSdk.Account, goSdk *goSdk.OntologySdk) []*goSdk.Account {
	//accts := make([]*goSdk.Account, cfg.AccountNum)
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
		txhash, err = goSdk.Native.Ong.Transfer(cfg.GasPrice, cfg.GasLimit, admin, admin, acct.Address, 1000000000)
		if err != nil {
			log.Errorf("send ONG tx failed, err: %s********", err)
		}
		//otoken.TransferAllTestToken(cfg, admin, goSdk, acct.Address.ToBase58())
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

func GetAccounts(cfg *config.Config) []*goSdk.Account {
	accounts := DbHelp.QueryAccountFromDb(0, cfg.AccountNum)
	return accounts
}
func NewAccountToDb(wallet *goSdk.Wallet) {
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

func UpdateWingBalance(wing_balance *big.Int, base58 string) {
	db := DbHelp.SetupConnect()
	reslut, err := DbHelp.Update(db, wing_balance.Uint64(), base58)
	if err != nil {
		log.Errorf("update error: %s", err)
	}
	log.Infof("execute resulte:%s", reslut)
	db.Close()
}

func UpdateStakingBalance(wing_balance *big.Int, base58 string) {
	db := DbHelp.SetupConnect()
	reslut, err := DbHelp.UpdateStakingBalance(db, wing_balance.Uint64(), base58)
	if err != nil {
		log.Errorf("update error: %s", err)
	}
	log.Infof("execute resulte:%s", reslut)
	db.Close()
}
