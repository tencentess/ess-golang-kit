package documentManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestCreateMultiFlowSignQRCode(t *testing.T) {
	// 模板ID
	templateId := "********************************"
	// 签署流程名称，最大长度不超过200字符
	flowName := "扫码签流程"

	response, err := CreateMultiFlowSignQRCode(essGolangKit.OperatorUserId, templateId, flowName)
	if err != nil {
		t.Errorf("CreateMultiFlowSignQRCode error: %v", err)
	}
	t.Logf("CreateMultiFlowSignQRCode finish: %v", response.ToJsonString())
}
