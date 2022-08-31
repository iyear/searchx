**USR**

[English](README.md) | 简体中文

## 使用

准备 `Telegram Bot`

- 获取 `Bot Token`: [如何创建 Bot](https://core.telegram.org/bots#6-botfather)

使用 `searchx`

- 下载 `searchx`：前往 [GitHub Releases](https://github.com/iyear/searchx/releases) 并解压
- 修改 `config/usr/config.min.yaml` 中 `YOUR_BOT_TOKEN` 为 `Bot Token`

使用 `login` 命令登录 `Telegram`

使用 `run` 命令启动 `searchx`

## 命令

- `-c`: 配置文件路径

### `login`

登录 `Telegram`

```shell
./searchx usr login # 使用最小化配置
./searchx usr login -c my/config.yaml # 使用自定义配置
```

### `run`

启动 Bot

```shell
./searchx usr run # 使用最小化配置
./searchx usr run -c my/config.yaml # 使用自定义配置
```

### `source`
索引传入的时间戳内的所有对话消息

- `--from`: 开始时间戳，默认为 0
- `--to`: 结束时间戳，默认为当前时间戳

```shell
./searchx usr source # 使用最小化配置，索引所有消息
./searchx usr source -c my/config.yaml # 使用自定义配置，索引所有消息
./searchx usr source -c my/config.yaml --from 1661703949 --to 1661903949 # 使用自定义配置，索引 2021-09-08 00:00:00 至 2021-09-10 00:00:00 之间的消息
```

## FAQ

**Q: 使用 `userbot` 会有封号风险吗？**

A: 有，但 `searchx` 通常不会造成封号。`searchx` 只读取消息，但不会发送、编辑任何消息。

相关链接: https://github.com/gotd/td/blob/main/.github/SUPPORT.md#how-to-not-get-banned
