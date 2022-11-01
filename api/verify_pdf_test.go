package api

import (
	"SdkTools"
	"testing"
)

// 合同文件验签调用样例
func TestVerifyPdf(t *testing.T) {
	flowId := ""

	response, err := VerifyPdf(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("VerifyPdf error: %v", err)
	}
	t.Logf("VerifyPdf finish: %v", response.ToJsonString())
}
