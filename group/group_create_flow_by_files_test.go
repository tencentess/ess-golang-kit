package group

import (
	ess_golang_kit "SdkTools"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

// 主企业代子企业使用文件发起合同的简单样例，如需更详细的参数使用说明，请参考 flow-manage/create_flow_by_files
// 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
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
