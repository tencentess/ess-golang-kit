package flowManage

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateFlowByFiles(userId, flowName, fileId string,
	approvers []*ess.ApproverInfo) (*ess.CreateFlowByFilesResponse, error) {
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowByFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	request.FlowName = common.StringPtr(flowName)
	request.Approvers = approvers
	request.FileIds = []*string{common.StringPtr(fileId)}

	response, err := client.CreateFlowByFiles(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateFlowByFilesExtended CreateFlowByFiles接口的详细参数使用样例，前面简要调用的场景不同，此版本旨在提供可选参数的填入参考。
// 如果您在实现基础场景外有进一步的功能实现需求，可以参考此处代码。
// 注意事项：此处填入参数仅为样例，请在使用时更换为实际值。
func CreateFlowByFilesExtended() (*ess.CreateFlowByFilesResponse, error) {
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)
	request := ess.NewCreateFlowByFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(essGolangKit.OperatorUserId),
	}

	request.FlowName = common.StringPtr("测试合同")

	// 构建签署方信息
	// 注意：文件发起时，签署方不能进行控件填写！！！如果有填写需求，请设置为发起方填写，或者使用模板发起！！！
	packFlowApprovers(request)

	request.FileIds = []*string{common.StringPtr("*************************")}

	request.FlowType = common.StringPtr("销售合同")

	// 经办人内容控件配置，必须在此处给控件进行赋值，合同发起时控件即被填写完成。
	request.Components = []*ess.Component{
		// 坐标定位，单行文本类型
		buildComponentNormal("TEXT", "单行文本"),
		// 表单域定位，单行文本类型
		buildComponentField("TEXT", "单行文本"),
		// 关键字定位，单行文本类型
		buildComponentKeyword("TEXT", "单行文本"),
	}

	request.NeedPreview = common.BoolPtr(false)

	request.PreviewType = common.Int64Ptr(0)

	request.Deadline = common.Int64Ptr(time.Now().Add(7 * 24 * time.Hour).Unix())

	request.Unordered = common.BoolPtr(false)

	request.UserData = common.StringPtr("UserData")

	response, err := client.CreateFlowByFiles(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// packFlowApprovers 构建签署方信息
func packFlowApprovers(request *ess.CreateFlowByFilesRequest) {
	request.Approvers = []*ess.ApproverInfo{
		{
			// 企业静默签署
			ApproverType: common.Int64Ptr(3),

			// 这里可以设置企业方自动签章，分别可以使用坐标、表单域、关键字进行定位
			SignComponents: []*ess.Component{
				// 坐标定位，印章类型，并传入印章Id
				buildComponentNormal("SIGN_SEAL", "*************"),
				// 表单域定位，印章类型，并传入印章Id
				buildComponentField("SIGN_SEAL", "*************"),
				// 关键字定位，印章类型，并传入印章Id
				buildComponentKeyword("SIGN_SEAL", "*************"),
			},
		},
		{
			// 个人签署
			ApproverType: common.Int64Ptr(1),

			ApproverName:   common.StringPtr("张三"),
			ApproverMobile: common.StringPtr("1*********8"),

			// 这里可以设置用户进行手动签名，分别可以使用坐标、表单域、关键字进行定位
			SignComponents: []*ess.Component{
				// 坐标定位，手写签名类型
				buildComponentNormal("SIGN_SIGNATURE", ""),
				// 表单域定位，手写签名类型
				buildComponentField("SIGN_SIGNATURE", ""),
				// 关键字定位，手写签名类型
				buildComponentKeyword("SIGN_SIGNATURE", ""),
			},
		},
		{
			// 企业签署
			ApproverType: common.Int64Ptr(0),

			ApproverName:     common.StringPtr("李四"),
			ApproverMobile:   common.StringPtr("1*********1"),
			OrganizationName: common.StringPtr("XXXXX公司"),

			// 这里可以设置企业手动签章（如果需要自动请使用静默签署），分别可以使用坐标、表单域、关键字进行定位
			SignComponents: []*ess.Component{
				// 坐标定位，印章类型
				buildComponentNormal("SIGN_SEAL", ""),
				// 表单域定位，印章类型
				buildComponentField("SIGN_SEAL", ""),
				// 关键字定位，印章类型
				buildComponentKeyword("SIGN_SEAL", ""),
			},
		},
	}
}

// buildSignComponentNormal 使用坐标模式进行控件定位
func buildComponentNormal(componentType, componentValue string) *ess.Component {
	// 可选参数传入请参考：https://cloud.tencent.com/document/api/1323/70369#Component
	component := &ess.Component{

		ComponentPosX:   common.Float64Ptr(146.15625),
		ComponentPosY:   common.Float64Ptr(472.78125),
		ComponentWidth:  common.Float64Ptr(112),
		ComponentHeight: common.Float64Ptr(40),

		FileIndex: common.Int64Ptr(0),

		ComponentPage: common.Int64Ptr(1),

		ComponentType: common.StringPtr(componentType),
	}

	if len(componentValue) != 0 {
		component.ComponentValue = common.StringPtr(componentValue)
	}
	return component
}

// buildComponentKeyword 使用关键字模式进行控件定位
func buildComponentKeyword(componentType, componentValue string) *ess.Component {
	component := &ess.Component{
		GenerateMode: common.StringPtr("KEYWORD"),

		ComponentId: common.StringPtr("签名"),

		ComponentWidth:  common.Float64Ptr(112),
		ComponentHeight: common.Float64Ptr(40),

		FileIndex: common.Int64Ptr(0),

		OffsetX: common.Float64Ptr(10.5),
		OffsetY: common.Float64Ptr(10.5),

		KeywordOrder: common.StringPtr("Reverse"),

		KeywordIndexes: []*int64{common.Int64Ptr(1)},

		ComponentType: common.StringPtr(componentType),
	}

	if len(componentValue) != 0 {
		component.ComponentValue = common.StringPtr(componentValue)
	}
	return component
}

// buildComponentField 使用表单域模式进行控件定位
func buildComponentField(componentType, componentValue string) *ess.Component {
	component := &ess.Component{

		GenerateMode: common.StringPtr("FIELD"),

		ComponentName: common.StringPtr("表单"),

		FileIndex: common.Int64Ptr(0),

		ComponentType: common.StringPtr(componentType),
	}

	if len(componentValue) != 0 {
		component.ComponentValue = common.StringPtr(componentValue)
	}

	return component
}
