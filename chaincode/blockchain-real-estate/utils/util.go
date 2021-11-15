package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// WriteLedger Writer
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//Creating a composite primary key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Error creating composite primary key %s", objectType, err))
	} else {
		key = val
	}
	// Serialized objects
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-Serialization of json data failed with error: %s", objectType, err))
	}
	// Write to blockchain ledger
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-Serialized json data failed to write to blockchain ledger with error: %s", objectType, err))
	}
	return nil
}

// DelLedger Delete books
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//Creating a composite primary key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Error creating composite primary key %s", objectType, err))
	} else {
		key = val
	}
	//Writing to the blockchain ledger
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-Error deleting blockchain ledger: %s", objectType, err))
	}
	return nil
}

// GetStateByPartialCompositeKeys Query data based on a composite primary key (suitable for fetching all, multiple, individual data)
// Splitting keys into queries
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// If the length of the keys passed in is 0, all data is found and returned
		// Finding relevant data from the blockchain by primary key is equivalent to a fuzzy query on the primary key
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-Error getting all data: %s", objectType, err))
		}
		defer resultIterator.Close()

		//Check if the returned data is empty, if not, iterate over the data, otherwise return the empty array
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Error in the returned data: %s", objectType, err))
			}

			results = append(results, val.GetValue())
		}
	} else {
		// The length of the keys passed in is not zero, the corresponding data is found and returned
		for _, v := range keys {
			// Create key combinations
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Error creating a key combination: %s", objectType, err))
			}
			// Getting data from the books
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Error getting data: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}
	return results, nil
}

// GetStateByPartialCompositeKeys2 根据复合主键查询数据(适合获取全部或指定的数据)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
