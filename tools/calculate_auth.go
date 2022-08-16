package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"
)

const (
	algorithm = "TC3-HMAC-SHA256"
	defaultTimeout = 5
)

// ApiCaller 云api调用请求体
type ApiCaller struct {
	//用于标识 API 调用者身份，可以简单类比为用户名
	SecretId string
	//用于验证 API 调用者的身份，可以简单类比为密码
	SecretKey string
	//所调用的云api接口的服务地址
	Host string
	//所调用的云api接口的版本号
	Version string
	//所调用的云api接口的名字
	Action string
	//所调用的云api接口的所在区域
	Region string
	//所调用的云api接口的所属服务
	Service string
	//所调用的云api接口的post请求体
	Payload string
	//所调用的云api接口的临时token
	Token string
	//返回云api的响应
	Timeout int
	//请求协议，默认使用https
	Protocol string
}

func CallCloudApiV3(apiCaller ApiCaller) (string, error) {
	timestamp := time.Now().Unix()
	// 1.拼接规范请求串
	signedHeaders, canonicalRequest := buildCanonicalRequest(apiCaller.Host, apiCaller.Payload)
	// 2.拼接待签名字符串
	date, credentialScope, string2sign := buildStringToSign(timestamp, apiCaller.Service, canonicalRequest)
	// 3.计算鉴权签名
	authorization := buildAuthorization(date, string2sign, credentialScope, signedHeaders, apiCaller)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if apiCaller.Timeout == 0 {
		apiCaller.Timeout = defaultTimeout
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   cast.ToDuration(apiCaller.Timeout) * time.Second,
	}
	req, err := buildRequest(authorization, timestamp, apiCaller)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 构建原始请求
func buildCanonicalRequest(host string, payload string) (string, string) {
	// step 1: build canonical request string
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	return signedHeaders, canonicalRequest
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

// 构建签名字段
func buildStringToSign(timestamp int64, service string, canonicalRequest string) (string, string, string) {
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	return date, credentialScope, string2sign
}

// 构建鉴权签名
func buildAuthorization(date string, string2sign string, credentialScope string,
	signedHeaders string, caller ApiCaller) string {
	// step 3: sign json
	secretDate := hmacsha256(date, "TC3"+caller.SecretKey)
	secretService := hmacsha256(caller.Service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		caller.SecretId,
		credentialScope,
		signedHeaders,
		signature)
	return authorization
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	_, _ = hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

// 构建请求
func buildRequest(authorization string, timestamp int64, apiCaller ApiCaller) (*http.Request, error) {
	req, err := http.NewRequest("POST",
		apiCaller.Protocol + "://" + apiCaller.Host, strings.NewReader(apiCaller.Payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Host", apiCaller.Host)
	req.Header.Set("X-TC-Action", apiCaller.Action)
	req.Header.Set("X-TC-Timestamp", cast.ToString(timestamp))
	req.Header.Set("X-TC-Version", apiCaller.Version)
	req.Header.Set("X-TC-Region", apiCaller.Region)
	if apiCaller.Token != "" {
		req.Header.Set("X-TC-Token", apiCaller.Token)
	}
	return req, nil
}
