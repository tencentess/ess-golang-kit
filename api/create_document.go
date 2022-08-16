package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 创建电子文档
// 适用场景：见创建签署流程接口。注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
func CreateDocument(userId, flowId, templateId, fileName string) (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 文件名列表,单个文件名最大长度200个字符
	request.FileNames = []*string{common.StringPtr(fileName)}
	// 签署流程编号,由CreateFlow接口返回
	request.FlowId = common.StringPtr(flowId)
	// 用户上传的模板ID,在控制台模版管理中可以找到
	// 单个个人签署模版
	request.TemplateId = common.StringPtr(templateId)

	return client.CreateDocument(request)
}
