package utils

import (
	"encoding/hex"
	"fmt"
	ontSDK "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"math/big"
)


func InvokeWasmVMTx2(sdk *ontSDK.OntologySdk, signConfig *ontSDK.Account, tx *types.MutableTransaction) (txHash string, err error) {
	err = SignTx(sdk, tx,0, signConfig)
	if err != nil {
		return "", fmt.Errorf("InvokeTx: %s", err)
	}
	sink := common.NewZeroCopySink(nil)
	im, err := tx.IntoImmutable()
	if err != nil {
		return "", fmt.Errorf("InvokeTx: %s", err)
	}
	im.Serialization(sink)
	// log.Info(hex.EncodeToString(sink.Bytes()))
	hash, err := sdk.SendTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("InvokeTx: %s, %s", err, hex.EncodeToString(sink.Bytes()))
	}
	return hash.ToHexString(), nil
}

func InvokeNeoVMTx(sdk *ontSDK.OntologySdk, signConfig *ontSDK.Account, gasPrice, gasLimit uint64,
	contract common.Address, method string, params []interface{}) (txHash string, err error) {
	tx, err := sdk.NeoVM.NewNeoVMInvokeTransaction(gasPrice, gasLimit, contract, []interface{}{method, params})
	if err != nil {
		return "", fmt.Errorf("InvokeNeoVMTx: %s", err)
	}
	err = SignTx(sdk, tx,0, signConfig)
	if err != nil {
		return "", fmt.Errorf("InvokeNeoVMTx: %s", err)
	}
	hash, err := sdk.SendTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("InvokeNeoVMTx: %s", err)
	}
	return hash.ToHexString(), nil
}

func PreExecuteBool(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) (bool, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return false, fmt.Errorf("PreExecuteBool: %s", err)
	}
	return res.Result.ToBool()
}

func PreExecuteAddrArray(sdk *ontSDK.OntologySdk, contract common.Address, method string,
	params []interface{}) ([]common.Address, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return nil, fmt.Errorf("PreExecuteAddrArray: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("PreExecuteAddrArray: %s", err)
	}
	result := make([]common.Address, 0)
	source := common.NewZeroCopySource(data)
	length, _, ill, eof := source.NextVarUint()
	if ill {
		return nil, fmt.Errorf("PreExecuteAddrArray: read len ill")
	}
	if eof {
		return nil, fmt.Errorf("PreExecuteAddrArray: read len eof")
	}
	for i := uint64(0); i < length; i++ {
		addr, eof := source.NextAddress()
		if eof {
			return nil, fmt.Errorf("PreExecuteAddrArray: read addr eof")
		}
		result = append(result, addr)
	}
	return result, nil
}

func PreExecuteAddress(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) (common.Address, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("PreExecuteAddress: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("PreExecuteAddress: %s", err)
	}
	addr, err := common.AddressParseFromBytes(data)
	if err != nil {
		return common.ADDRESS_EMPTY, fmt.Errorf("PreExecuteAddress: %s", err)
	}
	return addr, nil
}

func PreExecuteBigInt(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) (*big.Int, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return nil, fmt.Errorf("PreExecuteBigInt: %s", err)
	}
	return res.Result.ToInteger()
}

func NeoVMPreExecuteBigInt(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) (*big.Int, error) {
	res, err := sdk.NeoVM.PreExecInvokeNeoVMContract(contract, []interface{}{method, params})
	if err != nil {
		return nil, fmt.Errorf("NeoVMPreExecuteBigInt: %s", err)
	}
	return res.Result.ToInteger()
}

func PreExecuteString(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) (string, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return "", fmt.Errorf("PreExecuteBigInt: %s", err)
	}
	return res.Result.ToString()
}

func PreExecuteStringArray(sdk *ontSDK.OntologySdk, contract common.Address, method string, params []interface{}) ([]string, error) {
	res, err := sdk.WasmVM.PreExecInvokeWasmVMContract(contract, method, params)
	if err != nil {
		return nil, fmt.Errorf("PreExecuteBigInt: %s", err)
	}
	data, err := res.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("PreExecuteStringArray: %s", err)
	}
	source := common.NewZeroCopySource(data)
	length, _, ill, eof := source.NextVarUint()
	if ill {
		return nil, fmt.Errorf("PreExecuteStringArray: read len ill")
	}
	if eof {
		return nil, fmt.Errorf("PreExecuteStringArray: read len eof")
	}
	result := make([]string, 0)
	for i := uint64(0); i < length; i++ {
		addr, _, ill, eof := source.NextString()
		if ill {
			return nil, fmt.Errorf("PreExecuteStringArray: read str ill")
		}
		if eof {
			return nil, fmt.Errorf("PreExecuteStringArray: read str eof")
		}
		result = append(result, addr)
	}
	return result, nil
}
