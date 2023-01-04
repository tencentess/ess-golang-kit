package byfile

import (
	"SdkTools"
	"SdkTools/api"
	"SdkTools/api/file-upload-download"
	"encoding/base64"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"os"
	"testing"
)

// 如果您选择使用文件发起合同可以参考此处。通过此部分代码可以发起一份单C签署合同，帮您快速了解文件发起所必要的流程。您可以在体验后根据实际情况
// 修改参数，以满足业务场景的需求。
func TestCreateFlowByFile(t *testing.T) {
	// Step 1 准备合同参数

	// 定义合同文件所在的路径，一般文件发起的场景下我们需要先准备好合同文件，且为pdf格式。后续的填写、签章等操作均在此文件的基础上进行。
	filePath := "../testdata/test.pdf"

	// 定义合同名
	flowName := "我的第一个合同"

	// 构造签署人，设置签署人信息，您可以在此处填入自己的姓名和手机号，以体验合同签署的流程
	personName := "****************"   // 个人签署方的姓名，请保证信息的真实可用，否则合同不能签署
	personMobile := "****************" // 个人签署方的手机号，请保证信息的真实可用，否则合同不能签署

	// 这里我们构建一个单方个人签署的场景
	approvers := []*ess.ApproverInfo{
		BuildPersonApprover(personName, personMobile),
	}

	// 这里为所有签署人类型的信息构造样例，可以根据实际情况进行调整
	// approvers := BuildApprovers()

	// Step 2 上传文件并发起合同
	// 读取文件内容，并使用Base64进行编码
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fileBase64 := base64.StdEncoding.EncodeToString(bytes)

	// 将文件进行上传然后发起合同，并返回签署链接
	flowId, schemeUrl, err := api.CreateFlowByFileDirectly(ess_golang_kit.OperatorUserId, fileBase64, flowName, approvers)
	if err != nil {
		panic(err)
	}

	// 返回合同Id
	t.Logf("您创建的合同id为：\r\n")
	t.Logf(flowId)
	t.Logf("\r\n\r\n")
	// 返回签署的链接
	t.Logf("签署链接（请在手机浏览器中打开）为：\r\n")
	t.Logf(schemeUrl)
	t.Logf("\r\n\r\n")

	// Step 3
	// 下载合同
	fileUrlResp, err := file_upload_download.DescribeFileUrls(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		panic(err)
	}
	t.Logf("请访问以下地址下载您的合同：\r\n")
	t.Logf(*fileUrlResp.Response.FileUrls[0].Url)
}
