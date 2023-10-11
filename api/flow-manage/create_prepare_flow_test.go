package flowManage

import (
	essGolangKit "SdkTools"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 创建快速发起流程调用样例
func TestCreatePrepareFlow(t *testing.T) {
	flowName := "快速发起合同"
	resourceId := "********************"
	approvers := []*ess.FlowCreateApprover{
		{
			ApproverType: common.Int64Ptr(1),

			ApproverName: common.StringPtr("********************"),

			ApproverMobile: common.StringPtr("********************"),
		},
	}
	response, err := CreatePrepareFlow(essGolangKit.OperatorUserId, flowName, resourceId, approvers)
	if err != nil {
		t.Errorf("CreatePrepareFlow error: %v", err)
	}
	t.Logf("CreatePrepareFlow finish: %s", *response.Response.Url)
}
