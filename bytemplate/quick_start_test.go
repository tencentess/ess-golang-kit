package bytemplate

import (
	"SdkTools"
	"SdkTools/api"
	"SdkTools/api/file-upload-download"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

// 如果您选择使用模板发起合同可以参考此处。通过此部分代码可以发起一份单C签署合同，帮您快速了解文件发起所必要的流程。您可以在体验后根据实际情况
// 修改参数，以满足业务场景的需求。
func TestCreateFlowByTemplate(t *testing.T) {
	// Step 1 准备合同参数
	// 定义合同名
	flowName := "我的第一个合同"

	// 构造签署人，设置签署人信息，您可以在此处填入自己的姓名和手机号，以体验合同签署的流程
	personName := "****************"   // 个人签署方的姓名，必须是真实的才能正常签署
	personMobile := "****************" // 个人签署方的手机号，必须是真实的才能正常签署
	approvers := []*ess.FlowCreateApprover{
		BuildPersonFlowCreateApprover(personName, personMobile),
	}

	// 这里为所有签署人类型的信息构造样例，可以根据实际情况进行调整
	//approvers := BuildFlowCreateApprovers()

	// Step 2
	// 发起合同
	flowId, schemeUrl, err := api.CreateFlowByTemplateDirectly(ess_golang_kit.OperatorUserId, flowName, approvers)

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
