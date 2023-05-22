package organizationManagement

import (
	ess_golang_kit "SdkTools"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 查询员工信息调用样例
func TestDescribeIntegrationEmployees(t *testing.T) {
	limit := 10
	offset := 0
	filters := []*ess.Filter{{
		Key:    common.StringPtr("Status"),
		Values: common.StringPtrs([]string{"IsVerified"}),
	}}

	response, err := DescribeIntegrationEmployees(ess_golang_kit.OperatorUserId, int64(limit), int64(offset), filters)
	if err != nil {
		t.Errorf("DescribeIntegrationEmployees error: %v", err)
	}
	t.Logf("DescribeIntegrationEmployees finish: %v", response.ToJsonString())
}
