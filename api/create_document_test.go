package api

import (
	"SdkTools"
	"testing"
)

// 创建电子文档调用样例
func TestCreateDocument(t *testing.T) {
	// 签署流程编号,由CreateFlow接口返回
	flowId := "********************************"
	// 用户上传的模板ID,在控制台模版管理中可以找到
	// 单个个人签署模版
	templateId := "********************************"
	// 文件名列表,单个文件名最大长度200个字符
	fileName := "********************************"

	response, err := CreateDocument(ess_golang_kit.OperatorUserId, flowId, templateId, fileName)
	if err != nil {
		t.Errorf("CreateDocument error: %v", err)
	}
	t.Logf("CreateDocument finish: %v", response.ToJsonString())
}
