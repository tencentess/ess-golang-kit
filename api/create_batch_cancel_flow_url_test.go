package api

import (
	"SdkTools"
	"testing"
)

// 获取批量撤销签署流程链接样例
func TestCreateBatchCancelFlowUrl(t *testing.T) {
	// 需要执行撤回的签署流程id数组，最多100个
	flowIds := []string{"********************************", "********************************"}

	response, err := CreateBatchCancelFlowUrl(ess_golang_kit.OperatorUserId, flowIds)
	if err != nil {
		t.Errorf("CreateBatchCancelFlowUrl error: %v", err)
	}
	t.Logf("CreateBatchCancelFlowUrl finish: %v", response.ToJsonString())
}
