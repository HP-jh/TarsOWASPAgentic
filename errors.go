package main

import "fmt"

// FriendlyError 把底层错误翻译成人话
func FriendlyError(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	switch {
	case contains(s, "timeout") || contains(s, "deadline exceeded"):
		return "连接超时，请检查网络或稍后重试"
	case contains(s, "connection refused"):
		return "目标服务未启动或无法访问"
	case contains(s, "no such host") || contains(s, "DNS"):
		return "无法解析目标地址，请检查域名是否正确"
	case contains(s, "401") || contains(s, "Unauthorized"):
		return "认证失败，请检查 API Key 是否正确"
	case contains(s, "403"):
		return "权限不足，请联系管理员"
	case contains(s, "404"):
		return "请求的资源不存在"
	case contains(s, "500") || contains(s, "502") || contains(s, "503"):
		return "上游服务暂时不可用，请稍后重试"
	default:
		return fmt.Sprintf("遇到一点小问题：%s", s)
	}
}

func contains(a, b string) bool {
	return len(a) >= len(b) && (a == b || len(a) > 0 && containsSub(a, b))
}

func containsSub(a, b string) bool {
	for i := 0; i <= len(a)-len(b); i++ {
		if a[i:i+len(b)] == b {
			return true
		}
	}
	return false
}
