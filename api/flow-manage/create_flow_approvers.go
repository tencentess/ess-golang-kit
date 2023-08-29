package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateFlowApprovers(userId, flowId string,
	approvers []*ess.FillApproverInfo) (*ess.CreateFlowApproversResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowApproversRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.FlowId = common.StringPtr(flowId)

	// 补充签署人信息
	request.Approvers = approvers

	return client.CreateFlowApprovers(request)
}
