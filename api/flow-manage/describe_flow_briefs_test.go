package flowManage

import (
	ess_golang_kit "SdkTools"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func TestDescribeFlowBriefs(t *testing.T) {
	flowIds := []*string{common.StringPtr("")}
	response, err := DescribeFlowBriefs(ess_golang_kit.OperatorUserId, flowIds)
	if err != nil {
		t.Errorf("DescribeFlowBriefs error: %v", err)
	}
	t.Logf("DescribeFlowBriefs finish: %v", response.ToJsonString())
}
