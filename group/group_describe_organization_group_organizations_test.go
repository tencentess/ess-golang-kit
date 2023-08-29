package group

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestGroupDescribeOrganizationGroupOrganizations(t *testing.T) {
	limit := 10
	offset := 0
	response, err := GroupDescribeOrganizationGroupOrganizations(essGolangKit.OperatorUserId, uint64(limit), uint64(offset))
	if err != nil {
		t.Errorf("DescribeOrganizationGroupOrganizations error: %v", err)
	}
	t.Logf("DescribeOrganizationGroupOrganizations finish: %v", response.ToJsonString())
}
