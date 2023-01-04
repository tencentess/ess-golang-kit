package flow_manage

import (
	"SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateDocument 创建电子文档
//
// 官方文档：https://cloud.tencent.com/document/api/1323/70364
//
// 适用场景：见创建签署流程接口。注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
func CreateDocument(userId, flowId, templateId, fileName string) (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 文件名列表,单个文件名最大长度200个字符。目前仅支持单文件发起，此处传入任意自定义值即可
	request.FileNames = []*string{common.StringPtr(fileName)}
	// 签署流程编号,由CreateFlow接口返回
	request.FlowId = common.StringPtr(flowId)
	// 用户上传的模板ID,在控制台模版管理中可以找到
	// 如何创建模板见官网：https://cloud.tencent.com/document/product/1323/61357
	request.TemplateId = common.StringPtr(templateId)

	return client.CreateDocument(request)
}

// CreateDocumentExtended CreateDocument接口的详细参数使用样例，前面简要调用的场景不同，此版本旨在提供可选参数的填入参考。
// 如果您在实现基础场景外有进一步的功能实现需求，可以参考此处代码。
// 注意事项：此处填入参数仅为样例，请在使用时更换为实际值。
func CreateDocumentExtended() (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := client_service.GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(ess_golang_kit.OperatorUserId),
	}

	// 签署流程编号,由CreateFlow接口返回
	// 注意：注意每次创建电子文档前必须先创建流程，文档和流程为一对一的绑定关系！
	request.FlowId = common.StringPtr("**********************")

	// 文件名列表,单个文件名最大长度200个字符。目前仅支持单文件发起，此处传入任意自定义值即可
	request.FileNames = []*string{common.StringPtr("filename")}

	// 用户上传的模板ID,在控制台模版管理中可以找到
	// 如何创建模板见官网：https://cloud.tencent.com/document/product/1323/61357
	request.TemplateId = common.StringPtr("**********************")

	// 这里为需要发起方填写的控件进行赋值操作，推荐使用ComponentName+ComponentValue的方式进行赋值，ComponentName即模板编辑时设置的控件名称
	request.FormFields = []*ess.FormField{
		{
			ComponentName:  common.StringPtr("单行文本"),
			ComponentValue: common.StringPtr("文本内容"),
		},
		{
			ComponentName:  common.StringPtr("勾选框"),
			ComponentValue: common.StringPtr("true"),
		},
	}

	// 是否需要生成预览文件 默认不生成；
	//预览链接有效期300秒；
	// 注意：此处生成的链接只能访问一次，访问过后即失效！
	request.NeedPreview = common.BoolPtr(true)

	// 预览链接类型 默认:0-文件流, 1- H5链接 注意:此参数在NeedPreview 为true 时有效,
	request.PreviewType = common.Int64Ptr(1)

	// 客户端Token，保持接口幂等性,最大长度64个字符
	// 注意：传入相同的token会返回相同的结果，若无需要请不要进行传值！
	//request.ClientToken = common.StringPtr("*********token*******")

	return client.CreateDocument(request)
}
