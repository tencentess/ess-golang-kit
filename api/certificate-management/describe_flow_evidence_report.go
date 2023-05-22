package certificateManagement

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DescribeFlowEvidenceReport 查询出证报告
//
// 官网文档：https://cloud.tencent.com/document/product/1323/83441
//
// 查询出证报告，返回报告 URL。
func DescribeFlowEvidenceReport(userId, reportId string) (*ess.DescribeFlowEvidenceReportResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeFlowEvidenceReportRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 出证报告编号
	request.ReportId = common.StringPtr(reportId)

	return client.DescribeFlowEvidenceReport(request)
}
