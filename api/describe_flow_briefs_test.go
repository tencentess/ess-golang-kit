package api

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

func TestDescribeFlowBriefs(t *testing.T) {
	flowIds := []*string{common.StringPtr("")}
	response, err := DescribeFlowBriefs(ess_golang_kit.OperatorUserId, flowIds)
	if err != nil {
		t.Errorf("DescribeFlowBriefs error: %v", err)
	}
	t.Logf("UploadFiles finish: %v", response.ToJsonString())
}
