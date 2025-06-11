##  2025年6月10日
- 完成了 RAG 场景下的 MVP。
- 但在测试的过程中发现，RedisRetriever 需要使用索引来索引到数据文件，但是现在一切规则都符合，index 创建了之后 `JARVIS:doc:` 索引无法索引到相对应的数据

问题解决：索引中向量的维度定义与 embedding 模型生成的维度不一致，导致 index 不能自动扫描到数据（text-embedding-v1 向量维度为1536）。

Redis Stack 中索引自动追踪数据的前提：
- 前缀必须一致 `JARVIS:doc:` 为索引前缀，那么追踪的数据 Key 格式也肯定是 `JARVIS:doc:{doc id}`
- 索引定义与数据的字段定义必须完全严格一致

## 2025年6月11日
### 问题
- Redis Retriever 不能根据 query 目标字符串检索到相关的文档数据

### 问题解决
1. 不能检索到相关文档数据 
   - 在 `RetrieverConfig` 配置中，由于使用了自动填充结构体字段，Go 自动将配置中的 `DistanceThreshold` 字段初始化为 `new(float64)` 也就是零值，导致了无法检索到数据，注释掉后成功运行。


