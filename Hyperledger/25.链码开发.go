//链码源文件基本结构

package main

import(
    //链码开发相关API
    github.com/hyperledger/fabric/core/chaincode/shim
    //响应消息的封装
    github.com/hyperledger/fabric/protos/peer
)

// ChaincodeInterface

type SimpleChaincode struct{

}

// 在链码被实例化或升级时被自动调用
func (t *SimpleChaincode) Init (stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

// 对链码进行调用时自动执行Invoke方法
func (t *SimpleChaincode) Invoke (stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func main(){
    //启动链码
    err := shim.Start(new(SimpleChaincode));

}

// query invoke set get




链码API

GetState(key string)([]byte error) 通过key来返回数组特定值
PutState(key string,value []byte) error 账本中写入特定的键和值
DelState
GetStateByRange(startKey,endKey string) 根据指定的范围内的键值
GetHistoryForKey(key string) 返回指定键的所有历史值
GetQueryResult(query string) 对（支持富查询功能的）状态数据库进行查询
