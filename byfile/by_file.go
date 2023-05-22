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
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 这里我们设置为1，即身份为个人
		ApproverType: common.Int64Ptr(1),

		// 本环节需要操作人的名字
		ApproverName: common.StringPtr(name),

		// 本环节需要操作人的手机号
		ApproverMobile: common.StringPtr(mobile),

		// 合同发起后是否短信通知签署方进行签署： sms--短信，none--不通知
		NotifyType: common.StringPtr("sms"),

		// 签署人签署控件配置，数组传入，可以设置多个。此处我们为签署方设置一个手写签名控件。
		// 可选参数传入请参考：https://cloud.tencent.com/document/api/1323/70369#Component
		SignComponents: []*ess.Component{
			{
				// 以下4项确定了控件在pdf文件内的坐标位置以及长宽信息，这里我们给出一些预设值
				// 如何确定坐标请参考： https://doc.weixin.qq.com/doc/w3_AKgAhgboACgsf9NKAVqSOKVIkQ0vQ?scode=AJEAIQdfAAoz9916DRAKgAhgboACg
				ComponentPosX:   common.Float64Ptr(146.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				// 控件所属文件的序号，目前均为单文件发起，所以我们固定填入序号0
				FileIndex: common.Int64Ptr(0),
				// 控件所在的页面数，从1开始取值，请勿超出pdf文件的最大页数
				ComponentPage: common.Int64Ptr(1),
				// 控件类型，这里选择用户手写签名SIGN_SIGNATURE，阅读传参文档时请注意此处为SignComponent控件类型
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
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		// 这里我们设置为企业方手动签署0
		ApproverType: common.Int64Ptr(0),

		// 本环节需要操作人的名字
		ApproverName: common.StringPtr(name),

		// 本环节需要操作人的手机号
		ApproverMobile: common.StringPtr(mobile),

		// 本环节需要企业操作人的企业名称，请注意此处的企业名称要是真实有效的，企业需要在电子签平台进行注册且签署人有加入该企业方能签署。
		// 注：如果该企业尚未注册，或者签署人尚未加入企业，合同仍然可以发起
		OrganizationName: common.StringPtr(organizationName),

		// 合同发起后是否短信通知签署方进行签署：sms--短信，none--不通知
		NotifyType: common.StringPtr("none"),

		// 签署人签署控件配置，数组传入，可以设置多个。此处我们为签署方设置一个印章控件。
		// 可选参数传入请参考：https://cloud.tencent.com/document/api/1323/70369#Component
		SignComponents: []*ess.Component{
			{
				// 以下4项确定了控件在pdf文件内的坐标位置以及长宽信息，这里我们给出一些预设值
				// 如何确定坐标请参考： https://doc.weixin.qq.com/doc/w3_AKgAhgboACgsf9NKAVqSOKVIkQ0vQ?scode=AJEAIQdfAAoz9916DRAKgAhgboACg
				ComponentPosX:   common.Float64Ptr(246.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				// 控件所属文件的序号，目前均为单文件发起，所以我们固定填入序号0
				FileIndex: common.Int64Ptr(0),
				// 控件所在的页面数，从1开始取值，请勿超出pdf文件的最大页数
				ComponentPage: common.Int64Ptr(1),
				// 控件类型，这里选择印章控件SIGN_SEAL，阅读传参文档时请注意此处为SignComponent控件类型
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
		// 参与者类型：
		// 0：企业
		// 1：个人
		// 3：企业静默签署
		// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
		// 这里我们设置签署方类型为企业方静默签署3，注意当类型为静默签署时，签署人会默认设置为发起方经办人
		ApproverType: common.Int64Ptr(3),

		// 签署人签署控件配置，数组传入，可以设置多个。签署方类型为企业静默签署时，只允许类型为SIGN_SEAL（印章）和SIGN_DATE（日期）控件，并且传入印章编号
		// 可选参数传入请参考：https://cloud.tencent.com/document/api/1323/70369#Component
		SignComponents: []*ess.Component{
			{
				// 以下4项确定了控件在pdf文件内的坐标位置以及长宽信息，这里我们给出一些预设值
				// 如何确定坐标请参考： https://doc.weixin.qq.com/doc/w3_AKgAhgboACgsf9NKAVqSOKVIkQ0vQ?scode=AJEAIQdfAAoz9916DRAKgAhgboACg
				ComponentPosX:   common.Float64Ptr(346.15625),
				ComponentPosY:   common.Float64Ptr(472.78125),
				ComponentWidth:  common.Float64Ptr(112),
				ComponentHeight: common.Float64Ptr(40),
				// 控件所属文件的序号，目前均为单文件发起，所以我们固定填入序号0
				FileIndex: common.Int64Ptr(0),
				// 控件所在的页面数，从1开始取值，请勿超出pdf文件的最大页数
				ComponentPage: common.Int64Ptr(1),
				// 控件类型，这里选择印章控件SIGN_SEAL，阅读传参文档时请注意此处为SignComponent控件类型
				ComponentType: common.StringPtr("SIGN_SEAL"),
				// 印章Id，发起后会使用该印章在指定区域进行自动签章
				ComponentValue: common.StringPtr(essTools.ServerSignSealId),
			},
		},
	}
	return approver
}
