package byfile

import (
	"SdkTools"
	"SdkTools/api"
	"encoding/base64"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"os"
	"testing"
)

func TestCreateFlowByFile(t *testing.T) {

	// Step 1
	// 定义文件所在的路径
	filePath := "../testdata/test.pdf"

	// 定义合同名
	flowName := "我的第一个合同"

	// 构造签署人
	// 此块代码中的$approvers仅用于快速发起一份合同样例，非正式对接用
	personName := "****************"   // 个人签署方的姓名，必须是真实的才能正常签署
	personMobile := "****************" // 个人签署方的手机号，必须是真实的才能正常签署

	approvers := []*ess.ApproverInfo{
		BuildPersonApprover(personName, personMobile),
	}

	// 如果是正式接入，需使用这里注释的approvers。请进入BuildApprovers函数内查看说明，构造需要的场景参数
	// approvers := BuildApprovers()

	// Step 2
	// 将文件处理为Base64编码后的文件内容
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fileBase64 := base64.StdEncoding.EncodeToString(bytes)

	// 发起合同
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
	fileUrlResp, err := api.DescribeFileUrls(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		panic(err)
	}
	t.Logf("请访问以下地址下载您的合同：\r\n")
	t.Logf(*fileUrlResp.Response.FileUrls[0].Url)
}
