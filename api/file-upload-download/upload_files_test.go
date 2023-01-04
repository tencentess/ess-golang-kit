package file_upload_download

import (
	"SdkTools"
	"testing"
)

func TestUploadFiles(t *testing.T) {
	fileName := ""
	fileBase64 := ""
	response, err := UploadFiles(ess_golang_kit.OperatorUserId, fileBase64, fileName)
	if err != nil {
		t.Errorf("UploadFiles error: %v", err)
	}
	t.Logf("UploadFiles finish: %v", response.ToJsonString())
}
