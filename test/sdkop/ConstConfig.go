package sdkop
import (
	sdk "chainmaker.org/chainmaker-sdk-go"
	"fmt"
) 
	
const (
	sdkConfigPath  = "./sdk_config.yml"
	sdkConfigPath2  = "./sdk_config2.yml"
	chainId        = "chain1"
	orgId1         = "wx-org1.chainmaker.org"
	orgId2         = "wx-org2.chainmaker.org"
	orgId3         = "wx-org3.chainmaker.org"
	orgId4         = "wx-org4.chainmaker.org"
	orgNodeId1     = "node1"
	orgNodeId2     = "node2"
	orgNodeId3     = "node3"
	orgNodeId4     = "node4"
	tlsHostName    = "chainmaker.org"
	certPathPrefix = "../build"
	adminKeyPath = certPathPrefix + "/config/%s/certs/user/admin1/admin1.tls.key"
	adminCrtPath = certPathPrefix + "/config/%s/certs/user/admin1/admin1.tls.crt"

	userKeyPath = certPathPrefix + "/config/%s/certs/user/client1/client1.tls.key"
	userCrtPath = certPathPrefix + "/config/%s/certs/user/client1/client1.tls.crt"



	nodeAddr1 = "127.0.0.1:12301"
	connCnt1  = 5

	nodeAddr2 = "127.0.0.1:12302"
	connCnt2  = 5
	certPathFormat = "/config/%s/certs/ca"
	
)

var (
	node1 *sdk.NodeConfig
	node2 *sdk.NodeConfig
	caPaths = []string{
		certPathPrefix + fmt.Sprintf(certPathFormat, orgNodeId1),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgNodeId2),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgNodeId3),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgNodeId4),
	}
	
	assetContractName = "asset001"
	assetVersion      = "1.0.0"
	assetByteCodePath = "../chiancode/asset-wasm-demo/rust-asset-management-1.0.0.wasm"
	createContractTimeout = int64(50)
	
)