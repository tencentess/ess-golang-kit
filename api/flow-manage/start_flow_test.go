package flowManage

import (
	essGolangKit "SdkTools"
	"testing"
)

// 此接口用于发起流程调用样例
func TestStartFlow(t *testing.T) {

	flowId := "***********************"

	response, err := StartFlow(essGolangKit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("StartFlow error: %v", err)
	}
	t.Logf("StartFlow finish: %v", response.ToJsonString())
}
