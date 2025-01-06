package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		SeafoodList:  []*Seafood{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 水产品上链，传入用户ID、水产品上链信息和图片信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取水产品的上链信息
	SeafoodAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将水产品的信息转换为Seafood结构体
	var seafood Seafood
	if SeafoodAsBytes != nil {
		err = json.Unmarshal(SeafoodAsBytes, &seafood)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal seafood: %v", err)
		}
	}

	// 获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	timeStr := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给水产品信息中加上溯源码
	seafood.Traceability_code = traceability_code
	// 设置图片字段
	seafood.Image = arg6 // 假设 arg6 是 Image 字段

	// 不同用户类型的上链的参数不一致
	switch userType {
	// 渔民
	case "渔民":
		// 将传入的水产品上链信息转换为Fisherman_input结构体
		seafood.Fisherman_input.Sf_seafoodName = arg1
		seafood.Fisherman_input.Sf_origin = arg2
		seafood.Fisherman_input.Sf_salvageTime = arg3
		seafood.Fisherman_input.Sf_outOfWaterTime = arg4
		seafood.Fisherman_input.Sf_fishermanName = arg5
		seafood.Fisherman_input.Sf_Txid = txid
		seafood.Fisherman_input.Sf_Timestamp = timeStr
		seafood.Fisherman_input.Image = arg6 // 设置子结构体的 Image 字段
	// 加工厂
	case "加工厂":
		// 将传入的水产品上链信息转换为Factory_input结构体
		seafood.Factory_input.Fac_productName = arg1
		seafood.Factory_input.Fac_productionbatch = arg2
		seafood.Factory_input.Fac_productionTime = arg3
		seafood.Factory_input.Fac_factoryName = arg4
		seafood.Factory_input.Fac_contactNumber = arg5
		seafood.Factory_input.Fac_Txid = txid
		seafood.Factory_input.Fac_Timestamp = timeStr
		seafood.Factory_input.Image = arg6 // 设置子结构体的 Image 字段
	// 物流司机信息
	case "物流司机信息":
		// 将传入的水产品上链信息转换为Driver_input结构体
		seafood.Driver_input.Dr_name = arg1
		seafood.Driver_input.Dr_age = arg2
		seafood.Driver_input.Dr_phone = arg3
		seafood.Driver_input.Dr_carNumber = arg4
		seafood.Driver_input.Dr_transport = arg5
		seafood.Driver_input.Dr_Txid = txid
		seafood.Driver_input.Dr_Timestamp = timeStr
		seafood.Driver_input.Image = arg6 // 设置子结构体的 Image 字段
	// keyword_29
	case "keyword_29":
		// 将传入的水产品上链信息转换为Shop_input结构体
		seafood.Shop_input.Sh_storeTime = arg1
		seafood.Shop_input.Sh_sellTime = arg2
		seafood.Shop_input.Sh_shopName = arg3
		seafood.Shop_input.Sh_shopAddress = arg4
		seafood.Shop_input.Sh_shopPhone = arg5
		seafood.Shop_input.Sh_Txid = txid
		seafood.Shop_input.Sh_Timestamp = timeStr
		seafood.Shop_input.Image = arg6 // 设置子结构体的 Image 字段
	}

	// 将水产品的信息转换为json格式
	seafoodAsBytes, err := json.Marshal(seafood)
	if err != nil {
		return "", fmt.Errorf("failed to marshal seafood: %v", err)
	}
	// 将水产品的信息存入区块链
	err = ctx.GetStub().PutState(traceability_code, seafoodAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put seafood: %v", err)
	}
	// 将水产品存入用户的水产品列表
	err = s.AddSeafood(ctx, userID, &seafood)
	if err != nil {
		return "", fmt.Errorf("failed to add seafood to user: %v", err)
	}

	return txid, nil
}

// 添加水产品到用户的水产品列表
func (s *SmartContract) AddSeafood(ctx contractapi.TransactionContextInterface, userID string, seafood *Seafood) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.SeafoodList = append(user.SeafoodList, seafood)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取水产品的上链信息
func (s *SmartContract) GetSeafoodInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Seafood, error) {
	SeafoodAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Seafood{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Seafood结构体
	var seafood Seafood
	if SeafoodAsBytes != nil {
		err = json.Unmarshal(SeafoodAsBytes, &seafood)
		if err != nil {
			return &Seafood{}, fmt.Errorf("failed to unmarshal seafood: %v", err)
		}
		if seafood.Traceability_code != "" {
			return &seafood, nil
		}
	}
	return &Seafood{}, fmt.Errorf("the seafood %s does not exist", traceability_code)
}

// 获取用户的水产品ID列表
func (s *SmartContract) GetSeafoodList(ctx contractapi.TransactionContextInterface, userID string) ([]*Seafood, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.SeafoodList, nil
}

// 获取所有的水产品信息
func (s *SmartContract) GetAllSeafoodInfo(ctx contractapi.TransactionContextInterface) ([]Seafood, error) {
	seafoodListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer seafoodListAsBytes.Close()
	var seafoods []Seafood
	for seafoodListAsBytes.HasNext() {
		queryResponse, err := seafoodListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var seafood Seafood
		err = json.Unmarshal(queryResponse.Value, &seafood)
		if err != nil {
			return nil, err
		}
		// 过滤非水产品的信息
		if seafood.Traceability_code != "" {
			seafoods = append(seafoods, seafood)
		}
	}
	return seafoods, nil
}

// 获取水产品上链历史
func (s *SmartContract) GetSeafoodHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", traceability_code)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var seafood Seafood
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &seafood)
			if err != nil {
				return nil, err
			}
		} else {
			seafood = Seafood{
				Traceability_code: traceability_code,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
		formattedTime := timestamp.Format("2024-12-01 22:11:00")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &seafood,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}
