package api

import (
	ess_golang_kit "SdkTools"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

func TestCreateFlowByFiles(t *testing.T) {
	flowName := ""
	fileId := ""
	var approvers []*ess.ApproverInfo
	response, err := CreateFlowByFiles(ess_golang_kit.OperatorUserId, flowName, fileId, approvers)
	if err != nil {
		t.Errorf("CreateFlowByFiles error: %v", err)
	}
	t.Logf("UploadFiles finish: %v", response.ToJsonString())
}
