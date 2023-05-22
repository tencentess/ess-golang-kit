package group

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// GroupCreateFlowByFiles 此接口（CreateFlowByFiles）用来通过上传后的pdf资源编号来创建待签署的合同流程。
//
// 官网文档：https://cloud.tencent.com/document/api/1323/70360
//
// 适用场景1：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
// 适用场景2：可通过该接口传入制式合同文件，同时在指定位置添加签署控件。可以起到接口创建临时模板的效果。如果是标准的制式文件，建议使用模板功能生成模板ID进行合同流程的生成。
// 注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
func GroupCreateFlowByFiles(userId, flowName, fileId, proxyOrganizationId string,
	approvers []*ess.ApproverInfo) (*ess.CreateFlowByFilesResponse, error) {
	client := client_service.GetClientInstance(essGolangKit.SecretId, essGolangKit.SecretKey, essGolangKit.EndPoint)

	request := ess.NewCreateFlowByFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
	}
	// 设置集团子企业账号
	request.Agent = &ess.Agent{
		ProxyOrganizationId: common.StringPtr(proxyOrganizationId),
	}
	request.FlowName = common.StringPtr(flowName)
	request.Approvers = approvers
	request.FileIds = []*string{common.StringPtr(fileId)}

	response, err := client.CreateFlowByFiles(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
