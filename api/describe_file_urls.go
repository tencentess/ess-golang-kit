package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 查询文件下载URL
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
func DescribeFileUrls(userId, flowId string) (*ess.DescribeFileUrlsResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

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
	request.BusinessType = common.StringPtr("FLOW")

	// 业务编号的数组，如模板编号、文档编号、印章编号
	// 最大支持20个资源
	request.BusinessIds = []*string{common.StringPtr(flowId)}

	return client.DescribeFileUrls(request)
}
