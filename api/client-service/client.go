package client_service

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

// 构造客户端调用实例
func GetClientInstance(secretId, secretKey, endPoint string) *ess.Client {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = endPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(secretId, secretKey), "ap-guangzhou", pf)
	return client
}
