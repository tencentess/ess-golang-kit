package fileUploadDownload

import (
	essGolangKit "SdkTools"
	"testing"
)

// 查询转换任务状态调用样例
func TestGetTaskResultApi(t *testing.T) {

	taskId := "***********************"

	response, err := GetTaskResultApi(essGolangKit.OperatorUserId, taskId)
	if err != nil {
		t.Errorf("GetTaskResultApi error: %v", err)
	}
	t.Logf("GetTaskResultApi finish: %v", response.ToJsonString())
}
