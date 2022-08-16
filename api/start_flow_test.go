package api

import (
	"SdkTools"
	"testing"
)

// 此接口用于发起流程调用样例
func TestStartFlow(t *testing.T) {
	// 签署流程编号，由CreateFlow接口返回
	flowId := "***********************"

	response, err := StartFlow(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("StartFlow error: %v", err)
	}
	t.Logf("StartFlow finish: %v", response.ToJsonString())
}
