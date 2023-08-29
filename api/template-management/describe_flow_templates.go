package templateManagement

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func DescribeFlowTemplates(userId, templateId string) (*ess.DescribeFlowTemplatesResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeFlowTemplatesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.Filters = []*ess.Filter{
		{

			Key: common.StringPtr("template-id"),

			Values: []*string{common.StringPtr(templateId)},
		},
	}

	return client.DescribeFlowTemplates(request)
}
