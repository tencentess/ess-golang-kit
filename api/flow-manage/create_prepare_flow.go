package flow_manage

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreatePrepareFlow 创建快速发起流程
//
// 官网文档：https://cloud.tencent.com/document/product/1323/83412
//
// 适用场景：用户通过API 合同文件及签署信息，并可通过我们返回的URL在页面完成签署控件等信息的编辑与确认，快速发起合同.
// 注：该接口文件的resourceId 是通过上传文件之后获取的。
func CreatePrepareFlow(userId, flowName, resourceId string, approvers []*ess.FlowCreateApprover) (*ess.CreatePrepareFlowResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreatePrepareFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 资源Id,通过上传uploadFiles接口获得
	request.ResourceId = common.StringPtr(resourceId)
	// 签署流程参与者信息
	request.Approvers = approvers
	// 签署流程名称,最大长度200个字符
	request.FlowName = common.StringPtr(flowName)

	return client.CreatePrepareFlow(request)
}
