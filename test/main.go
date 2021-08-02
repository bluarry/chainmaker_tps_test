package main

import (
	"fmt"
	"nss/xncs/chainmaker/sdkop"
	sdk "chainmaker.org/chainmaker-sdk-go"
	"sync"
	"time"
	"strconv"
	"flag"
)

func Init(){
	//1. 注册转账智能合约
	err:=sdkop.ContractInstance()
	if err!=nil{
		fmt.Printf("遇到错误,无法安装合约: %v\n",err)
		return
	}
	//2. 注册新用户
	sdkop.RegisterUser();
}
func Invoke(){
	client:=sdkop.Connect_chain(1)
	client2:=sdkop.Connect_chain(2)
	fmt.Println("====================== 3)查询钱包地址 ======================")
	addr1 := sdkop.UserContractAssetQuery(client,client2,true)  //true 为1 ，else 0
	fmt.Printf("client1 address: %s\n", addr1)
	addr2 := sdkop.UserContractAssetQuery(client,client2,false)
	fmt.Printf("client2 address: %s\n", addr2)


	fmt.Println("====================== 给用户分别发币10000000 ======================")
	amount := "10000000"
	sdkop.UserContractAssetInvoke(client,"issue_amount", amount, addr1, true)
	sdkop.UserContractAssetInvoke(client,"issue_amount", amount, addr2, true)

	fmt.Printf("%v 余额 %v\n",addr1,sdkop.GetBalance(addr1))
	fmt.Printf("%v 余额 %v\n",addr1,sdkop.GetBalance(addr2))
	sdkop.UserContractAssetInvoke(client,"transfer", "100", addr2, true)
	fmt.Printf("%v 余额 %v\n",addr1,sdkop.GetBalance(addr1))
	fmt.Printf("%v 余额 %v\n",addr1,sdkop.GetBalance(addr2))

	//fmt.Printf("hello")
}

func init(){
	
}

var Max_Count = 1000  //循环次数    每个并发循环次数
const MAX_CONNECT = 100 //连接网关数  并发数
var wg = sync.WaitGroup{}

//0 不需要，1需要
var winit=flag.Int("init",0,"是否需要初始化(安装链码，注册用户)") 
var winvoke=flag.Int("invoke",0,"是否需要测试链码调用") 
var wtest=flag.Int("test",1,"性能测试")


func invoceChaincode(client1,client2 *sdk.ChainClient){
	//addr1 := sdkop.UserContractAssetQuery(true)  //true 为node1 ，else node0
	//fmt.Printf("client1 address: %s\n", addr1)

	addr2 := sdkop.UserContractAssetQuery(client1,client2,false) 
	for i := 0; i < Max_Count; i++ {
		sdkop.UserContractAssetInvoke(client1,"transfer", "1", addr2, false) //最后一个参数为是否同步获取交易结果？
	}
	wg.Done()
}

func TpsTest(){
	clients1:=make([]*sdk.ChainClient,MAX_CONNECT)
	clients2:=make([]*sdk.ChainClient,MAX_CONNECT)

	for i:=0;i<MAX_CONNECT;i++{
		clients1[i]=sdkop.Connect_chain(1)
		clients2[i]=sdkop.Connect_chain(2)
	}
	fmt.Println("============ application-golang starts ============")

	wg.Add(MAX_CONNECT)
	for i := 0; i < MAX_CONNECT; i++ {
		go invoceChaincode(clients1[i],clients2[i])
	}
	timeStart := time.Now().UnixNano()
	wg.Wait()

	timeCount := Max_Count * MAX_CONNECT
	timeEnd := time.Now().UnixNano()
	count := float64(timeCount)
	timeResult := float64((timeEnd-timeStart)/1e6) / 1000.0
	fmt.Println("Throughput:", timeCount, "Duration:", strconv.FormatFloat(timeResult, 'g', 30, 32)+" s", "TPS:", count/timeResult)

}


func main(){
	flag.Parse()
	if 1==*winit {
		Init()
	}
	if 1==*winvoke{

		Invoke()
	}
	if 1==*wtest{
		TpsTest()
	}

	
}