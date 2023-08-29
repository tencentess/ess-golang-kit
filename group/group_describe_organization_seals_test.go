package group

import (
	ess_golang_kit "SdkTools"
	"testing"
)


func TestGroupDescribeOrganizationSeals(t *testing.T) {
	limit := 10

	proxyOrganizationId := "********************************"
	response, err := GroupDescribeOrganizationSeals(ess_golang_kit.OperatorUserId, int64(limit), proxyOrganizationId)
	if err != nil {
		t.Errorf("DescribeOrganizationSeals error: %v", err)
	}
	t.Logf("DescribeOrganizationSeals finish: %v", response.ToJsonString())
}
