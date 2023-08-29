package organizationManagement

import (
	essTools "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func DeleteIntegrationEmployees(userId string,
	employees []*ess.Staff) (*ess.DeleteIntegrationEmployeesResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essTools.SecretId, essTools.SecretKey, essTools.EndPoint)

	request := ess.NewDeleteIntegrationEmployeesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.Employees = employees

	return client.DeleteIntegrationEmployees(request)
}
