package flowManage

import (
	essGolangKit "SdkTools"
	"testing"
)

// 创建电子文档调用样例
func TestCreateDocument(t *testing.T) {

	flowId := "********************************"

	templateId := "********************************"

	fileName := "********************************"

	response, err := CreateDocument(essGolangKit.OperatorUserId, flowId, templateId, fileName)
	if err != nil {
		t.Errorf("CreateDocument error: %v", err)
	}
	t.Logf("CreateDocument finish: %v", response.ToJsonString())
}
