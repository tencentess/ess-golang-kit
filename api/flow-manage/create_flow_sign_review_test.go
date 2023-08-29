package flowManage

import (
	"SdkTools"
	"testing"
)

// 提交企业签署流程审批结果样例
func TestCreateFlowSignReview(t *testing.T) {
	// 签署流程编号
	flowId := "********************************"

	reviewType := "********************************"

	reviewMessage := "********************************"

	response, err := CreateFlowSignReview(essGolangKit.OperatorUserId, flowId, reviewType, reviewMessage)
	if err != nil {
		t.Errorf("CreateFlowSignReview error: %v", err)
	}
	t.Logf("CreateFlowSignReview finish: %v", response.ToJsonString())
}
