package models

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	_ "github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"github.com/hyperledger/fabric-sdk-go/def/fabapi"
	_ "github.com/hyperledger/fabric-sdk-go/def/fabapi"
)

// 维护一个与区块链操作相关的结构体
type ChainCodeSpec struct{
	client apitxn.ChannelClient
	chaincodeID string
}

// 账本数据的上传
func (this *ChainCodeSpec) ChainCodeInvoke(function string,args [][]byte)(responseArgs[]byte,err error){
	request := apitxn.ExecuteTxRequest{ChaincodeID:this.chaincodeID,Fcn:function,Args:args}
	id,err := this.client.ExecuteTx(request)
	return []byte(id.ID),err
}

// 账本数据查询
func(this *ChainCodeSpec) ChainCodeQuery(funcction string,args [][]byte)(responseArgs[]byte,err error){
	request :=apitxn.QueryRequest{ChaincodeID:this.chaincodeID,Fcn:funcction,Args:args}
	return this.client.Query(request)
}

func getSDK(conf string)(*fabapi.FabricSDK, error){
	option := fabapi.Options{ConfigFile:conf}
	sdk,err := fabapi.NewSDK(option)
	if err!=nil{
		beego.Error("fabapi.NewSDK err!")
		return nil,err
	}
	return sdk,nil
}

func Init(channelID string,userID string,ccID string)(*ChainCodeSpec,error){
	// 获取 SDK,需要给 getSDK 函数传递一个配置文件路径
	config := beego.AppConfig.String("CORE_XXX1_CONFIG_PATH")
	sdk,err := getSDK(config)
	if err!=nil{
		beego.Error("getSDK error")
		return nil,err
	}

	// 通过 SDK 创建客户端
	client,err :=sdk.NewChannelClient(channelID,userID)
	if err!=nil{
		beego.Error("sdk.NewChannelClient err!")
		return nil,err
	}

	return &ChainCodeSpec{client:client,chaincodeID:ccID},nil

}
