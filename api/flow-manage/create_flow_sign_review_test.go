package flowManage

import (
	"SdkTools"
	"testing"
)

// 提交企业签署流程审批结果样例
func TestCreateFlowSignReview(t *testing.T) {
	// 签署流程编号
	flowId := "********************************"

	// 企业内部审核结果
	//PASS: 通过
	//REJECT: 拒绝
	reviewType := "********************************"

	// 审核原因
	//当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	reviewMessage := "********************************"

	response, err := CreateFlowSignReview(essGolangKit.OperatorUserId, flowId, reviewType, reviewMessage)
	if err != nil {
		t.Errorf("CreateFlowSignReview error: %v", err)
	}
	t.Logf("CreateFlowSignReview finish: %v", response.ToJsonString())
}
