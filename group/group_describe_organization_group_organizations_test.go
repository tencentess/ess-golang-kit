package group

import (
	essGolangKit "SdkTools"
	"testing"
)

/**
 * 主企业代子企业查询模板信息样例
 * 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
 */
func TestGroupDescribeOrganizationGroupOrganizations(t *testing.T) {
	limit := 10
	offset := 0
	response, err := GroupDescribeOrganizationGroupOrganizations(essGolangKit.OperatorUserId, uint64(limit), uint64(offset))
	if err != nil {
		t.Errorf("DescribeOrganizationGroupOrganizations error: %v", err)
	}
	t.Logf("DescribeOrganizationGroupOrganizations finish: %v", response.ToJsonString())
}
