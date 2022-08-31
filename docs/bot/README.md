**BOT**

English | [简体中文](README.zh.md)

## Usage

Preparing the `Telegram Bot`

- Get `Bot Token`: [How to create a Bot](https://core.telegram.org/bots#6-botfather)

Using `searchx`

- Download `searchx`: Go to [GitHub Releases](https://github.com/iyear/searchx/releases) and unpack it
- Modify `YOUR_BOT_TOKEN` `YOUR_ADMIN_ID` in `config/bot/config.min.yaml` to `Bot Token` `Telegram User ID`

Use the `run` command to start and invite `Bot` to the group/channel

## Commands

- `-c`: path to configuration file, default is `config/bot/config.min.yaml`

### `run`

Start Bot

```shell
./searchx bot run # use the min configuration
./searchx bot run -c my/config.yaml # use the specified configuration
```

### `source`
Bot only indexes messages during the join period, use this command if you want to index history messages.

Support groups/channels, support very large `JSON` file.

Use **official client** to export history messages. Export options: **Uncheck all, format as `JSON`**.

- `-f`: exported history messages in `JSON` file, default: `result.json`

```shell
./searchx bot source # use the min config, and the default file
./searchx bot source -c my/config.yaml -f my/result.json # use the specified config and file
```

### `query`

Command line query

- `-q`: keyword
- `--pn`: page number, start from 0, default is 0
- `--ps`: number of result per page, default is 10
- `--json`: output as `JSON` format, default is `false`

```shell
./searchx bot query -q KEYWORD # use the min config, and the default value
./searchx bot query -c my/config.yaml -q KEYWORD --pn 1 --ps 7 --json # use the specified config and value
```

## FAQ

**Q: Why do I need to disable `Group Privacy`? Does it cause security issues?**

A: When `Group Privacy` is enabled, it will cause `Bot` to not receive all messages in the group, resulting in missing
indexes. Please refer to: https://core.telegram.org/bots#privacy-mode
At the same time, it does not create security issues. The project is self-deployed and the data is stored locally, so it
will not lead to data leakage.
