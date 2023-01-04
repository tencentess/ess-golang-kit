package flow_manage

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlowApprovers 补充签署流程本企业签署人信息
//
// 官网文档：https://cloud.tencent.com/document/product/1323/80033
//
//适用场景：在通过模版或者文件发起合同时，若未指定本企业签署人信息，则流程发起后，可以调用此接口补充签署人。
//同一签署人可以补充多个员工作为候选签署人,最终签署人取决于谁先领取合同完成签署。
//
//注：目前暂时只支持补充来源于企业微信的员工作为候选签署人
func CreateFlowApprovers(userId, flowId string, approvers []*ess.FillApproverInfo) (*ess.CreateFlowApproversResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateFlowApproversRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程编号
	request.FlowId = common.StringPtr(flowId)

	// 补充签署人信息
	request.Approvers = approvers

	return client.CreateFlowApprovers(request)
}
