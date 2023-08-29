package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateDocument(userId, flowId, templateId, fileName string) (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.FileNames = []*string{common.StringPtr(fileName)}

	request.FlowId = common.StringPtr(flowId)

	request.TemplateId = common.StringPtr(templateId)

	return client.CreateDocument(request)
}

// CreateDocumentExtended CreateDocument接口的详细参数使用样例，前面简要调用的场景不同，此版本旨在提供可选参数的填入参考。
// 如果您在实现基础场景外有进一步的功能实现需求，可以参考此处代码。
// 注意事项：此处填入参数仅为样例，请在使用时更换为实际值。
func CreateDocumentExtended() (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(essGolangKit.OperatorUserId),
	}

	request.FlowId = common.StringPtr("**********************")

	request.FileNames = []*string{common.StringPtr("filename")}

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

	request.NeedPreview = common.BoolPtr(true)

	request.PreviewType = common.Int64Ptr(1)

	return client.CreateDocument(request)
}
