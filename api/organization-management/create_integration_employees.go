package organizationManagement

import (
	essTools "SdkTools"
	cService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateIntegrationEmployees 创建员工
//
// 官网文档：https://cloud.tencent.com/document/product/1323/81117
//
// 创建员工
func CreateIntegrationEmployees(userId string,
	employees []*ess.Staff) (*ess.CreateIntegrationEmployeesResponse, error) {
	// 构造客户端调用实例
	client := cService.GetClientInstance(essTools.SecretId, essTools.SecretKey, essTools.EndPoint)

	request := ess.NewCreateIntegrationEmployeesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 待创建员工的信息，Mobile和DisplayName必填
	request.Employees = employees

	return client.CreateIntegrationEmployees(request)
}
