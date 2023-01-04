package document_management

import (
	ess_golang_kit "SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// VerifyPdf 合同文件验签
//
// 官网文档：https://cloud.tencent.com/document/product/1323/80797
//
// 验证合同文件
func VerifyPdf(userId, flowId string) (*ess.VerifyPdfResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewVerifyPdfRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程id
	request.FlowId = common.StringPtr(flowId)

	return client.VerifyPdf(request)
}
