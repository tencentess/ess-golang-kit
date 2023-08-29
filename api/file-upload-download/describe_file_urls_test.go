package fileUploadDownload

import (
	essGolangKit "SdkTools"
	"testing"
)

// 查询文件下载URL调用样例
func TestDescribeFileUrls(t *testing.T) {

	flowId := "********************************"

	response, err := DescribeFileUrls(essGolangKit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("DescribeFileUrls error: %v", err)
	}
	t.Logf("DescribeFileUrls finish: %v", response.ToJsonString())
}
