package group

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestGroupCreateSchemeUrl(t *testing.T) {
	// 成功发起合同的flowId
	flowId := "********************************"
	// 需要代发起的子企业的企业id
	proxyOrganizationId := "********************************"

	response, err := GroupCreateSchemeUrl(essGolangKit.OperatorUserId, flowId, proxyOrganizationId)
	if err != nil {
		t.Errorf("CreateSchemeUrl error: %v", err)
	}
	t.Logf("CreateSchemeUrl finish: %v", response.ToJsonString())
}
