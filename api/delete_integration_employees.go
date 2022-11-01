package api

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// DeleteIntegrationEmployees 移除员工
func DeleteIntegrationEmployees(userId string, employees []*ess.Staff) (*ess.DeleteIntegrationEmployeesResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewDeleteIntegrationEmployeesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 待移除员工的信息，userId和openId二选一，必填一个
	request.Employees = employees

	return client.DeleteIntegrationEmployees(request)
}
