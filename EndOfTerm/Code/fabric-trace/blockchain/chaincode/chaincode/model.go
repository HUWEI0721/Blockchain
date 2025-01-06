package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
野生水产品列表
*/
type User struct {
	UserID       string     `json:"userID"`
	UserType     string     `json:"userType"`
	RealInfoHash string     `json:"realInfoHash"`
	SeafoodList  []*Seafood `json:"seafoodList"`
}

/*
定义野生水产品结构体
溯源码
渔民输入
加工厂输入
物流司机信息输入
keyword_29输入
图片字段
*/
type Seafood struct {
	Traceability_code string          `json:"traceability_code"`
	Fisherman_input   Fisherman_input `json:"fisherman_input"`
	Factory_input     Factory_input   `json:"factory_input"`
	Driver_input      Driver_input    `json:"driver_input"`
	Shop_input        Shop_input      `json:"shop_input"`
	Image             string          `json:"image"` // 添加的图片字段
}

// HistoryQueryResult 结构体用于处理历史查询结果
type HistoryQueryResult struct {
	Record    *Seafood `json:"record"`
	TxId      string   `json:"txId"`
	Timestamp string   `json:"timestamp"`
	IsDelete  bool     `json:"isDelete"`
}

/*
渔民
野生水产品的溯源码，一物一码，主打高端市场（自动生成）
水产品名称
捕捞区域
捕捞起始时间
出水时间
渔民名称
*/
type Fisherman_input struct {
	Sf_seafoodName    string `json:"sf_seafoodName"`
	Sf_origin         string `json:"sf_origin"`
	Sf_salvageTime    string `json:"sf_salvageTime"`
	Sf_outOfWaterTime string `json:"sf_outOfWaterTime"`
	Sf_fishermanName  string `json:"sf_fishermanName"`
	Sf_Txid           string `json:"sf_txid"`
	Sf_Timestamp      string `json:"sf_timestamp"`
	Image             string `json:"image"` // 添加的图片字段
}

/*
加工厂
水产品商品名称
加工批次
出厂时间（可以防止黑心商家修改时间）
加工厂名称与地址
物流司机联系电话
*/
type Factory_input struct {
	Fac_productName     string `json:"fac_productName"`
	Fac_productionbatch string `json:"fac_productionbatch"`
	Fac_productionTime  string `json:"fac_productionTime"`
	Fac_factoryName     string `json:"fac_factoryName"`
	Fac_contactNumber   string `json:"fac_contactNumber"`
	Fac_Txid            string `json:"fac_txid"`
	Fac_Timestamp       string `json:"fac_timestamp"`
	Image               string `json:"image"` // 添加的图片字段
}

/*
物流司机信息
姓名
冷链车厢温度
电话
运输车辆车牌号
运输与冷链记录
*/
type Driver_input struct {
	Dr_name      string `json:"dr_name"`
	Dr_age       string `json:"dr_temperature"`
	Dr_phone     string `json:"dr_phone"`
	Dr_carNumber string `json:"dr_carNumber"`
	Dr_transport string `json:"dr_transport"`
	Dr_Txid      string `json:"dr_txid"`
	Dr_Timestamp string `json:"dr_timestamp"`
	Image        string `json:"image"` // 添加的图片字段
}

/*
keyword_29
存入时间
销售时间
销售点名称
销售点位置
销售点电话
*/
type Shop_input struct {
	Sh_storeTime   string `json:"sh_storeTime"`
	Sh_sellTime    string `json:"sh_sellTime"`
	Sh_shopName    string `json:"sh_shopName"`
	Sh_shopAddress string `json:"sh_shopAddress"`
	Sh_shopPhone   string `json:"sh_shopPhone"`
	Sh_Txid        string `json:"sh_txid"`
	Sh_Timestamp   string `json:"sh_timestamp"`
	Image          string `json:"image"` // 添加的图片字段
}
