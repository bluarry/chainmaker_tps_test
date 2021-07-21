module nss/xncs/chainmaker

go 1.16

require (
	chainmaker.org/chainmaker-go/common v0.0.0
	chainmaker.org/chainmaker-sdk-go v1.2.3
	go.uber.org/zap v1.16.0
)

replace chainmaker.org/chainmaker-sdk-go => ../../chainmaker-sdk-go

replace chainmaker.org/chainmaker-go/common => ../../chainmaker-sdk-go/common
