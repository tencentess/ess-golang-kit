package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func CreateFlowByFiles(userId, flowName, fileId string, approvers []*ess.ApproverInfo) (*ess.CreateFlowByFilesResponse, error) {
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.EndPoint)

	request := ess.NewCreateFlowByFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	// 调用方用户信息，参考通用结构
	request.Operator = &ess.UserInfo{
		UserId: common.StringPtr(userId),
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
