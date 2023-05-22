package flowManage

import (
	essGolangKit "SdkTools"
	"testing"
)

// 撤销签署流程调用样例
func TestCancelFlow(t *testing.T) {
	flowId := ""
	cancelMessage := ""

	response, err := CancelFlow(essGolangKit.OperatorUserId, flowId, cancelMessage)
	if err != nil {
		t.Errorf("CancelFlow error: %v", err)
	}
	t.Logf("CancelFlow finish: %v", response.ToJsonString())
}
