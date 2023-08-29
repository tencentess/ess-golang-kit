package group

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestGroupDescribeFlowTemplates(t *testing.T) {
	templateId := "********************************"
	// 需要代发起的子企业的企业id
	proxyOrganizationId := "********************************"
	response, err := GroupDescribeFlowTemplates(essGolangKit.OperatorUserId, templateId, proxyOrganizationId)
	if err != nil {
		t.Errorf("DescribeFlowTemplates error: %v", err)
	}
	t.Logf("DescribeFlowTemplates finish: %v", response.ToJsonString())
}
