package documentManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 合同文件验签调用样例
func TestVerifyPdf(t *testing.T) {
	flowId := ""

	response, err := VerifyPdf(essGolangKit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("VerifyPdf error: %v", err)
	}
	t.Logf("VerifyPdf finish: %v", response.ToJsonString())
}
