package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateFlow(userId, flowName string, approvers []*ess.FlowCreateApprover) (*ess.CreateFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程参与者信息
	request.Approvers = approvers

	request.FlowName = common.StringPtr(flowName)

	return client.CreateFlow(request)
}

// CreateFlowExtended CreateFlow接口的详细参数使用样例，前面简要调用的场景不同，此版本旨在提供可选参数的填入参考。
// 如果您在实现基础场景外有进一步的功能实现需求，可以参考此处代码。
// 注意事项：此处填入参数仅为样例，请在使用时更换为实际值。
func CreateFlowExtended() (*ess.CreateFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(essGolangKit.OperatorUserId),
	}

	request.FlowName = common.StringPtr("测试合同")

	// 构建签署方信息，注意这里的签署人类型、数量、顺序需要和模板中的设置保持一致
	request.Approvers = []*ess.FlowCreateApprover{
		{
			// 企业静默签署
			ApproverType: common.Int64Ptr(3),
		},
		{
			// 个人签署
			ApproverType: common.Int64Ptr(1),

			ApproverName: common.StringPtr("张三"),

			ApproverMobile: common.StringPtr("15912345678"),

			NotifyType: common.StringPtr("sms"),
		},
		{
			// 企业签署
			ApproverType: common.Int64Ptr(0),

			ApproverName: common.StringPtr("李四"),

			ApproverMobile: common.StringPtr("15987654321"),

			OrganizationName: common.StringPtr("XXXXX公司"),

			NotifyType: common.StringPtr("sms"),
		},
	}

	request.Unordered = common.BoolPtr(false)

	request.UserData = common.StringPtr("UserData")
	
	request.DeadLine = common.Int64Ptr(time.Now().Add(7 * 24 * time.Hour).Unix())

	return client.CreateFlow(request)
}
