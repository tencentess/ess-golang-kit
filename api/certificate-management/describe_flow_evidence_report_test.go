package certificateManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 查询出证报告调用样例
func TestDescribeFlowEvidenceReport(t *testing.T) {
	reportId := "********************************"

	response, err := DescribeFlowEvidenceReport(essGolangKit.OperatorUserId, reportId)
	if err != nil {
		t.Errorf("DescribeFlowEvidenceReport error: %v", err)
	}
	t.Logf("CreateFlowEvidenceReport finish: %v", response.ToJsonString())
}
