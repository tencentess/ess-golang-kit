package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm3"
)

// CalcMd5 ...
func CalcMd5(content string) string {
	md5Instance := md5.New()
	result := md5Instance.Sum([]byte(content))
	return hex.EncodeToString(result)
}

// CalcSha1 ...
func CalcSha1(content string) string {
	sha1Instance := sha1.New()
	result := sha1Instance.Sum([]byte(content))
	return hex.EncodeToString(result)
}

// CalcSha256 ...
func CalcSha256(content string) string {
	sha256Instance := sha256.New()
	result := sha256Instance.Sum([]byte(content))
	return hex.EncodeToString(result)
}

// CalcSm3 ...
func CalcSm3(content string) string {
	sm3Instance := sm3.New()
	result := sm3Instance.Sum([]byte(content))
	return hex.EncodeToString(result)
}
