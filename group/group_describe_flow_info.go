package group

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func GroupDescribeFlowInfo(userId string, flowIds []*string, proxyOrganizationId string) (*ess.DescribeFlowInfoResponse, error) {
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeFlowInfoRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}
	request.FlowIds = flowIds

	return client.DescribeFlowInfo(request)
}
