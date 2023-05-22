package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlowSignReview 提交企业签署流程审批结果
//
// 官网地址：https://cloud.tencent.com/document/product/1323/78980
//
// 适用场景:
// 在通过接口(CreateFlow 或者CreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
func CreateFlowSignReview(userId, flowId, reviewType, reviewMessage string) (*ess.CreateFlowSignReviewResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowSignReviewRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程编号
	request.FlowId = common.StringPtr(flowId)

	// 企业内部审核结果
	//PASS: 通过
	//REJECT: 拒绝
	request.ReviewType = common.StringPtr(reviewType)

	// 审核原因
	//当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	request.ReviewMessage = common.StringPtr(reviewMessage)

	return client.CreateFlowSignReview(request)
}
