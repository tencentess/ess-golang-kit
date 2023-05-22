package group

import (
	essGolangKit "SdkTools"
	"testing"
)

/**
 * 主企业代子企业查询模板信息样例
 * 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
 */
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
