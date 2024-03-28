# ChatCopilot

## 项目介绍

`ChatCopilot` 是一款用 `golang` 实现的获取微信聊天记录并存储到数据库的脚本工具，目前只在 `MacOS` 上可运行（因为我没有 windows 电脑）。

### 当前实现功能

- 1、获取群聊记录

## 工具安装

```sh

brew install sqlite
brew install sqlcipher

```

## 使用方法

### 添加配置文件

将一下配置文件 `config/app.cfg` 添加到可执行文件的根目录下，修改你本机中对应的 `mysql` 和 `redis` 配置信息.

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
```

`wechat.key` 为微信数据库密钥，获取方式见[这里](doc/mac数据库解密.md)。

`wechat.path` 为 `mac` 微信聊天记录的目录，具体需要看你电脑存放位置的实际情况，例：`/Users/james/Library/Containers/com.tencent.xinWeChat/Data/Library/Application\ Support/com.tencent.xinWeChat/2.0b4.0.9/5a22781f14219edfffa333cb38aa92cf/Message`

### 创建数据库表

执行 `migration` 文件夹中的 `sql` 语句创建对应的数据库表.

### 执行可执行文件

在 `Releases` 中选择对应的系统下载可执行文件，解压后直接执行。

#### 运行 `api` 服务

```sh
./ChatCopilot api

```

对应的 `api` 文档[地址](doc/api.md)

#### 运行定时服务

```sh
./ChatCopilot crontab

```

每 30 秒同步一次新的群聊记录。

## 常见问题

### 未找到选项 `-L/usr/local/opt/openssl/lib` 的目录

```sh

export CGO_CFLAGS="-I/opt/homebrew/include"
export CGO_LDFLAGS="-L/opt/homebrew/lib"

```

## 参考文献

- 导出多年微信聊天记录 [https://sspai.com/post/82577](https://sspai.com/post/82577)

- PyWxDump [https://github.com/xaoyaoo/PyWxDump](https://github.com/xaoyaoo/PyWxDump)
