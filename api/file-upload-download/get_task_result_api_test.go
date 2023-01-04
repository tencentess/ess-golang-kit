package file_upload_download

import (
	"SdkTools"
	"testing"
)

// 查询转换任务状态调用样例
func TestGetTaskResultApi(t *testing.T) {
	// 任务Id，“创建文件转换任务”接口返回
	taskId := "***********************"

	response, err := GetTaskResultApi(ess_golang_kit.OperatorUserId, taskId)
	if err != nil {
		t.Errorf("GetTaskResultApi error: %v", err)
	}
	t.Logf("GetTaskResultApi finish: %v", response.ToJsonString())
}
