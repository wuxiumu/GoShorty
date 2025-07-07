# GoShorty

🎯 一个用于演示 Golang 高级特性（Goroutine、Channel、pprof、泛型、反射等）的高并发短链接跳转服务。

---

## 🚀 功能特性

- 短链接生成：`POST /api/shorten`
- 短链接跳转：`GET /{short_key}`
- 访问统计查看：`GET /api/stats/{short_key}`
- 异步日志采集（channel 非阻塞）
- 限流机制（channel 控制并发）
- pprof 支持性能分析
- 泛型工具函数：`MapReduce`
- 反射 + CSV 导出

---

## 📦 本地运行

```bash
go mod tidy
go run ./cmd/main.go

```

好的，以下是对这个项目 **《Golang 高级编程实战模板工程》** 的整体结构与功能模块进行的**重新梳理**，便于你总结、复用、打磨成长期可维护的架构学习工程：



------





## **📦 项目名称：GoShorty（Golang 高级编程实战模板）**







### **🔍 项目目标：**





构建一个可运行、可压测、可分析的微型服务，通过真实业务功能串联起 **Golang 的高级特性与性能调优实践**，并提供易扩展的模板工程架构。



------





## **🧱 一、功能结构总览**



| **模块** | **说明**                                       |
| -------- | ---------------------------------------------- |
| URL 缩短 | POST /api/shorten 接收长链接生成短码           |
| URL 跳转 | GET /abc123 重定向至原始链接                   |
| 访问统计 | GET /api/stats/abc123 查询访问次数             |
| 日志采集 | 异步日志记录到控制台（channel 实现）           |
| 数据导出 | 支持导出访问记录为 CSV（reflect+encoding/csv） |



------





## **🚀 二、Golang 高级特性覆盖点**



| **特性**       | **实践位置**                       | **验证方式**                   |
| -------------- | ---------------------------------- | ------------------------------ |
| Goroutine 调度 | 每次跳转为独立协程                 | 使用 pprof 查看数量与堆栈      |
| Channel 通信   | 日志异步处理 / 限流控制            | 滥用可导致堵塞，合理使用有缓冲 |
| 逃逸分析       | &Link{}、map[string]*Link          | go run -gcflags="-m" main.go   |
| 内存对齐       | Link 字段排序                      | 使用 unsafe.Sizeof()           |
| 泛型函数       | MapReduce[T, R]                    | 自定义类型运行聚合             |
| 反射操作       | exportCSV() 动态字段访问           | 输出 CSV 内容可见字段值        |
| pprof 调试     | 自动挂载 /debug/pprof/             | 浏览器可视化或 CLI 工具分析    |
| 限流机制       | channel 限制并发数（如最多10跳转） | wrk 压测时触发 429             |



------





## **🧪 三、压测与调试建议**







### **✅ 工具推荐**



| **工具**       | **用法**                      |
| -------------- | ----------------------------- |
| wrk            | 并发压测性能瓶颈              |
| hey            | 快速验证接口正确性            |
| pprof          | Goroutine / Heap / Block 分析 |
| go test -bench | 基准测试（扩展可加）          |
| curl / Postman | 手动验证 API 行为             |



------





## **📁 四、建议文件结构（可拆为包）**



```
/cmd
  main.go                // 服务入口
/internal
  /handler
    shorten.go           // POST + 跳转逻辑
    stats.go             // 查看统计
  /core
    link.go              // Link结构定义与操作
    logger.go            // 异步日志器
    limiter.go           // channel 限流器
  /util
    reflect_export.go    // CSV导出
    generic.go           // MapReduce 泛型
    pprof.go             // pprof启动辅助
```



------





## **🔧 五、可选拓展方向（持续进阶）**



| **拓展项**               | **目的**                    |
| ------------------------ | --------------------------- |
| 使用 Redis 作为存储层    | 观察连接池、缓存策略        |
| 接入 Prometheus          | 更专业监控指标采集          |
| 增加 JWT 登录鉴权        | 演示上下文传递与 token 校验 |
| 使用 context.Context     | 演示跨 goroutine 中止、超时 |
| Dockerfile + Makefile    | 模板工程可部署化            |
| 集成 gRPC                | 高性能服务间通信演示        |
| 加入测试用例 & Benchmark | 展示 *_test.go 工程规范     |



------





## **📜 六、一键运行与测试说明（README 示例）**



```
go run main.go

# 生成短链接
curl -X POST http://localhost:8080/api/shorten -d "url=https://golang.org"

# 跳转链接
curl -v http://localhost:8080/abc123

# 查看统计
curl http://localhost:8080/api/stats/abc123

# 压测入口
wrk -t4 -c100 -d10s http://localhost:8080/abc123

# pprof 查看 goroutine
http://localhost:8080/debug/pprof/goroutine?debug=1
```



------





## **🧩 七、核心设计思想总结**





> 用一个小服务，把 Go 的高级特性一次性串通，跑得动、压得住、看得懂。





- 🧠 学习原理：不是纸上谈兵，是写代码时看见它怎么跑
- 🔍 学习调优：调度、逃逸、channel 使用都能实测对比
- 🎯 适合作为：技术面试展示工程、内训培训课、性能调优手册





------



如需我继续帮你：



- 拆分目录结构（生成实际文件）
- 添加 Prometheus metrics 接口
- 编写 README、Makefile、Dockerfile
- 添加单元测试框架





只需告诉我「下一步做什么」，我可以继续帮你完善。