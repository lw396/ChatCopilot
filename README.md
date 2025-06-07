# ChatCopilot

<div align="center"><img width = "150" height = "150" src="./assets/logo-universal.png" /></div>

## 项目介绍

`ChatCopilot` 是一款用 `golang` 实现的获取微信聊天记录并支持实时存储到数据库的工具，目前只支持在 `MacOS` 上可运行（因为我没有 windows 电脑）

### 当前实现功能

1. 群聊

- 通过群昵称获取群聊基本信息
- 将群聊历史记录保存至 mysql
- 实时更新接收到的群聊消息

2. 联系人聊天

- 通过昵称获取联系人基本信息
- 将聊天历史记录保存至 mysql
- 实时更新接收到的聊天消息

3. 消息处理

- 图片：明文显示图片说在路径
- 视频：明文显示视频说在路径
- 表情包：保存表情包图片到本地
- 语音：解码语音消息转为 wav 格式
- ......

## 使用方法

### 添加配置文件

将一下配置文件 `config/app.cfg` 添加到可执行文件的根目录下，修改你本机中对应的 `mysql` 和 `redis` 配置信息

```cfg
pod-id=1

[mysql]
host=127.0.0.1
port=3306
user=root
password=secret
db=ChatCopilot
timezone=Asia/Shanghai

[redis]
host=127.0.0.1
port=6379
auth=secret
db=0

[log]
dir=logs
max-age=7

[wechat]
key=
path=./test

[task]
interval=10
crontab=*/10 * * * * *
```

`wechat.key` 为微信数据库密钥，获取方式见[这里](doc/mac数据库解密.md)

`wechat.path` 为 `mac` 微信聊天记录的目录，具体需要看你电脑存放位置的实际情况

例：`/Users/james/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat/2.0b4.0.9/5a22781f14219edfffa333cb38aa92cf/Message`

注：路径中若存在空格，`不需要`在空格前添加 `\`

`task.interval`: 为执行同步任务的间隔时间，单位为 `秒`(范围为 1-59)，默认为 `10`，优先级高于 `task.crontab`

`task.crontab` : 当 `task.interval` 满足不了您执行任务的需求时，你可以使用 `crontab` 来设置定时任务

### 创建数据库表

执行 `migration` 文件夹中的 `sql` 语句创建对应的数据库表

### 执行可执行文件

在 `Releases` 中选择对应的系统下载可执行文件，解压后直接执行

#### 运行 `api` 服务

```sh
./chat-copilot api

```

对应的 `api` 文档[地址](doc/api.md)

#### 运行定时服务

```sh
./chat-copilot crontab

```

每 10 秒同步一次新的群聊记录

## 项目存在问题说明

1.当收到较大的原图图片或视频时，可能遇到微信没有自动下载原图的，在该情况下无法同步文件，目前只能手动点击图片下载后才可进行同步，目前未找到解决方案。

## 常见问题

### 未找到选项 `-L/usr/local/opt/openssl/lib` 的目录

```sh

export CGO_CFLAGS="-I/opt/homebrew/include"
export CGO_LDFLAGS="-L/opt/homebrew/lib"

```

### 读取同步图片或视频时，出现文件路径为空的情况

进入 微信 -> 设置 -> 通用 -> 勾选 文件设置中 `小于 20MB 的文件自动下载`，并将自动下载大小文件设置为 `1024MB`（最大只能设为 `1024MB`，所以在同步时未下载文件的请况下无法同步大于`1024MB`的文件）

![Alt](doc/img/wechat-file-setting.png)

## 参考文献

- 导出多年微信聊天记录 [https://sspai.com/post/82577](https://sspai.com/post/82577)

- PyWxDump [https://github.com/xaoyaoo/PyWxDump](https://github.com/xaoyaoo/PyWxDump)

- 使用 macOS 微信提取自定义表情 [https://blog.jogle.top/2022/08/14/macos-wechat-sticker-dump/](https://blog.jogle.top/2022/08/14/macos-wechat-sticker-dump/)

## 其他

[![Powered by DartNode](https://dartnode.com/branding/DN-Open-Source-sm.png)](https://dartnode.com "Powered by DartNode - Free VPS for Open Source")

- silk-v3-decoder [https://github.com/kn007/silk-v3-decoder](https://github.com/kn007/silk-v3-decoder)

- WeChatMsg [https://github.com/LC044/WeChatMsg](https://github.com/LC044/WeChatMsg)
