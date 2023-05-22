package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CancelFlow 撤销签署流程
//
// 官网文档：https://cloud.tencent.com/document/product/1323/70362
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
// 注：如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同。
func CancelFlow(userId, flowId, cancelMessage string) (*ess.CancelFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCancelFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程id
	request.FlowId = common.StringPtr(flowId)
	// 撤销原因，最长200个字符
	request.CancelMessage = common.StringPtr(cancelMessage)

	return client.CancelFlow(request)
}
