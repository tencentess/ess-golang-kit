package certificate_management

import (
	"SdkTools"
	"testing"
)

// 创建并返回出证报告调用样例
func TestCreateFlowEvidenceReport(t *testing.T) {
	flowId := "********************************"

	response, err := CreateFlowEvidenceReport(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("CreateFlowEvidenceReport error: %v", err)
	}
	t.Logf("CreateFlowEvidenceReport finish: %v", response.ToJsonString())
}
