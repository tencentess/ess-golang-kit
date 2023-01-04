package flow_manage

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

// 创建快速发起流程调用样例
func TestCreatePrepareFlow(t *testing.T) {
	flowName := "快速发起合同"
	resourceId := "********************"
	approvers := []*ess.FlowCreateApprover{
		// 签署流程参与者信息
		// 个人签署方
		{
			// 参与者类型：
			// 0：企业
			// 1：个人
			// 3：企业静默签署
			// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
			ApproverType: common.Int64Ptr(1),
			// 本环节需要操作人的名字
			ApproverName: common.StringPtr("********************"),
			// 本环节需要操作人的手机号
			ApproverMobile: common.StringPtr("********************"),
		},
	}
	response, err := CreatePrepareFlow(ess_golang_kit.OperatorUserId, flowName, resourceId, approvers)
	if err != nil {
		t.Errorf("CreatePrepareFlow error: %v", err)
	}
	t.Logf("CreatePrepareFlow finish: %s", *response.Response.Url)
}
