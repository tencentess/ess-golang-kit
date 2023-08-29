package flowManage

import (
	essGolangKit "SdkTools"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 创建签署流程调用样例
func TestCreateFlow(t *testing.T) {
	flowName := "我的第一份模板合同"
	approvers := []*ess.FlowCreateApprover{
			ApproverType: common.Int64Ptr(1),
			ApproverName: common.StringPtr("********************************"),
			ApproverMobile: common.StringPtr("********************************"),
		},
	}
	response, err := CreateFlow(essGolangKit.OperatorUserId, flowName, approvers)
	if err != nil {
		t.Errorf("CreateFlow error: %v", err)
	}
	t.Logf("CreateFlow finish: %v", response.ToJsonString())
}
