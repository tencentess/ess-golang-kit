package fileUploadDownload

import (
	essGolangKit "SdkTools"
	"testing"
)

// 创建文件转换任务调用样例
func TestCreateConvertTaskApi(t *testing.T) {

	resourceId := "********************************"

	resourceType := "********************************"

	resourceName := "********************************"

	response, err := CreateConvertTaskApi(essGolangKit.OperatorUserId, resourceId, resourceType, resourceName)
	if err != nil {
		t.Errorf("CreateConvertTaskApi error: %v", err)
	}
	t.Logf("CreateConvertTaskApi finish: %v", response.ToJsonString())
}
