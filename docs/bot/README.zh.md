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
./searchx bot run # 使用最小化配置
./searchx bot run -c my/config.yaml # 使用自定义配置
```

### `source`

`Bot` 仅索引加入期间的含文本消息，如需索引历史消息请使用该命令。支持群组/频道，支持超大文件导入。

使用**官方客户端**导出历史消息。导出选项: **取消所有勾选、格式为 `JSON`**。

- `-f`: 导出的历史消息`JSON`

```shell
./searchx bot source # 使用最小化配置，使用默认文件
./searchx bot source -c my/config.yaml -f my/result.json # 使用自定义配置和文件
```

### `query`

命令行查询

- `-q`: 关键词
- `--pn`: 页码,从 0 开始,默认为 0
- `--ps`: 每页条数,默认为 10
- `--json`: 输出为 `JSON` 格式

```shell
./searchx bot query -q KEYWORD # 使用最小化配置
./searchx bot query -c my/config.yaml -q KEYWORD --pn 1 --ps 7 --json # 使用自定义配置
```

## FAQ

**Q: 为什么需要禁用 `Group Privacy`？它会造成安全问题吗？**

A: `Group Privacy` 开启后会导致 `Bot`
无法接收所有群组内的消息而导致索引缺失。具体请参考: https://core.telegram.org/bots#privacy-mode

同时，它不会产生安全问题。本项目 `Bot` 为自行部署，数据均存放于本地，不会导致数据泄露。
