package group

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// GroupDescribeFlowTemplates 集团账号-查询模板(DescribeFlowTemplates)
//
// 官网文档：https://cloud.tencent.com/document/product/1323/74803
//
// 适用场景：当模板较多或模板中的控件较多时，可以通过查询模板接口更方便的获取自己主体下的模板列表，以及每个模板内的控件信息。
// 该接口常用来配合“创建电子文档”接口作为前置的接口使用。
func GroupDescribeFlowTemplates(userId, templateId, proxyOrganizationId string) (*ess.DescribeFlowTemplatesResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeFlowTemplatesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}
	// 搜索条件，具体参考Filter结构体。本接口取值：template-id：按照【 模板唯一标识 】进行过滤
	request.Filters = []*ess.Filter{
		{
			// 查询过滤条件的Key
			Key: common.StringPtr("template-id"),
			// 查询过滤条件的Value列表
			Values: []*string{common.StringPtr(templateId)},
		},
	}

	return client.DescribeFlowTemplates(request)
}
