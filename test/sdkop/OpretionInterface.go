package sdkop

import (
	"fmt"
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
)

func Connect_chain(tp int) *sdk.ChainClient{
	
	chainClient, err :=createClientWithConfig(sdkConfigPath)
	if 2==tp{
		chainClient, err =createClientWithConfig(sdkConfigPath2)
	}
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	return chainClient
}

/*
实例化代币合约
*/
func ContractInstance(){
	chainClient, err :=createClientWithConfig(sdkConfigPath)
	if err!=nil{
		fmt.Println(err)
		return
	}
	Client1=chainClient
	//fmt.Printf("%v\n",chainClient)
	//fmt.Println("====================== create admin1 ======================")
	admin1, err := createAdminWithConfig(orgNodeId1)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("====================== create admin2 ======================")
	admin2, err := createAdminWithConfig(orgNodeId2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("====================== create admin3 ======================")
	admin3, err := createAdminWithConfig(orgNodeId3)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("====================== create admin4 ======================")
	admin4, err := createAdminWithConfig(orgNodeId4)
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("====================== 1)安装钱包合约 ======================")
	pairs := []*common.KeyValuePair{
		{
			Key:   "issue_limit",
			Value: "100000000",
		},
		{
			Key:   "total_supply",
			Value: "100000000",
		},
	}
	userContractAssetCreate(chainClient, admin1, admin2, admin3, admin4, pairs, true, true)
	fmt.Println("====================== 安装钱包合约成功 ======================")
}
/*---------------上边是通用的，下边和智能合约相关----------------*/
func RegisterUser(){
	client2, err := createClientWithConfig(sdkConfigPath2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	Client2=client2
	fmt.Println("====================== 2)注册另一个用户 ======================")
	userContractAssetInvokeRegister(client2, "register", true)
   	fmt.Println("====================== 注册另一个用户成功 ======================")
   
}	
func UserContractAssetQuery(client1,client2 *sdk.ChainClient,id bool) string {
	/*
	client *sdk.ChainClient, method string, params map[string]string
	*/
	method:="query_address"
	var params map[string]string
	resp, err := client2.QueryContract(assetContractName, method, params, -1)
	if id{
		resp, err = client1.QueryContract(assetContractName, method, params, -1)
	}

	if err!=nil{
		fmt.Printf("get error: %+v\n",err);
		return ""
	}
	//fmt.Printf("QUERY asset contract [%s] resp: %+v\n", method, resp)

	err = sdk.CheckProposalRequestResp_ext(resp, true)
	if err!=nil{
		fmt.Printf("check get error: %+v\n",err);
		return ""
	}
	return string(resp.ContractResult.Result)
}


func UserContractAssetInvoke(client *sdk.ChainClient, method, amount, addr string, withSyncResult bool){
	err:=userContractAssetInvoke(client,method,amount,addr,withSyncResult)
	if err!=nil{
		fmt.Printf("invoke error : %v\n",err);
	}
}

func GetBalance(addr string)int32{
	return getBalance(addr)
}