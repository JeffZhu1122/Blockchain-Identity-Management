package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	

	var userNames = [6]string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}

	var uname = [6]string{"admin", "manager", "agent-1", "agent-2", "agent-3", "agent-4"}
	var passwords = [6]string{"123456", "123456", "123456", "123456", "123456", "123456"}
	var phone = [6]string{"0402367888", "0402367111", "0402367222", "0402367333", "0402367444", "0402367555"}
	var email = [6]string{"admin@bci.com", "manager@bci.com", "agent-1@bci.com", "agent-2@bci.com", "agent-3@bci.com", "agent-4@bci.com"}
	var photo = [6]string{"", "", "", "", "", ""}
	var level = [6]int{0, 1, 2, 2, 2, 2}


	//初始化账号数据
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			Balance: balances[i],
			Password:  passwords[i],
			UserName:  userNames[i],
			Phone:   phone[i],
			Email: email[i],
			UName: uname[i],
			Photo: photo[i],
			Level: level[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryAccountList":
		return routers.QueryAccountList(stub, args)
	case "createRealEstate":
		return routers.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return routers.QueryRealEstateList(stub, args)
	case "createSelling":
		return routers.CreateSelling(stub, args)
	case "createSellingByBuy":
		return routers.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return routers.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return routers.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return routers.UpdateSelling(stub, args)
	case "createDonating":
		return routers.CreateDonating(stub, args)
	case "queryDonatingList":
		return routers.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return routers.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return routers.UpdateDonating(stub, args)
	case "deleteData":
		return routers.DeleteData(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
