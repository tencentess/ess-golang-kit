package byfile

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 构造签署人 - 以B2B2C为例, 实际请根据自己的场景构造签署方、控件
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

// 打包个人签署方参与者信息
func BuildPersonApprover(name, mobile string) *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{
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
		// 本环节操作人签署控件配置，为企业静默签署时，只允许类型为SIGN_SEAL（印章）和SIGN_DATE（日期）控件，并且传入印章编号
		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(146.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				FileIndex:       common.Int64Ptr(0),
				ComponentPage:   common.Int64Ptr(1),
				ComponentType:   common.StringPtr("SIGN_SIGNATURE"),
			},
		},
	}
	return approver
}

// 打包企业签署方参与者信息
func BuildOrganizationApprover(name, mobile, organizationName string) *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{
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
		// 本环节操作人签署控件配置，为企业静默签署时，只允许类型为SIGN_SEAL（印章）和SIGN_DATE（日期）控件，并且传入印章编号
		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(246.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				FileIndex:       common.Int64Ptr(0),
				ComponentPage:   common.Int64Ptr(1),
				ComponentType:   common.StringPtr("SIGN_SEAL"),
			},
		},
	}
	return approver
}

// 打包企业静默签署方参与者信息
func BuildServerSignApprover() *ess.ApproverInfo {
	// 签署参与者信息
	approver := &ess.ApproverInfo{
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		ApproverType: common.Int64Ptr(3),
		// 本环节操作人签署控件配置，为企业静默签署时，只允许类型为SIGN_SEAL（印章）和SIGN_DATE（日期）控件，并且传入印章编号
		SignComponents: []*ess.Component{
			{
				ComponentPosX:   common.Float64Ptr(346.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				FileIndex:       common.Int64Ptr(0),
				ComponentPage:   common.Int64Ptr(1),
				ComponentType:   common.StringPtr("SIGN_SEAL"),
				ComponentValue:  common.StringPtr(ess_golang_kit.ServerSignSealId),
			},
		},
	}
	return approver
}

// 构建控件信息
func BuildComponent(componentPosX, componentPosY, componentWidth, componentHeight float64,
	fileIndex, componentPage int64, componentType, componentValue string) *ess.Component {
	component := &ess.Component{
		// 参数控件X位置，单位pt
		ComponentPosX: common.Float64Ptr(componentPosX),
		// 参数控件Y位置，单位pt
		ComponentPosY: common.Float64Ptr(componentPosY),
		// 参数控件宽度，单位pt
		ComponentWidth: common.Float64Ptr(componentWidth),
		// 参数控件高度，单位pt
		ComponentHeight: common.Float64Ptr(componentHeight),
		// 控件所属文件的序号（取值为：0-N）
		FileIndex: common.Int64Ptr(fileIndex),
		// 参数控件所在页码，取值为：1-N
		ComponentPage: common.Int64Ptr(componentPage),
		// 如果是 Component 控件类型，则可选类型为：
		// TEXT - 单行文本
		// MULTI_LINE_TEXT - 多行文本
		// CHECK_BOX - 勾选框
		// ATTACHMENT - 附件
		// SELECTOR - 选择器
		// 如果是 SignComponent 控件类型，则可选类型为：
		// SIGN_SEAL - 签署印章控件，静默签署时需要传入印章id作为ComponentValue
		// SIGN_DATE - 签署日期控件
		// SIGN_SIGNATURE - 手写签名控件，静默签署时不能使用
		ComponentType: common.StringPtr(componentType),
		// 自动签署所对应的印章Id
		ComponentValue: common.StringPtr(componentValue),
	}
	return component
}
