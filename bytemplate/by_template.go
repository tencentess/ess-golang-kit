package bytemplate

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// BuildFlowCreateApprovers 构造签署人 - 以B2B2C为例, 实际请根据自己的场景构造签署方
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

// BuildPersonFlowCreateApprover 打包个人签署方参与者信息
func BuildPersonFlowCreateApprover(name, mobile string) *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{

		ApproverType: common.Int64Ptr(1),

		ApproverName: common.StringPtr(name),

		ApproverMobile: common.StringPtr(mobile),

		NotifyType: common.StringPtr("sms"),
	}
	return approver
}

// BuildOrganizationFlowCreateApprover 打包企业签署方参与者信息
func BuildOrganizationFlowCreateApprover(name, mobile, organizationName string) *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{

		ApproverType: common.Int64Ptr(0),

		ApproverName: common.StringPtr(name),

		ApproverMobile: common.StringPtr(mobile),

		OrganizationName: common.StringPtr(organizationName),

		NotifyType: common.StringPtr("none"),
	}
	return approver
}

// BuildServerSignFlowCreateApprover 打包企业静默签署方参与者信息
func BuildServerSignFlowCreateApprover() *ess.FlowCreateApprover {
	// 签署参与者信息
	approver := &ess.FlowCreateApprover{

		ApproverType: common.Int64Ptr(3),
	}
	return approver
}
