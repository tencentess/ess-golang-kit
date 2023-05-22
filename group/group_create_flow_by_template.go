package group

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// GroupCreateFlow 创建签署流程(CreateFlow)
//
// 官网文档：https://cloud.tencent.com/document/api/1323/70361
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
// 注：该接口是通过模板生成合同流程的前置接口，先创建一个不包含签署文件的流程。配合“创建电子文档”接口和“发起流程”接口使用。
func GroupCreateFlow(userId, flowName string, approvers []*ess.FlowCreateApprover, proxyOrganizationId string) (*ess.CreateFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}
	// 签署流程参与者信息
	request.Approvers = approvers
	// 签署流程名称,最大长度200个字符
	request.FlowName = common.StringPtr(flowName)

	return client.CreateFlow(request)
}

// GroupCreateDocument 创建电子文档(CreateDocument)
//
// 官方文档：https://cloud.tencent.com/document/api/1323/70364
//
// 适用场景：见创建签署流程接口。注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
func GroupCreateDocument(userId, flowId, templateId, fileName, proxyOrganizationId string) (*ess.CreateDocumentResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateDocumentRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
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

// GroupStartFlow 此接口用于发起流程(StartFlow)
//
// 官网文档：https://cloud.tencent.com/document/product/1323/70357
//
// 适用场景：见创建签署流程接口。
// 注：该接口是“创建电子文档”接口的后置接口，用于激活包含完整合同信息（模板及内容信息）的流程。激活后的流程就是一份待签署的电子合同。
func GroupStartFlow(userId, flowId, proxyOrganizationId string) (*ess.StartFlowResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewStartFlowRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}
	// 签署流程编号，由CreateFlow接口返回
	request.FlowId = common.StringPtr(flowId)

	return client.StartFlow(request)
}
