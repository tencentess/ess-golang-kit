package group

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestGroupDescribeFlowInfo(t *testing.T) {

	flowIds := []*string{common.StringPtr("")}

	proxyOrganizationId := ""
	response, err := GroupDescribeFlowInfo(ess_golang_kit.OperatorUserId, flowIds, proxyOrganizationId)
	if err != nil {
		t.Errorf("DescribeFlowInfo error: %v", err)
	}
	t.Logf("DescribeFlowInfo finish: %v", response.ToJsonString())
}
