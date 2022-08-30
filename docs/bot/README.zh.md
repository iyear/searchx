**BOT**

[English](README.md) | 简体中文

## 使用

准备 `Telegram Bot`

- 获取 `Bot Token`: [如何创建 Bot](https://core.telegram.org/bots#6-botfather)

使用 `searchx`

- 下载 `searchx`：前往 [GitHub Releases](https://github.com/iyear/searchx/releases) 并解压
- 修改 `config/bot/config.min.yaml` 中 `YOUR_BOT_TOKEN` `YOUR_ADMIN_ID` 为 `Bot Token` `Telegram User ID`

使用 `run` 命令启动并邀请 `Bot` 至群组/频道

## 命令

- `-c`: 配置文件路径

### `run`

启动 Bot

```shell
# 使用最小化配置启动
./searchx bot run -c config/bot/config.min.yaml
```

### `source`

`Bot` 仅索引加入期间的含文本消息，如需索引历史消息请使用该命令。支持群组/频道，支持超大文件导入。

使用**官方客户端**导出历史消息。导出选项: **取消所有勾选、格式为 `JSON`**。

- `-f`: 导出的历史消息`JSON`

```shell
./searchx bot source -c config/bot/config.min.yaml -f YOUR_PATH/result.json
```

### `query`

命令行查询

- `-q`: 关键字
- `--pn`: 页码,从 0 开始,默认为 0
- `--ps`: 每页条数,默认为 10
- `--json`: 输出为 `JSON` 格式

```shell
./searchx bot query -c  -q KEYWORD --pn 0 --ps 10 --json
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

A: `Group Privacy` 开启后会导致 `Bot`
无法接收所有群组内的消息而导致索引缺失。具体请参考: https://core.telegram.org/bots#privacy-mode

同时，它不会产生安全问题。本项目 `Bot` 为自行部署，数据均存放于本地，不会导致数据泄露。
