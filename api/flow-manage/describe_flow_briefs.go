package flow_manage

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DescribeFlowBriefs 查询流程摘要
//
// 官网文档：https://cloud.tencent.com/document/product/1323/70358
//
// 适用场景：可用于主动查询某个合同流程的签署状态信息。可以配合回调通知使用。
// 日调用量默认10W
//
// tips: 如果需要查询合同的详细情况，需要使用查询合同详情接口 https://cloud.tencent.com/document/product/1323/80032
func DescribeFlowBriefs(userId string, flowIds []*string) (*ess.DescribeFlowBriefsResponse, error) {
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewDescribeFlowBriefsRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	request.FlowIds = flowIds

	return client.DescribeFlowBriefs(request)
}
