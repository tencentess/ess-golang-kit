package flowManage

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestCreateSchemeUrl(t *testing.T) {
	// 成功发起合同的flowId
	flowId := "********************************"

	response, err := CreateSchemeUrl(essGolangKit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("CreateSchemeUrl error: %v", err)
	}
	t.Logf("CreateSchemeUrl finish: %v", response.ToJsonString())
}
