package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 通过AuthCode查询用户是否实名
func DescribeThirdPartyAuthCode(userId, authCode string) (*ess.DescribeThirdPartyAuthCodeResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewDescribeThirdPartyAuthCodeRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 电子签小程序跳转客户小程序时携带的授权查看码
	request.AuthCode = common.StringPtr(authCode)

	return client.DescribeThirdPartyAuthCode(request)
}
