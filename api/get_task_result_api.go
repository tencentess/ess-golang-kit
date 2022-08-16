package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 此接口用于查询转换任务状态
// 适用场景：将doc/docx文件转化为pdf文件
// 注：该接口是“创建文件转换任务”接口的后置接口，用于查询转换任务的执行结果
func GetTaskResultApi(userId, taskId string) (*ess.GetTaskResultApiResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewGetTaskResultApiRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 任务Id，“创建文件转换任务”接口返回
	request.TaskId = common.StringPtr(taskId)

	return client.GetTaskResultApi(request)
}
