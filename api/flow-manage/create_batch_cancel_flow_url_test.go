package flowManage

import (
	essGolangKit "SdkTools"
	"testing"
)

// 获取批量撤销签署流程链接样例
func TestCreateBatchCancelFlowUrl(t *testing.T) {

	flowIds := []string{"********************************", "********************************"}

	response, err := CreateBatchCancelFlowUrl(essGolangKit.OperatorUserId, flowIds)
	if err != nil {
		t.Errorf("CreateBatchCancelFlowUrl error: %v", err)
	}
	t.Logf("CreateBatchCancelFlowUrl finish: %v", response.ToJsonString())
}
