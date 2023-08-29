package group

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestGroupCreateSchemeUrl(t *testing.T) {

	flowId := "********************************"

	proxyOrganizationId := "********************************"

	response, err := GroupCreateSchemeUrl(essGolangKit.OperatorUserId, flowId, proxyOrganizationId)
	if err != nil {
		t.Errorf("CreateSchemeUrl error: %v", err)
	}
	t.Logf("CreateSchemeUrl finish: %v", response.ToJsonString())
}
