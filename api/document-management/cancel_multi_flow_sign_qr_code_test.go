package documentManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 取消一码多扫二维码调用样例
func TestCancelMultiFlowSignQRCode(t *testing.T) {

	qrCodeId := "********************************"

	response, err := CancelMultiFlowSignQRCode(essGolangKit.OperatorUserId, qrCodeId)
	if err != nil {
		t.Errorf("CancelMultiFlowSignQRCode error: %v", err)
	}
	t.Logf("CancelMultiFlowSignQRCode finish: %v", response.ToJsonString())
}
