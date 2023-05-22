package group

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// GroupDescribeFlowInfo 查询合同详情(DescribeFlowInfo)
//
// 官网文档：https://cloud.tencent.com/document/product/1323/80032
//
// 查询合同详情
// 适用场景：可用于主动查询某个合同详情信息。
//
// tips: 如果仅需查询合同摘要，需要使用查询合同摘要接口 https://cloud.tencent.com/document/product/1323/70358
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
