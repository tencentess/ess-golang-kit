package group

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func GroupDescribeOrganizationGroupOrganizations(userId string, limit, offset uint64) (*ess.DescribeOrganizationGroupOrganizationsResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewDescribeOrganizationGroupOrganizationsRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId:   common.StringPtr(userId),
		ClientIp: common.StringPtr("8.8.8.8"),
		Channel:  common.StringPtr("YUFU"),
	}
	request.Limit = &limit
	request.Offset = &offset

	return client.DescribeOrganizationGroupOrganizations(request)
}
