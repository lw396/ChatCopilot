# API Reference

## 获取群聊名称列表

### 描述

`http://localhost:6978/group_contact`

因为微信中可能存在多个名称相似或者完全相同的群聊，需要输入对应的群昵称获取群对应的 `user_name` 列表。

### 请求方式

`GET`

### Query Params

- | 参数名 | 类型 | 必填 | 描述 |
- | `nickname` | `string` | 是 | 群聊昵称 |

## 获取群聊基本信息

### 描述

`http://localhost:6978/message_info`

获取到群聊对应的 `user_name` 值后，我们需要查找该群聊的所在的数据库中， 因为微信将所有的聊天记分散存储在不同的 10 个数据库中，我们需要遍历这个 10 个数据库中的表，找到其所在的数据库。

### 请求方式

`GET`

### Query Params

- | 参数名 | 类型 | 必填 | 描述 |
- | `user_name` | `string` | 是 | -- |

## 保存群聊聊天记录

### 描述

保存历史聊天记录，并更新定时任务。

`http://localhost:6978/message_content`

### 请求方式

`POST`

### Body Params

- | 参数名 | 类型 | 必填 | 描述 |
- | `user_name` | `string` | 是 | -- |
- | `db_name` | `string` | 是 | 数据库名 |
