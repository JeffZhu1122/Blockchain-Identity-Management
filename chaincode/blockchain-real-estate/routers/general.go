package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

//Tmp comment
type Tmp struct {
	Key string  `json:"key"` 
	Id  string  `json:"idk"` 
}

// DeleteData comment
func DeleteData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key := args[0] 
	idk := args[1]

	if err := utils.DelLedger(stub, key, []string{idk}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	tmp := &Tmp{
		Key: key,
		Id: idk,
	}

	tmpByte, err := json.Marshal(tmp)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(tmpByte)
}

