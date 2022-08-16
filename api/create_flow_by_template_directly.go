package api

import (
	"SdkTools"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"time"
)

// 通过模板发起签署流程，并查询签署链接
func CreateFlowByTemplateDirectly(userId, flowName string, approvers []*ess.FlowCreateApprover) (flowId,
	schemeUrl string, err error) {

	// 创建流程
	createFlowResp, err := CreateFlow(userId, flowName, approvers)
	if err != nil {
		return "", "", err
	}
	flowId = *createFlowResp.Response.FlowId

	// 创建电子文档
	_, err = CreateDocument(userId, flowId, ess_golang_kit.TemplateId, flowName)
	if err != nil {
		return "", "", err
	}

	// 等待文档异步合成
	time.Sleep(time.Duration(3) * time.Second)

	// 开启流程
	_, err = StartFlow(userId, flowId)
	if err != nil {
		return "", "", err
	}

	// 获取签署链接
	schemeResp, err := CreateSchemeUrl(userId, flowId)
	if err != nil {
		return "", "", err
	}
	schemeUrl = *schemeResp.Response.SchemeUrl

	return flowId, schemeUrl, nil
}
