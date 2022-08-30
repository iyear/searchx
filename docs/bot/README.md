**BOT**

English | [简体中文](README.zh.md)

## Usage

Preparing the `Telegram Bot`

- Get `Bot Token`: [How to create a Bot](https://core.telegram.org/bots#6-botfather)

Using `searchx`

- Download `searchx`: Go to [GitHub Releases](https://github.com/iyear/searchx/releases) and unpack it
- Modify `YOUR_BOT_TOKEN` `YOUR_ADMIN_ID` in `config/bot/config.min.yaml` to `Bot Token` `Telegram User ID`

Use the `run` command to start and invite `Bot` to the group/channel

## Command

- `-c`: path to configuration file

### `run`

Start Bot

```shell
## start with a minimal config
./searchx bot run -c config/bot/config.min.yaml
```

### source

Bot only indexes messages during the join period, use this command if you want to index history messages.

Support groups/channels, support very large `JSON` file.

Use **official client** to export history messages. Export options: **Uncheck all, format as `JSON`**.

- `-f`: exported history messages in `JSON` file

```shell
./searchx source -c config/bot/config.min.yaml -f YOUR_PATH/result.json
```

### query

Command line query

- `--pn`: page number, start from 0, default is 0
- `--ps`: number of result per page, default is 10
- `--json`: output as `JSON` format

```shell
./searchx query -c config/bot/config.min.yaml -q KEYWORD --pn 0 --ps 10 --json
```

## FAQ

**Q: Why don't I use the search that comes with Telegram?**

A: As we all know, the search function that comes with Telegram is not very useful. The purpose of this project is to
solve these problems.

**Q: I'm having problems using it?**

A: If you still can't solve the problem after searching, give feedback
by [SUBMIT ISSUE](https://github.com/iyear/searchx/issues/new).

When submit an `ISSUE`, we recommend describing the problem in English and providing relevant screenshots and steps to
reproduce it.

**Q: I want to add a feature?**

A: Same as above

**Q: Why do I need to disable `Group Privacy`? Does it cause security issues?**

A: When `Group Privacy` is enabled, it will cause `Bot` to not receive all messages in the group, resulting in missing
indexes. Please refer to: https://core.telegram.org/bots#privacy-mode
At the same time, it does not create security issues. The project is self-deployed and the data is stored locally, so it
will not lead to data leakage.
