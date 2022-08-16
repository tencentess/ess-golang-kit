package api

import (
	"SdkTools"
	"testing"
)

// 查询模板调用样例
func TestDescribeThirdPartyAuthCode(t *testing.T) {
	// 电子签小程序跳转客户小程序时携带的授权查看码
	authCode := "********************************"

	response, err := DescribeThirdPartyAuthCode(ess_golang_kit.OperatorUserId, authCode)
	if err != nil {
		t.Errorf("DescribeThirdPartyAuthCode error: %v", err)
	}
	t.Logf("DescribeThirdPartyAuthCode finish: %v", response.ToJsonString())
}
