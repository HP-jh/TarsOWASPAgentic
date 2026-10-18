# Changelog

## v1.0.1
- 添加上游请求重试机制（指数退避，最多 3 次）
- 新增 CHANGES.md 变更日志
- 增加重试相关单元测试

## v1.0.0
- MCP OWASP Agentic 集成基线
- 支持 stdio 与 HTTP 双模式
- owasp_scan / owasp_status 两个工具

## v1.0.2
- 结构化 HTTP 请求日志中间件（记录方法、路径、状态码、耗时、客户端地址）
- 新增 logger.go / logger_test.go

## v1.0.3
- 支持从 JSON 配置文件加载参数（endpoint / api_key / mode / timeout_seconds）
- 默认值自动补全，容错友好
- 新增 config.go / config_test.go

## v1.0.4
- 简单 Token Bucket 速率限制器（可配置 QPS 与间隔）
- HTTP 中间件保护所有端点
- 超限返回 429 + 小白化提示
- 新增 ratelimit.go / ratelimit_test.go

## v1.0.5
- 审计日志模块：内存缓存 + 可选文件持久化
- 记录每次工具调用的时间、工具名、目标、状态与错误
- 新增 audit.go / audit_test.go

## v1.0.6
- HTTP 连接池优化：MaxIdleConns=100、MaxIdleConnsPerHost=10、HTTP/2 强制尝试
- 细化拨号超时（5s）与保活（30s）
- 提取 client.go 便于复用与测试

## v1.0.7
- 小白化错误翻译层：把超时/拒绝连接/401/403/500 等翻译成普通人能看懂的话
- 新增 errors.go / errors_test.go

## v1.0.8
- 新增 /metrics 统计端点：uptime、请求总数/成功/失败、扫描总数/成功/失败
- 原子计数器保证并发安全
- 新增 metrics.go / metrics_test.go

## v1.0.9
- 扫描目标输入验证：空值、URL 格式、协议白名单、内网地址黑名单
- 安全基线：禁止扫描 localhost/127.0.0.1/192.168.x/10.x/172.16.x
- 新增 validate.go / validate_test.go

## v1.1.0
- 批量扫描 /batch 端点：并发扫描最多 50 个目标
- 可配置并发度（默认 3），信号量限流防上游打挂
- 自动调用 ValidateTarget 做批量前置校验
- 新增 batch.go / batch_test.go
