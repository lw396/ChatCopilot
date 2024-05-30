# Project: ChatCopilot

## 💡Related templates

## End-point: 登录
### Method: POST
>```
>{{url}}/auth/login
>```
### Body (**raw**)

```json
{
    "username": "admin",
    "password": "admin"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 获取群聊基本信息

### Method: GET
>```
>{{base_url}}/group_contact?nickname=xxx
>```
### Body (**raw**)

```json

```

### Query Params

|Param|value|
|---|---|
|nickname|xxx|


### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 保存群聊天记录

### Method: POST
>```
>{{base_url}}/group_contact
>```
### Body (**raw**)

```json
{
    "user_name": "xxx"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 删除群聊天记录

### Method: DELETE
>```
>{{base_url}}/group_contact
>```
### Body (**raw**)

```json
{
    "user_name": "xxx"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 获取群聊天列表

### Method: GET
>```
>{{base_url}}/group_contact_list?nickname=xxx&offset=1
>```
### Body (**raw**)

```json

```

### Query Params

|Param|value|
|---|---|
|nickname|xxx|
|offset|1|


### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 获取联系人基本信息

### Method: GET
>```
>{{base_url}}/contact_person?nickname=xxx
>```
### Body (**raw**)

```json

```

### Query Params

|Param|value|
|---|---|
|nickname|xxx|


### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 保存联系人聊天记录

### Method: POST
>```
>{{base_url}}/contact_person
>```
### Body (**raw**)

```json
{
    "user_name": "xxx"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 删除联系人聊天记录

### Method: DELETE
>```
>{{base_url}}/group_contact
>```
### Body (**raw**)

```json
{
    "user_name": "xxx"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: 获取消息记录列表

### Method: GET

>```
>{{base_url}}/message_content_list?user_name=xxx&offset=1
>```
### Body (**raw**)

```json

```

### Query Params

|Param|value|
|---|---|
|user_name|xxx|
|offset|1|


### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|{{tokan}}|string|