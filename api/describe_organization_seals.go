package api

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DescribeOrganizationSeals 查询企业印章的列表，需要操作者具有查询印章权限
func DescribeOrganizationSeals(userId string, limit int64) (*ess.DescribeOrganizationSealsResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewDescribeOrganizationSealsRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.Limit = common.Int64Ptr(limit)

	return client.DescribeOrganizationSeals(request)
}
