# base +record-retrieval-poll

> **前置条件：** 先阅读 [`../lark-shared/SKILL.md`](../../lark-shared/SKILL.md) 了解认证、全局参数和安全规则。

使用 `+record-retrieval` 返回的 `retrieval_id` 轮询语义检索任务。任务仍在执行时返回状态；任务完成后返回原语义检索结果。

## 适用场景

- 已调用 `+record-retrieval` 发起语义检索任务。
- 需要查询任务是否完成，或读取最终命中的记录。

## 推荐命令

```bash
lark-cli base +record-retrieval-poll \
  --base-token app_xxx \
  --retrieval-id 7637482394915802653
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--base-token <token>` | 是 | Base Token |
| `--retrieval-id <id>` | 是 | `+record-retrieval` 返回的 `retrieval_id` |

## API 入参详情

**HTTP 方法和路径：**

```
GET /open-apis/base/v3/bases/:base_token/records/retrieval/:retrieval_id
```

## 返回重点

- 任务执行中通常返回 `status` 和 `status_message`，例如 `running`。
- 任务完成后返回语义检索结果，Payload 与 `+record-retrieval` 对应的最终结果一致。

## 工作流

1. 使用 `+record-retrieval-switch --enable true` 确认 Base 已开启语义检索。
2. 调用 `+record-retrieval` 发起任务，并记录返回的 `retrieval_id`。
3. 调用 `+record-retrieval-poll --retrieval-id <id>` 查询状态或最终结果。

## 坑点

- ⚠️ `retrieval_id` 只对发起该任务的用户和 Base 有效，跨 Base 或跨用户轮询会失败。
- ⚠️ 任务执行中返回的是状态，不是最终记录列表；需要间隔一段时间后再次轮询。
- ⚠️ 该接口作用于整个 Base（跨表），不能指定具体 `table-id`。

## 参考

- [lark-base-record-retrieval.md](lark-base-record-retrieval.md) — 发起语义检索
- [lark-base-record-retrieval-switch.md](lark-base-record-retrieval-switch.md) — 启用 / 停用 retrieval
- [lark-base-record.md](lark-base-record.md) — record 索引页
