package api

import (
	"SdkTools"
	"testing"
)

// 取消一码多扫二维码调用样例
func TestCancelMultiFlowSignQRCode(t *testing.T) {
	// 二维码id
	qrCodeId := "********************************"

	response, err := CancelMultiFlowSignQRCode(ess_golang_kit.OperatorUserId, qrCodeId)
	if err != nil {
		t.Errorf("CancelMultiFlowSignQRCode error: %v", err)
	}
	t.Logf("CancelMultiFlowSignQRCode finish: %v", response.ToJsonString())
}
