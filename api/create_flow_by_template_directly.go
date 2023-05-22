package api

import (
	essGolangKit "SdkTools"
	flowManage "SdkTools/api/flow-manage"
	"time"

	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlowByTemplateDirectly 通过模板发起签署流程，并查询签署链接
func CreateFlowByTemplateDirectly(userId, flowName string, approvers []*ess.FlowCreateApprover) (flowId,
	schemeUrl string, err error) {

	// 创建流程
	createFlowResp, err := flowManage.CreateFlow(userId, flowName, approvers)
	if err != nil {
		return "", "", err
	}
	flowId = *createFlowResp.Response.FlowId

	// 创建电子文档，注意每次创建电子文档前必须先创建流程，文档和流程为一对一的绑定关系
	_, err = flowManage.CreateDocument(userId, flowId, essGolangKit.TemplateId, flowName)
	if err != nil {
		return "", "", err
	}

	// 注：文档合成操作是异步完成的，所以此处需要进行适当等待，或者在StartFlow接口报错文档不可用时进行重试（使用相同参数）
	time.Sleep(time.Duration(3) * time.Second)

	// 开启流程
	_, err = flowManage.StartFlow(userId, flowId)
	if err != nil {
		return "", "", err
	}

	// 获取签署链接
	schemeResp, err := flowManage.CreateSchemeUrl(userId, flowId)
	if err != nil {
		return "", "", err
	}
	schemeUrl = *schemeResp.Response.SchemeUrl

	return flowId, schemeUrl, nil
}
