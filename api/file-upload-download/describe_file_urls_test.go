package file_upload_download

import (
	"SdkTools"
	"testing"
)

// 查询文件下载URL调用样例
func TestDescribeFileUrls(t *testing.T) {
	// 业务编号的数组，如模板编号、文档编号、印章编号
	// 最大支持20个资源
	flowId := "********************************"

	response, err := DescribeFileUrls(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		t.Errorf("DescribeFileUrls error: %v", err)
	}
	t.Logf("DescribeFileUrls finish: %v", response.ToJsonString())
}
