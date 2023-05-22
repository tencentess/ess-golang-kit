package certificateManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 创建并返回出证报告调用样例
func TestCreateFlowEvidenceReport(t *testing.T) {
	flowId := "********************************"

	response, err := CreateFlowEvidenceReport(essGolangKit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("CreateFlowEvidenceReport error: %v", err)
	}
	t.Logf("CreateFlowEvidenceReport finish: %v", response.ToJsonString())
}
