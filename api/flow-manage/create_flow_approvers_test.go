package flowManage

import (
	essGolangKit "SdkTools"
	"testing"

	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 补充签署流程本企业签署人信息样例
func TestCreateFlowApprovers(t *testing.T) {
	flowId := "********************************"

	recipientId := "********************************"

	approverSource := "********************************"

	customUserId := "********************************"

	approver := ess.FillApproverInfo{
		RecipientId:    &recipientId,
		ApproverSource: &approverSource,
		CustomUserId:   &customUserId,
	}

	approvers := []*ess.FillApproverInfo{&approver}

	response, err := CreateFlowApprovers(essGolangKit.OperatorUserId, flowId, approvers)
	if err != nil {
		t.Errorf("CreateFlowApprovers error: %v", err)
	}
	t.Logf("CreateFlowApprovers finish: %v", response.ToJsonString())
}
