package flowManage

import (
	essGolangKit "SdkTools"
	"testing"

	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 补充签署流程本企业签署人信息样例
func TestCreateFlowApprovers(t *testing.T) {
	// 签署流程编号
	flowId := "********************************"

	// 签署人签署Id
	recipientId := "********************************"

	// 签署人来源
	//WEWORKAPP: 企业微信
	approverSource := "********************************"

	// 企业自定义账号Id
	//WEWORKAPP场景下指企业自有应用获取企微明文的userid
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
