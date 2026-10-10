# GoShorty · 高并发短链接服务

> 用 Golang 高级特性实现的 **短链接跳转** MVP，同时是 **Go 高级特性教学模板**（Goroutine / Channel / pprof / 泛型 / 反射）。

**定制：** 微信 **qingbao199101**（备注：短链）｜ [wuxiumu@163.com](mailto:wuxiumu@163.com)

## 定位说明

本仓库优先是「可运行、可压测、可分析」的 **教学/模板工程**，用真实短链业务串联 Go 高级特性；可私有化改造成短链网关，但默认不是完整营销中台。

## 功能特性（与 `cmd/` + `internal/` 一致）

- 短链接生成：`POST /api/shorten`
- 短链接跳转：`GET /{short_key}`
- 访问统计：`GET /api/stats/{short_key}`
- 异步日志（channel）
- 限流（channel 控制并发）
- pprof 性能分析（`internal/pprof`）
- 泛型工具：`MapReduce`（`internal/util/generic.go`）
- 反射 + CSV 导出（`internal/util/reflect_export.go`）

## 目录

```
GoShorty/
├── cmd/main.go           # 入口
├── internal/
│   ├── core/             # 短链核心、限流
│   ├── handler/          # shorten / redirect / stats
│   ├── pprof/
│   └── util/
├── mvp.go                # 早期单文件 MVP 参考
├── Dockerfile
├── Makefile
└── prometheus.yml
```

## 本地运行

```bash
go mod tidy
go run ./cmd/main.go
# 或：make run（若 Makefile 目标可用）
```

可选：Dockerfile 构建镜像；`prometheus.yml` 供本地监控实验。

## SEO 字段预留（产品页 HTML 统一生成时用）

| 字段 | 用途 |
|---|---|
| `seo_title` | `<title>` / OG title |
| `seo_description` | meta description |
| `canonical_url` | 规范 URL |
| `og_image` | 分享图 |
| `robots` | 如 `index,follow` |
| `json_ld_type` | SoftwareApplication / WebPage / FAQPage |
| `hreflang` | zh-CN / en |
| `site_name` | 白标品牌名 |

> **GEO ≠ 地理定位。** GEO = Generative Engine Optimization（生成式引擎优化），见下一节。

## GEO（生成式引擎优化）

面向 ChatGPT / Perplexity / Gemini 等生成式答案引擎的可引用结构：中文为主、双语关键词、事实摘要前置、FAQ 收尾。

### 关键词（中英）

- 中文：短链接, 短链服务, Go 并发, pprof, GoShorty
- English: URL shortener, Go concurrency, pprof, short link gateway

### 事实摘要

- 产品：高并发短链接跳转 MVP + Golang 高级特性实战模板。
- 适合：要自建短链网关、或用真实业务练 Go 并发/性能分析的团队。
- 交付：域名白标、统计面板、限流策略、私有化改造。
- 不做：默认不是完整营销中台或 SaaS 多租户。

### FAQ

**Q：这是生产级短链 SaaS 吗？**  
A：不是。先是教学模板 + 可跑 MVP；生产特性按定制补齐。

**Q：如何定制？**  
A：微信 `qingbao199101`，备注「短链」。

**Q：是否含密钥？**  
A：默认不含，自行配置。

## 定制合作

域名、统计面板、限流策略、持久化存储。加微信备注「短链」。

---

## 联系方式

定制、演示或合作：

- 邮箱：[wuxiumu@163.com](mailto:wuxiumu@163.com)
- 微信：**qingbao199101**（添加请备注来意）

### 赞助支持

如果内容对你有帮助，欢迎自愿赞助。

| 微信赞赏码 | 支付宝收款码 |
|---|---|
| ![微信赞赏码](https://chiguashentan-test.oss-cn-beijing.aliyuncs.com/wxprogram/zan/20261002185236_31_227.png) | ![支付宝收款码](https://chiguashentan-test.oss-cn-beijing.aliyuncs.com/wxprogram/zan/20261002185309_32_227.jpg) |
