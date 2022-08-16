package api

import (
	"SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

func UploadFiles(userId, fileBase64, fileName string) (*ess.UploadFilesResponse, error) {
	client := GetClientInstance(ess_golang_kit.SecretId, ess_golang_kit.SecretKey, ess_golang_kit.FileEndPoint)
	uploadFile := &ess.UploadFile{
		FileBody: common.StringPtr(fileBase64),
		FileName: common.StringPtr(fileName),
	}

	request := ess.NewUploadFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")
	request.BusinessType = common.StringPtr("DOCUMENT")
	request.Caller = &ess.Caller{OperatorId: common.StringPtr(userId)}
	request.FileInfos = []*ess.UploadFile{uploadFile}

	return client.UploadFiles(request)
}
