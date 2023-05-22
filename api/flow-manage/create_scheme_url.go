package flowManage

import (
	essGolangKit "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateSchemeUrl 获取小程序跳转链接
//
// 官网文档： https://cloud.tencent.com/document/product/1323/70359
//
// 适用场景：如果需要签署人在自己的APP、小程序、H5应用中签署，可以通过此接口获取跳转腾讯电子签小程序的签署跳转链接。
// 注：如果签署人是在PC端扫码签署，可以通过生成跳转链接自主转换成二维码，让签署人在PC端扫码签署。
// 跳转到小程序的实现，参考官方文档（分为全屏、半屏两种方式）
// 如您需要自主配置小程序跳转链接，请参考: 跳转小程序链接配置说明
func CreateSchemeUrl(userId, flowId string) (*ess.CreateSchemeUrlResponse, error) {
	// 构造客户端调用实例
	client := clientService.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateSchemeUrlRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 签署流程编号 (PathType=1时必传)
	request.FlowId = common.StringPtr(flowId)
	// 跳转页面 1: 小程序合同详情 2: 小程序合同列表页 0: 不传, 默认主页
	request.PathType = common.Uint64Ptr(1)

	return client.CreateSchemeUrl(request)
}
