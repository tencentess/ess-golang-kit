package personAccountManagement

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)


func DescribeThirdPartyAuthCode(userId, authCode string) (*ess.DescribeThirdPartyAuthCodeResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeThirdPartyAuthCodeRequest()
	request.BaseRequest.SetHttpMethod("POST")

	request.AuthCode = common.StringPtr(authCode)

	return client.DescribeThirdPartyAuthCode(request)
}
