package personAccountManagement

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DescribeThirdPartyAuthCode 通过AuthCode查询用户是否实名
//
// 官网文档：https://cloud.tencent.com/document/product/1323/70368
func DescribeThirdPartyAuthCode(userId, authCode string) (*ess.DescribeThirdPartyAuthCodeResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeThirdPartyAuthCodeRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 电子签小程序跳转客户小程序时携带的授权查看码
	request.AuthCode = common.StringPtr(authCode)

	return client.DescribeThirdPartyAuthCode(request)
}
