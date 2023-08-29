package personAccountManagement

import (
	essGolangKit "SdkTools"
	"testing"
)

// 查询模板调用样例
func TestDescribeThirdPartyAuthCode(t *testing.T) {

	authCode := "********************************"

	response, err := DescribeThirdPartyAuthCode(essGolangKit.OperatorUserId, authCode)
	if err != nil {
		t.Errorf("DescribeThirdPartyAuthCode error: %v", err)
	}
	t.Logf("DescribeThirdPartyAuthCode finish: %v", response.ToJsonString())
}
