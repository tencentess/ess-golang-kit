package organization_management

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

// 创建员工调用样例
func TestCreateIntegrationEmployees(t *testing.T) {

	employees := []*ess.Staff{{
		Mobile:      common.StringPtr("*************"),
		DisplayName: common.StringPtr("张三"),
	}}

	response, err := CreateIntegrationEmployees(ess_golang_kit.OperatorUserId, employees)
	if err != nil {
		t.Errorf("CreateIntegrationEmployees error: %v", err)
	}
	t.Logf("CreateIntegrationEmployees finish: %v", response.ToJsonString())
}
