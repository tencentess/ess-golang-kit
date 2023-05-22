package group

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
	"time"
)

/**
 * 主企业代子企业使用模板发起合同的简单样例，如需更详细的参数使用说明，请参考 flow-manage 目录下的 create_flow/_create_document/_start_flow
 * 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
 */
func TestGroupCreateFlowByTemplate(t *testing.T) {

	// 需要代发起的子企业的企业id
	proxyOrganizationId := "*****************"
	// 子企业模板id，需要在控制台查询获取。请勿使用主企业的模板id!
	templateId := "**************"
	// 合同名称
	flowName := "****************"
	// 姓名
	name := "*****************"
	// 手机号
	mobile := "*****************"
	// 文件名
	fileName := "****************"

	// FlowCreateApprover https://cloud.tencent.com/document/api/1323/70369#FlowCreateApprover
	// 参与者类型：
	//0：企业
	//1：个人
	//3：企业静默签署
	//注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。
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
