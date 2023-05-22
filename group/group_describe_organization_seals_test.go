package group

import (
	ess_golang_kit "SdkTools"
	"testing"
)

/**
 * 主企业代子企业查询企业印章信息样例
 * 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
 */
func TestGroupDescribeOrganizationSeals(t *testing.T) {
	limit := 10
	// 需要代发起的子企业的企业id
	proxyOrganizationId := "********************************"
	response, err := GroupDescribeOrganizationSeals(ess_golang_kit.OperatorUserId, int64(limit), proxyOrganizationId)
	if err != nil {
		t.Errorf("DescribeOrganizationSeals error: %v", err)
	}
	t.Logf("DescribeOrganizationSeals finish: %v", response.ToJsonString())
}
