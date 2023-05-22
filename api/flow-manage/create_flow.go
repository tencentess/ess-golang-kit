package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlow 创建签署流程
//
// 官网文档：https://cloud.tencent.com/document/api/1323/70361
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
// 注：该接口是通过模板生成合同流程的前置接口，先创建一个不包含签署文件的流程。配合“创建电子文档”接口和“发起流程”接口使用。
func CreateFlow(userId, flowName string, approvers []*ess.FlowCreateApprover) (*ess.CreateFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程参与者信息
	request.Approvers = approvers
	// 签署流程名称,最大长度200个字符
	request.FlowName = common.StringPtr(flowName)

	return client.CreateFlow(request)
}

// CreateFlowExtended CreateFlow接口的详细参数使用样例，前面简要调用的场景不同，此版本旨在提供可选参数的填入参考。
// 如果您在实现基础场景外有进一步的功能实现需求，可以参考此处代码。
// 注意事项：此处填入参数仅为样例，请在使用时更换为实际值。
func CreateFlowExtended() (*ess.CreateFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(essGolangKit.OperatorUserId),
	}

	// 签署流程名称,最大长度200个字符
	request.FlowName = common.StringPtr("测试合同")

	// 构建签署方信息，注意这里的签署人类型、数量、顺序需要和模板中的设置保持一致
	request.Approvers = []*ess.FlowCreateApprover{
		{
			// 企业静默签署
			// 这里我们设置签署方类型为企业方静默签署3，注意当类型为静默签署时，签署人会默认设置为发起方经办人。此时姓名手机号企业名等信息无需填写，且填写无效
			ApproverType: common.Int64Ptr(3),

			// 合同发起后是否短信通知签署方进行签署：sms--短信，none--不通知
			// 企业静默签署不会发送短信通知，此处设置无效
			// NotifyType: common.StringPtr("sms"),
		},
		{
			// 个人签署
			ApproverType: common.Int64Ptr(1),

			// 个人身份签署一般设置姓名+手机号，请确保实际签署时使用的信息和此处一致
			// 本环节需要操作人的名字
			ApproverName: common.StringPtr("张三"),

			// 本环节需要操作人的手机号
			ApproverMobile: common.StringPtr("15912345678"),

			// 合同发起后是否短信通知签署方进行签署：sms--短信，none--不通知
			NotifyType: common.StringPtr("sms"),
		},
		{
			// 企业签署
			ApproverType: common.Int64Ptr(0),

			// 本环节需要操作人的名字
			ApproverName: common.StringPtr("李四"),

			// 本环节需要操作人的手机号
			ApproverMobile: common.StringPtr("15987654321"),

			// 本环节需要企业操作人的企业名称，请注意此处的企业名称要是真实有效的，企业需要在电子签平台进行注册且签署人有加入该企业方能签署。
			// 注：如果该企业尚未注册，或者签署人尚未加入企业，合同仍然可以发起
			OrganizationName: common.StringPtr("XXXXX公司"),

			// 合同发起后是否短信通知签署方进行签署：sms--短信，none--不通知
			NotifyType: common.StringPtr("sms"),
		},
	}

	// 客户端Token，保持接口幂等性,最大长度64个字符
	// 注意：传入相同的token会返回相同的结果，若无需要请不要进行传值！
	//request.ClientToken = common.StringPtr("*********token*******")

	// 发送类型：
	//true：无序签
	//false：有序签
	//注：默认为false（有序签），请和模板中的配置保持一致。如果传值不一致会以模板中设置的为准！
	request.Unordered = common.BoolPtr(false)

	// 用户自定义字段，回调的时候会进行透传，长度需要小于20480
	request.UserData = common.StringPtr("UserData")

	// 签署流程的签署截止时间。
	//值为unix时间戳,精确到秒,不传默认为当前时间一年后
	request.DeadLine = common.Int64Ptr(time.Now().Add(7 * 24 * time.Hour).Unix())

	return client.CreateFlow(request)
}
