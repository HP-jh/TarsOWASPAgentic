package main

import (
	"fmt"
	"net/url"
	"strings"
)

// ValidateTarget 校验扫描目标是否合法
func ValidateTarget(target string) error {
	if strings.TrimSpace(target) == "" {
		return fmt.Errorf("目标不能为空")
	}
	// 基本 URL 或域名校验
	if strings.Contains(target, "://") {
		u, err := url.Parse(target)
		if err != nil {
			return fmt.Errorf("目标地址格式不正确: %w", err)
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return fmt.Errorf("仅支持 http/https 协议")
		}
		if u.Host == "" {
			return fmt.Errorf("目标缺少主机名")
		}
	} else {
		// 简单域名/IP 校验
		if strings.Contains(target, " ") {
			return fmt.Errorf("目标包含非法空格")
		}
	}
	// 禁止内网/本地地址（生产安全基线）
	lower := strings.ToLower(target)
	forbidden := []string{"localhost", "127.0.0.1", "0.0.0.0", "::1", "192.168.", "10.0.0.", "172.16."}
	for _, f := range forbidden {
		if strings.Contains(lower, f) {
			return fmt.Errorf("出于安全考虑，禁止扫描内网地址 (%s)", f)
		}
	}
	return nil
}
