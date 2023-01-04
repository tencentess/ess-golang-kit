package document_management

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CancelMultiFlowSignQRCode 取消一码多扫二维码
//
// 官网文档：https://cloud.tencent.com/document/product/1323/75451
//
// 此接口（CancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
func CancelMultiFlowSignQRCode(userId, qrCodeId string) (*ess.CancelMultiFlowSignQRCodeResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCancelMultiFlowSignQRCodeRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程id
	request.QrCodeId = common.StringPtr(qrCodeId)

	return client.CancelMultiFlowSignQRCode(request)
}
