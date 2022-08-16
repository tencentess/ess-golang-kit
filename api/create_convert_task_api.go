package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 此接口用于创建文件转换任务
// 适用场景：将doc/docx文件转化为pdf文件
func CreateConvertTaskApi(userId, resourceId, resourceType, resourceName string) (*ess.CreateConvertTaskApiResponse, error) {
	// 构造客户端调用实例
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateConvertTaskApiRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 资源Id，由UploadFiles接口返回
	request.ResourceId = common.StringPtr(resourceId)
	// 资源类型，2-doc 3-docx
	request.ResourceType = common.StringPtr(resourceType)
	// 资源名称
	request.ResourceName = common.StringPtr(resourceName)

	return client.CreateConvertTaskApi(request)
}
