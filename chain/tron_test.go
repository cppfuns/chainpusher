package chain

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tc "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
)

func TestTronTransactionTransferContract(t *testing.T) {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	transaction, err := client.GetTransactionByID("68fc9d0cd12807e0d0ef25dba843c1e0a06f34b72e102b8ed3be051bda3c989a")
	if err != nil {
		t.Error(err)
	}

	contract := transaction.RawData.GetContract()[0]

	var transferContract core.TransferContract
	contract.Parameter.UnmarshalTo(&transferContract)

	payer := "TBREsCfBdPyD612xZnwvGPux7osbXvtzLh"
	payee := "TYqDptDgPUBwhrETDqJjjrCocKgwrQ5tyw"
	contractType := "type.googleapis.com/protocol.TransferContract"

	var amount int64 = 20000000

	if transferContract.Amount != amount {
		t.Error("Amount is incorrect")
	}

	if payer != tc.EncodeCheck(transferContract.OwnerAddress) {
		t.Error("Payer is incorrect")
	}

	if payee != tc.EncodeCheck(transferContract.ToAddress) {
		t.Error("Payee is incorrect")
	}

	if contractType != contract.GetParameter().TypeUrl {
		t.Error("Contract type is incorrect")
	}
}

// test the transaction of tron
func TestTronTransaction(t *testing.T) {

	expectJson := `{"id":"W/GjEYDJYiaVKSXjG4U5CvbKWtBikf7Juo1cQBffUQw=","fee":345000,"blockNumber":62207075,"blockTimeStamp":1717264812000,"contractResult":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="],"contract_address":"QaYU+AO2/XgJhqQseOycf3fm3tE8","receipt":{"energy_usage":31895,"energy_usage_total":31895,"net_fee":345000,"result":1,"energy_penalty_total":17245},"log":[{"address":"phT4A7b9eAmGpCx47Jx/d+be0Tw=","topics":["3fJSrRviyJtpwrBo/DeNqpUrp/FjxKEWKPVaTfUjs+8=","AAAAAAAAAAAAAAAAcIv0/OSYo2A3ZPO6SPzwNT+bmdM=","AAAAAAAAAAAAAAAAhnyR0himtkQMSVpaR09soe2AEMU="],"data":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZcHOA="}]}`

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	transaction, err := client.GetTransactionInfoByID("5bf1a31180c96226952925e31b85390af6ca5ad06291fec9ba8d5c4017df510c")
	if err != nil {
		t.Error(err)
	}
	addr := transaction.ContractAddress
	addrHex := hex.EncodeToString(addr)
	addrBase58 := address.Address(addr).String()

	if len(addr) != 21 {
		t.Error("Contract address length is incorrect")
	}

	if addrHex != "41a614f803b6fd780986a42c78ec9c7f77e6ded13c" {
		t.Error("Contract address is incorrect")
	}

	if addrBase58 != "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" {
		t.Error("Contract address is incorrect")
	}

	jsonBytes, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}

	if string(jsonBytes) != expectJson {
		t.Error("Transaction is incorrect")
	}
}

// PLAN: application logic
func TestTronParseArgumentsOfSmartContract(t *testing.T) {
	// payee := "TNEJn1gqWKbNo26TsimuazmxxpjqufdS83"
	payer := "TVfDGG67P8778zZKmgaMnU1sYegfwsQaWg"

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())

	input, _ := hex.DecodeString("a9059cbb000000000000000000000041d7fb3f45980187dbbe41c30fbba76f008e16e89f0000000000000000000000000000000000000000000000000000000000000186")
	if len(input) != 68 {
		t.Error("Input length is incorrect")
	}
	contractAbi, err := GetUsdtSmartContract(client)

	if err != nil {
		return
	}

	args, err := abi.GetInputsParser(contractAbi, "transfer")
	if err != nil {
		t.Error(err)
	}
	unpacked, err := args.Unpack(input[4:])

	ethAddress := unpacked[0].(common.Address)
	amount := unpacked[1].(*big.Int)

	tronAddress := ToTronAddress(ethAddress).String()

	if err != nil {
		t.Error(err)
	}

	if tronAddress != payer {
		t.Error("Address is incorrect")
	}

	if amount.Int64() != 390 {
		t.Error("Amount is incorrect")
	}
}
