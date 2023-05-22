package fileUploadDownload

import (
	essGolangKit "SdkTools"
	"testing"
)

// 创建文件转换任务调用样例
func TestCreateConvertTaskApi(t *testing.T) {
	// 资源Id，由UploadFiles接口返回
	resourceId := "********************************"
	// 资源类型，2-doc 3-docx
	resourceType := "********************************"
	// 资源名称
	resourceName := "********************************"

	response, err := CreateConvertTaskApi(essGolangKit.OperatorUserId, resourceId, resourceType, resourceName)
	if err != nil {
		t.Errorf("CreateConvertTaskApi error: %v", err)
	}
	t.Logf("CreateConvertTaskApi finish: %v", response.ToJsonString())
}
