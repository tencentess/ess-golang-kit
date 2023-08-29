package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateBatchCancelFlowUrl(userId string, flowIds []string) (*ess.CreateBatchCancelFlowUrlResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateBatchCancelFlowUrlRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}

	request.FlowIds = common.StringPtrs(flowIds)

	return client.CreateBatchCancelFlowUrl(request)
}
