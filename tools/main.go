package main

import (
	"encoding/json"
	"log"
)

const (
	secretId = "****"
	secretKey = "****"
	host = "ess.tencentcloudapi.com"
	version = "2020-11-11"
	region = "ap-guangzhou"
	service = "ess"
	protocol = "https"
	userId = "****"
	flowId = "****"
)

type UserInfo struct {
	UserId string
}

type DescribeFlowBriefsRequest struct {
	Operator *UserInfo
	FlowIds []string
}

type DescribeFlowBriefsResponse struct {
	Response *DescribeFlowBriefsResponseInner
}

type DescribeFlowBriefsResponseInner struct {
	FlowBriefs []*FlowBriefs
	RequestId  string
	Error      *ErrorResponse
}

type FlowBriefs struct {
	FlowId string
	FlowName string
	FlowDescription string
	FlowType string
	FlowStatus int64
	CreatedOn int64
	FlowMessage string
}

type ErrorResponse struct {
	Code    string
	Message string
}

func main() {
	action := "DescribeFlowBriefs"
	requestData := &DescribeFlowBriefsRequest{
		Operator: &UserInfo{UserId: userId},
		FlowIds: []string{flowId},
	}
	payload, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("marshal request fail: %v", err)
		return
	}
	apiCaller := ApiCaller{
		SecretId:  secretId,
		SecretKey: secretKey,
		Host:      host,
		Version:   version,
		Action:    action,
		Region:    region,
		Service:   service,
		Payload:   string(payload),
		Token:     "",
		Timeout:   10,
		Protocol:  protocol,
	}
	resp, err := CallCloudApiV3(apiCaller)
	if err != nil {
		log.Printf("call DescribeFlowBriefs fail: %v", err)
		return
	}

	respData := &DescribeFlowBriefsResponse{}
	if err := json.Unmarshal([]byte(resp), respData); err != nil {
		log.Printf("parse resp fail: %v", err)
		return
	}

	log.Printf("DescribeFlowBriefs result: %v", resp)
}
