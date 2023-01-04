package organization_management

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

// 移除员工调用样例
func TestDeleteIntegrationEmployees(t *testing.T) {

	employees := []*ess.Staff{{
		UserId: common.StringPtr("************"),
	}}

	response, err := DeleteIntegrationEmployees(ess_golang_kit.OperatorUserId, employees)
	if err != nil {
		t.Errorf("DeleteIntegrationEmployees error: %v", err)
	}
	t.Logf("DeleteIntegrationEmployees finish: %v", response.ToJsonString())
}
