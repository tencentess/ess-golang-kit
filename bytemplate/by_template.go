package bytemplate

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 构造签署人 - 以B2B2C为例, 实际请根据自己的场景构造签署方
func BuildFlowCreateApprovers() []*ess.FlowCreateApprover {
	// 个人签署方构造参数
	personName := "********************"
	personMobile := "********************"

	// 企业签署方构造参数
	organizationName := "********************"
	organizationUserName := "********************"
	organizationUserMobile := "********************"

	approvers := []*ess.FlowCreateApprover{
		BuildServerSignFlowCreateApprover(),
		BuildOrganizationFlowCreateApprover(organizationUserName, organizationUserMobile, organizationName),
		BuildPersonFlowCreateApprover(personName, personMobile),
	}
	return approvers
}

// 打包个人签署方参与者信息
func BuildPersonFlowCreateApprover(name, mobile string) *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		ApproverType: common.Int64Ptr(1),
		// 本环节需要操作人的名字
		ApproverName: common.StringPtr(name),
		// 本环节需要操作人的手机号
		ApproverMobile: common.StringPtr(mobile),
		// sms--短信，none--不通知
		NotifyType: common.StringPtr("sms"),
	}
	return approver
}

// 打包企业签署方参与者信息
func BuildOrganizationFlowCreateApprover(name, mobile, organizationName string) *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		ApproverType: common.Int64Ptr(0),
		// 本环节需要操作人的名字
		ApproverName: common.StringPtr(name),
		// 本环节需要操作人的手机号
		ApproverMobile: common.StringPtr(mobile),
		// 本环节需要企业操作人的企业名称
		OrganizationName: common.StringPtr(organizationName),
		// sms--短信，none--不通知
		NotifyType: common.StringPtr("none"),
	}
	return approver
}

// 打包企业静默签署方参与者信息
func BuildServerSignFlowCreateApprover() *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		ApproverType: common.Int64Ptr(3),
	}
	return approver
}
