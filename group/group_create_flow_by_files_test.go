package group

import (
	ess_golang_kit "SdkTools"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

func TestGroupCreateFlowByFiles(t *testing.T) {
	flowName := ""
	fileId := ""
	// 需要代发起的子企业的企业id
	proxyOrganizationId := ""

	var approvers []*ess.ApproverInfo
	response, err := GroupCreateFlowByFiles(ess_golang_kit.OperatorUserId, flowName, fileId, proxyOrganizationId, approvers)
	if err != nil {
		t.Errorf("CreateFlowByFiles error: %v", err)
	}
	t.Logf("CreateFlowByFiles finish: %v", response.ToJsonString())
}
