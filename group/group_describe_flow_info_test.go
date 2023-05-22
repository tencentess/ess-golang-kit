package group

import (
	ess_golang_kit "SdkTools"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

/**
 * 主企业代子企业查询合同信息的使用样例
 * 注意：使用集团代发起功能，需要主企业和子企业均已加入集团，并且主企业OperatorUserId对应人员被赋予了对应操作权限
 */
func TestGroupDescribeFlowInfo(t *testing.T) {
	// 子企业的合同id
	flowIds := []*string{common.StringPtr("")}
	// 需要代发起的子企业的企业id
	proxyOrganizationId := ""
	response, err := GroupDescribeFlowInfo(ess_golang_kit.OperatorUserId, flowIds, proxyOrganizationId)
	if err != nil {
		t.Errorf("DescribeFlowInfo error: %v", err)
	}
	t.Logf("DescribeFlowInfo finish: %v", response.ToJsonString())
}
