package ess_golang_kit

const (
	// 调用API的密钥对，通过腾讯云控制台获取
	SecretId  = "**************"
	SecretKey = "**************"

	// 经办人Id，电子签控制台获取
	OperatorUserId = "**************"

	// 模板Id，电子签控制台获取，仅在通过模板发起时使用
	TemplateId = "**************"

	// 企业方静默签用的印章Id，电子签控制台获取
	ServerSignSealId = "****************"

	// API域名，现网使用 ess.tencentcloudapi.com
	EndPoint = "ess.test.ess.tencent.cn"
	// 文件服务域名，现网使用 file.ess.tencent.cn
	FileEndPoint = "file.test.ess.tencent.cn"
)
