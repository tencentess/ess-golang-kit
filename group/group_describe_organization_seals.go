package group

import (
	ess_golang_kit "SdkTools"
	client_service "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func GroupDescribeOrganizationSeals(userId string, limit int64, proxyOrganizationId string) (*ess.DescribeOrganizationSealsResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewDescribeOrganizationSealsRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}

	request.Limit = common.Int64Ptr(limit)

	return client.DescribeOrganizationSeals(request)
}
