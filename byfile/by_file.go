package byfile

import (
	essTools "SdkTools"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// BuildApprovers 构造签署人 - 以B2B2C为例, 实际请根据自己的场景构造签署方、控件
func BuildApprovers() []*ess.ApproverInfo {

	// 个人签署方构造参数
	personName := "********************"
	personMobile := "********************"

	// 企业签署方构造参数
	organizationName := "********************"
	organizationUserName := "********************"
	organizationUserMobile := "********************"

	approvers := []*ess.ApproverInfo{
		BuildServerSignApprover(),
		BuildOrganizationApprover(organizationUserName, organizationUserMobile, organizationName),
		BuildPersonApprover(personName, personMobile),
	}

	return approvers
}

// BuildPersonApprover 打包个人签署方参与者信息
// 可选参数传入请参考：https://cloud.tencent.com/document/api/1323/70369#ApproverInfo
func BuildPersonApprover(name, mobile string) *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{

		ApproverType: common.Int64Ptr(1),

		ApproverName: common.StringPtr(name),

		ApproverMobile: common.StringPtr(mobile),

		NotifyType: common.StringPtr("sms"),

		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(146.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),

				FileIndex: common.Int64Ptr(0),

				ComponentPage: common.Int64Ptr(1),

				ComponentType: common.StringPtr("SIGN_SIGNATURE"),
			},
		},
	}
	return approver
}

// BuildOrganizationApprover 打包企业签署方参与者信息
func BuildOrganizationApprover(name, mobile, organizationName string) *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{

		ApproverType: common.Int64Ptr(0),

		ApproverName: common.StringPtr(name),

		ApproverMobile: common.StringPtr(mobile),

		OrganizationName: common.StringPtr(organizationName),

		NotifyType: common.StringPtr("none"),

		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(246.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),

				FileIndex: common.Int64Ptr(0),

				ComponentPage: common.Int64Ptr(1),

				ComponentType: common.StringPtr("SIGN_SEAL"),
			},
		},
	}
	return approver
}

// BuildServerSignApprover 打包企业静默签署方参与者信息
func BuildServerSignApprover() *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{

		ApproverType: common.Int64Ptr(3),

		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(346.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),

				FileIndex: common.Int64Ptr(0),

				ComponentPage: common.Int64Ptr(1),

				ComponentType: common.StringPtr("SIGN_SEAL"),
				
				ComponentValue: common.StringPtr(essTools.ServerSignSealId),
			},
		},
	}
	return approver
}
