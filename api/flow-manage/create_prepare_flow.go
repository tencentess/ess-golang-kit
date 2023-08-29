package flowManage

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreatePrepareFlow(userId, flowName, resourceId string,
	approvers []*ess.FlowCreateApprover) (*ess.CreatePrepareFlowResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreatePrepareFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.ResourceId = common.StringPtr(resourceId)

	request.Approvers = approvers

	request.FlowName = common.StringPtr(flowName)

	return client.CreatePrepareFlow(request)
}
