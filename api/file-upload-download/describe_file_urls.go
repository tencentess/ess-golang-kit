package fileUploadDownload

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DescribeFileUrls 查询文件下载URL
//
// 官网文档：https://cloud.tencent.com/document/product/1323/70366
//
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
func DescribeFileUrls(userId, flowId string) (*ess.DescribeFileUrlsResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeFileUrlsRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 文件对应的业务类型，目前支持：
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章 “SEAL”
	// - 流程 "FLOW"
	// 一般情况下，我们下载合同文件选择使用 FLOW
	request.BusinessType = common.StringPtr("FLOW")

	// 业务编号的数组，如模板编号、文档编号、印章编号
	// 最大支持20个资源
	request.BusinessIds = []*string{common.StringPtr(flowId)}

	return client.DescribeFileUrls(request)
}
