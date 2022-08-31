**USR**

English | [简体中文](README.zh.md)

## Usage

Preparing the `Telegram Bot

- Get `Bot Token`: [How to create a Bot](https://core.telegram.org/bots#6-botfather)

Using `searchx`

- Download `searchx`: Go to [GitHub Releases](https://github.com/iyear/searchx/releases) and unzip it
- Change `YOUR_BOT_TOKEN` to `Bot Token` in `config/usr/config.min.yaml`

Log in to `Telegram` with the `login` command

Start `searchx` with the `run` command

## Commands

- `-c`: path to configuration file, default: `config/usr/config.min.yaml`

### `login`

Login to `Telegram`

```shell
./searchx usr login # start with default min config
./searchx usr login -c my/config.yaml # start with specified config
```

### `run`

Start Bot

```shell
./searchx usr run # start with a min config
./searchx usr run -c my/config.yaml # start with specified config
```

### `source`
Index all dialog messages within the specific timestamp

- `--from`: start timestamp, default is 0
- `--to`: end timestamp, defaults to current timestamp

```shell
./searchx usr source # start with a min config, index all messages
./searchx usr source -c my/config.yaml # start with specified config, index all messages
./searchx usr source -c my/config.yaml --from 1661703949 --to 1661903949 # start with specified config, index messages from 2021-09-08 00:00:00 to 2021-09-10 00:00:00
```

## FAQ

**Q: Is there any risk of blocking with `userbot`?**

A: Yes, but `searchx` does not usually cause blocking. `searchx` only reads messages, but does not send or edit any messages.

Related link: https://github.com/gotd/td/blob/main/.github/SUPPORT.md#how-to-not-get-banned
