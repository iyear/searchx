## 简介
![](https://img.shields.io/github/go-mod/go-version/iyear/searchx?style=flat-square)
![](https://img.shields.io/github/license/iyear/searchx?style=flat-square)
![](https://img.shields.io/github/v/release/iyear/searchx?color=red&style=flat-square)
![](https://img.shields.io/github/last-commit/iyear/searchx?style=flat-square)

[English](README.md) | 简体中文

🔎 五分钟强化 Telegram 群组/频道搜索 🚀

[DEMO](https://t.me/e5subs_bot) (需先加入[E5SubBot交流群](https://t.me/e5subbot))

## 特性

- 最小化配置 & 单文件一键启动
- 组件化扩展设计，可选多种存储后端
- 跨平台支持 macOS/Windows/Linux
- 特别为中文优化的搜索引擎
- 大文件历史消息导入
- 自定义国际化消息模板
- ……

## 部署

1. **准备 `Telegram Bot`**
    1. 获取 `Bot Token`: [如何创建 Bot](https://core.telegram.org/bots#6-botfather)
    2. **禁用** `Group Privacy`: 向 [@BotFather](https://t.me/BotFather) 发送 `/setprivacy`
2. **获取“我”的 `Telegram User ID`: [@userinfobot](https://t.me/userinfobot)**
3. **准备 `searchx`**
   1. 下载 `searchx`：前往 [GitHub Releases](https://github.com/iyear/searchx/releases) 并解压
   2. 修改 `config/config.min.yaml` 中 `YOUR_BOT_TOKEN` `YOUR_ADMIN_ID` 为 `Bot Token` `Telegram User ID`
4. **启动: `./searchx run -c config/config.min.yaml`**
5. **邀请 `Bot` 至群组/频道**

## 自定义
对于绝大部分用户，默认的最小化配置即可满足使用需求。

如果你有更多自定义配置需求，请参考 [config.full.yaml](config/config.full.yaml) 及注释。

## 命令
### version
查看版本信息

```shell
./searchx version
```

```
Version: 0.0.0
Commit: 5cae8dc
Date: 2022-07-30T07:58:05Z

go1.17.3 windows/amd64
```

### run
启动 Bot

- `-c`: 配置文件路径

```shell
# 使用最小化配置启动
./searchx run -c config/config.min.yaml
```

### source
Bot仅索引加入期间的含文本消息，如需索引历史消息请使用该命令。支持群组/频道，支持超大文件导入。

使用**官方客户端**导出历史消息。导出选项: **取消所有勾选、格式为 `JSON`**。

- `-f`: 导出的历史消息`JSON`
- `-d`: 搜索引擎
- `-o`: 搜索引擎选项

```shell
# 使用默认的 bleve 导入消息
./searchx source -f YOUR_PATH/result.json -d bleve -o path="index" -o dict="config/dict.txt"
```

### query
命令行查询

- `-d`: 搜索引擎
- `-o`: 搜索引擎选项
- `-q`: 关键字
- `--pn`: 页码,从 0 开始,默认为 0
- `--ps`: 每页条数,默认为 10
- `--json`: 输出为 `JSON` 格式

```shell
# 使用默认的 bleve 查询
./searchx query -d bleve -o path="index" -o dict="config/dict.txt" -q KEYWORD --pn 0 --ps 10 --json
```

## FAQ
**Q: 为什么我不使用 Telegram 自带的搜索？**

A: 众所周知，Telegram 自带的搜索功能并不好用，尤其是对中文的支持很差。本项目的目的就是解决这些搜索痛点。

**Q: 我在使用过程中遇到了问题？**
A: 在确认搜索后依旧无法解决，通过 [发起 ISSUE](https://github.com/iyear/searchx/issues/new) 的方式反馈。

在发起 `ISSUE` 的过程中，我们提倡使用英文描述问题，并在 `ISSUE` 中提供相关的截图和复现步骤。

**Q: 我想要增加一个功能？**

A: 同上

**Q: 为什么需要禁用 `Group Privacy`？它会造成安全问题吗？**

A: `Group Privacy` 开启后会导致 `Bot` 无法接收所有群组内的消息而导致索引缺失。具体请参考: https://core.telegram.org/bots#privacy-mode
同时，它不会产生安全问题。本项目 `Bot` 为自行部署，数据均存放于本地，不会导致数据泄露。

**Q: 可以用于个人账户吗？未来会支持吗？**

A: 目前不可以，未来可能会支持，但 `userbot` 的使用有封号的风险。本项目当前的目标群体是群主/频道主，自建为成员提供便利。

## LICENSE
Apache License 2.0
