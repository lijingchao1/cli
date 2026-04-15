# base +record-retrieval-switch

> **前置条件：** 先阅读 [`../lark-shared/SKILL.md`](../../lark-shared/SKILL.md) 了解认证、全局参数和安全规则。

在指定 Base 上启用或停用 retrieval（语义检索）能力。启用后 `+record-retrieval` 才能返回有效结果。

## 适用场景

- 首次启用 Base 的语义检索能力。
- 关闭不再需要的 Base 的检索能力，以节省索引资源。

## 推荐命令

```bash
# 启用 retrieval
lark-cli base +record-retrieval-switch \
  --base-token app_xxx \
  --enable true

# 停用 retrieval
lark-cli base +record-retrieval-switch \
  --base-token app_xxx \
  --enable false
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--base-token <token>` | 是 | Base Token |
| `--enable <true\|false>` | 是 | 是否启用 retrieval；必须显式传 `true` 或 `false`，无默认值 |

## API 入参详情

**HTTP 方法和路径：**

```
PUT /open-apis/base/v3/bases/:base_token/records/retrieval-switch
```

**请求体：**

```json
{
  "enable": true
}
```

## 返回重点

- 直接返回接口 `data` 字段内容，通常含最新的开关状态。

## 工作流

1. 启用 retrieval 后，索引构建需要一定时间，刚写入的记录可能不会立即被检索到。
2. 停用 retrieval 后，再次启用时索引可能需要重新构建。
3. 对同一个 Base 频繁切换开关没有必要，保持稳定状态即可。

## 坑点

- ⚠️ 该操作会影响 Base 的检索行为（写入操作），切换前请确认目标。
- ⚠️ `--enable` 必须显式传 `true` 或 `false`，不传会被 CLI 拦截，传其他值会被 enum 校验拒绝。
- ⚠️ 切换后 `+record-retrieval` 的可用性和结果会随之变化。

## 参考

- [lark-base-record-retrieval.md](lark-base-record-retrieval.md) — 按语义检索记录
- [lark-base-record.md](lark-base-record.md) — record 索引页
