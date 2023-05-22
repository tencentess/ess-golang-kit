package api

import (
	fileUploadDownload "SdkTools/api/file-upload-download"
	flowManage "SdkTools/api/flow-manage"

	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// CreateFlowByFileDirectly 通过文件base64直接发起签署流程，返回flowid
func CreateFlowByFileDirectly(userId, fileBase64, flowName string, approvers []*ess.ApproverInfo) (flowId,
	schemeUrl string, err error) {

	// 上传文件获取fileId
	uploadResp, err := fileUploadDownload.UploadFiles(userId, fileBase64, flowName)
	if err != nil {
		return "", "", err
	}
	fileId := uploadResp.Response.FileIds[0]

	// 使用文件创建签署流程
	createFlowResp, err := flowManage.CreateFlowByFiles(userId, flowName, *fileId, approvers)
	if err != nil {
		return "", "", err
	}
	flowId = *createFlowResp.Response.FlowId

	// 获取签署链接
	schemeResp, err := flowManage.CreateSchemeUrl(userId, flowId)
	if err != nil {
		return "", "", err
	}
	schemeUrl = *schemeResp.Response.SchemeUrl

	return flowId, schemeUrl, nil
}
