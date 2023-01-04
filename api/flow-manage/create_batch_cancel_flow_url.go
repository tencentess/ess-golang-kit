package flow_manage

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateBatchCancelFlowUrl 获取批量撤销签署流程链接
//
// 官网地址：https://cloud.tencent.com/document/product/1323/78262
//
// 电子签企业版：指定需要批量撤回的签署流程Id，获取批量撤销链接
// 客户指定需要撤回的签署流程Id，最多100个，超过100不处理；接口调用成功返回批量撤回合同的链接，通过链接跳转到电子签小程序完成批量撤回
func CreateBatchCancelFlowUrl(userId string, flowIds []string) (*ess.CreateBatchCancelFlowUrlResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateBatchCancelFlowUrlRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 需要执行撤回的签署流程id数组，最多100个
	request.FlowIds = common.StringPtrs(flowIds)

	return client.CreateBatchCancelFlowUrl(request)
}
