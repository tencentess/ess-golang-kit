package seal_manage

import (
	ess_golang_kit "SdkTools"
	"testing"
)

// 查询企业印章的列表
func TestDescribeOrganizationSeals(t *testing.T) {
	limit := 10

	response, err := DescribeOrganizationSeals(ess_golang_kit.OperatorUserId, int64(limit))
	if err != nil {
		t.Errorf("DescribeOrganizationSeals error: %v", err)
	}
	t.Logf("DescribeOrganizationSeals finish: %v", response.ToJsonString())
}
