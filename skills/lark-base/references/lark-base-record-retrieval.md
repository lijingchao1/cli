# base +record-retrieval

> **前置条件：** 先阅读 [`../lark-shared/SKILL.md`](../../lark-shared/SKILL.md) 了解认证、全局参数和安全规则。

按语义化查询（semantic query）在 Base 维度上检索记录，返回匹配的记录列表。该接口使用 Base 内置的向量检索能力，适合自然语言描述的查询场景。

## 适用场景

- 基于自然语言描述查找相关记录（例如"和合同履约相关的记录"）。
- 已启用该 Base 的 retrieval 开关（见 `+record-retrieval-switch`）。
- 不知道具体字段或关键词，只有大致的语义描述。

## 推荐命令

```bash
lark-cli base +record-retrieval \
  --base-token app_xxx \
  --query "与合同风险相关的记录"
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--base-token <token>` | 是 | Base Token |
| `--query <string>` | 是 | 检索查询字符串，支持自然语言描述 |

## API 入参详情

**HTTP 方法和路径：**

```
POST /open-apis/base/v3/bases/:base_token/records/retrieval
```

**请求体：**

```json
{
  "query": "与合同风险相关的记录"
}
```

## 返回重点

- 直接返回接口 `data` 字段内容，通常包含命中的记录列表。

## 工作流

1. 调用前请先确认目标 Base 已经启用 retrieval（使用 `+record-retrieval-switch --enable` 启用）。
2. 未启用时调用该接口可能返回空结果或错误。

## 坑点

- ⚠️ 该接口作用于整个 Base（跨表），不能指定具体 `table-id`。
- ⚠️ 需要 Base 级别启用 retrieval 开关；未开启时检索会失效。
- ⚠️ 检索效果依赖底层索引构建进度，新写入的记录可能未被索引。
- ⚠️ `--query` 非空，建议提供具体语义描述而非单词。

## 参考

- [lark-base-record-retrieval-switch.md](lark-base-record-retrieval-switch.md) — 启用 / 停用 retrieval
- [lark-base-record.md](lark-base-record.md) — record 索引页
- [lark-base-record-search.md](lark-base-record-search.md) — 关键词搜索记录（按字段精确匹配）
