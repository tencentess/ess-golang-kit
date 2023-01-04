package flow_manage

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestDescribeFlowInfo(t *testing.T) {
	flowIds := []*string{common.StringPtr("")}
	response, err := DescribeFlowInfo(ess_golang_kit.OperatorUserId, flowIds)
	if err != nil {
		t.Errorf("DescribeFlowInfo error: %v", err)
	}
	t.Logf("DescribeFlowInfo finish: %v", response.ToJsonString())
}
