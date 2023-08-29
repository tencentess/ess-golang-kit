package group

import (
	essGolangKit "SdkTools"
	client_service "SdkTools/api/client-service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

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
