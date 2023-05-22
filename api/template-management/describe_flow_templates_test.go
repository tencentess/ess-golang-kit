package templateManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 查询模板调用样例
func TestDescribeFlowTemplates(t *testing.T) {
	templateId := "********************************"

	response, err := DescribeFlowTemplates(essGolangKit.OperatorUserId, templateId)
	if err != nil {
		t.Errorf("DescribeFlowTemplates error: %v", err)
	}
	t.Logf("DescribeFlowTemplates finish: %v", response.ToJsonString())
}
