package certificateManagement

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlowEvidenceReport 创建并返回出证报告
//
// 官网文档：https://cloud.tencent.com/document/product/1323/79686
//
// 创建出证报告，返回报告 ID。
func CreateFlowEvidenceReport(userId, flowId string) (*ess.CreateFlowEvidenceReportResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowEvidenceReportRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程编号
	request.FlowId = common.StringPtr(flowId)

	return client.CreateFlowEvidenceReport(request)
}
