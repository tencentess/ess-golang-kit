package group

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func GroupCreateSchemeUrl(userId, flowId, proxyOrganizationId string) (*ess.CreateSchemeUrlResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateSchemeUrlRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}

	request.FlowId = common.StringPtr(flowId)

	request.PathType = common.Uint64Ptr(1)

	return client.CreateSchemeUrl(request)
}
