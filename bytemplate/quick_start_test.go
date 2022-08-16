package bytemplate

import (
	"SdkTools"
	"SdkTools/api"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	"testing"
)

func TestCreateFlowByTemplate(t *testing.T) {
	// Step 1
	// 定义合同名
	flowName := "我的第一个合同"
	// 构造签署人
	// 此块代码中的$approvers仅用于快速发起一份合同样例，非正式对接用
	personName := "****************"   // 个人签署方的姓名，必须是真实的才能正常签署
	personMobile := "****************" // 个人签署方的手机号，必须是真实的才能正常签署
	approvers := []*ess.FlowCreateApprover{
		BuildPersonFlowCreateApprover(personName, personMobile),
	}

	// 如果是正式接入，需使用这里注释的approvers。请进入BuildFlowCreateApprovers函数内查看说明，构造需要的场景参数
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
	fileUrlResp, err := api.DescribeFileUrls(ess_golang_kit.OperatorUserId, flowId)
	if err != nil {
		panic(err)
	}
	t.Logf("请访问以下地址下载您的合同：\r\n")
	t.Logf(*fileUrlResp.Response.FileUrls[0].Url)
}
