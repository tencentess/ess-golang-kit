package fileUploadDownload

import (
	essGolangKit "SdkTools"
	"testing"
)

func TestUploadFiles(t *testing.T) {
	fileName := ""
	fileBase64 := ""
	response, err := UploadFiles(essGolangKit.OperatorUserId, fileBase64, fileName)
	if err != nil {
		t.Errorf("UploadFiles error: %v", err)
	}
	t.Logf("UploadFiles finish: %v", response.ToJsonString())
}
