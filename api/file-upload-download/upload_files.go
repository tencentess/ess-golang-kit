package fileUploadDownload

import (
	essTools "SdkTools"
	clientService "SdkTools/api/client-service"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// UploadFiles 此接口（UploadFiles）用于文件上传。
//
// 官网文档：https://cloud.tencent.com/document/api/1323/73066
//
// 适用场景：用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
// 调用是请注意此处的 Endpoint 和其他接口不同
func UploadFiles(userId, fileBase64, fileName string) (*ess.UploadFilesResponse, error) {
	client := clientService.GetClientInstance(essTools.SecretId, essTools.SecretKey, essTools.FileEndPoint)
	// 上传文件内容数组，最多支持20个文件
	uploadFile := &ess.UploadFile{
		FileBody: common.StringPtr(fileBase64),
		FileName: common.StringPtr(fileName),
	}

	request := ess.NewUploadFilesRequest()
	request.BaseRequest.SetHttpMethod("POST")

	// 文件对应业务类型
	//1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	//2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	//3. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	// 上传pdf文件我们选择 DOCUMENT
	request.BusinessType = common.StringPtr("DOCUMENT")
	request.Caller = &ess.Caller{OperatorId: common.StringPtr(userId)}
	request.FileInfos = []*ess.UploadFile{uploadFile}

	return client.UploadFiles(request)
}
