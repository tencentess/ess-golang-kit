package group

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
	"time"
)

func TestGroupCreateFlowByTemplate(t *testing.T) {

	// 需要代发起的子企业的企业id
	proxyOrganizationId := "*****************"

	templateId := "**************"

	flowName := "****************"

	name := "*****************"

	mobile := "*****************"

	fileName := "****************"

	approvers := []*ess.FlowCreateApprover{
		{
			ApproverType:   common.Int64Ptr(1),
			ApproverName:   common.StringPtr(name),
			ApproverMobile: common.StringPtr(mobile),
		},
	}

	// 1、调用CreateFlow接口创建流程
	createFlowResponse, createFlowErr := GroupCreateFlow(ess_golang_kit.OperatorUserId, flowName, approvers, proxyOrganizationId)
	if createFlowErr != nil {
		t.Errorf("GroupCreateFlow error: %v", createFlowErr)
		return
	}
	t.Logf("GroupCreateFlow finish: %v", createFlowResponse.ToJsonString())

	flowId := *createFlowResponse.Response.FlowId

	// 2、调用CreateDocument接口创建文档
	createDocumentResponse, createDocumentErr := GroupCreateDocument(ess_golang_kit.OperatorUserId, flowId, templateId, fileName, proxyOrganizationId)
	if createDocumentErr != nil {
		t.Errorf("GroupCreateDocument error: %v", createDocumentErr)
		return
	}
	t.Logf("GroupCreateDocument finish: %v", createDocumentResponse.ToJsonString())

	time.Sleep(3 * time.Second)

	// 3、调用StartFlow接口开启流程
	startFlowResponse, startFlowErr := GroupStartFlow(ess_golang_kit.OperatorUserId, flowId, proxyOrganizationId)
	if startFlowErr != nil {
		t.Errorf("GroupStartFlow error: %v", startFlowErr)
		return
	}
	t.Logf("GroupStartFlow finish: %v", startFlowResponse.ToJsonString())

}
