## Introduction
![](https://img.shields.io/github/go-mod/go-version/iyear/searchx?style=flat-square)
![](https://img.shields.io/github/license/iyear/searchx?style=flat-square)
![](https://img.shields.io/github/v/release/iyear/searchx?color=red&style=flat-square)
![](https://img.shields.io/github/last-commit/iyear/searchx?style=flat-square)

English | [简体中文](README_zh.md) | [DEMO](https://t.me/e5subs_bot)

🔍 Enhance Telegram Group/Channel Search In 5 Minutes 🚀

[DEMO](https://t.me/e5subs_bot) (Join [E5SubBot Group](https://t.me/e5subbot) in advance)

## Features

- Minimal configuration & single file one-click launch
- Component-based scalable design & multiple storage engine options
- Chinese optimization of search engine
- Large number of historical messages to import
- Customizable i18n message templates
- ......

## Deployment

1. **Prepare `Telegram Bot`**
    1. Get `Bot Token`: [How to create a Bot](https://core.telegram.org/bots#6-botfather)
    2. **Disable** `Group Privacy`: send `/setprivacy` to [@BotFather](https://t.me/BotFather)
2. **Get `Telegram User ID`: [@userinfobot](https://t.me/userinfobot)**
3. **Prepare `searchx`**
    1. Download `searchx`: go to [GitHub Releases](https://github.com/iyear/searchx/releases) and uncompress it.
    2. Change `YOUR_BOT_TOKEN` `YOUR_ADMIN_ID` in `config/config.min.yaml` to `Bot Token` `Telegram User ID`
4. **Start: `./searchx run -c config/config.min.yaml`**
5. **Invite `Bot` to groups/channels**

## Customization
For the vast majority of users, the default minimal configuration is sufficient. 

If you have more customization needs, please refer to [config.full.yaml](config/config.full.yaml) and its comments.

## Command
### version
View version information

```shell
./searchx version
```

```
Version: 0.0.0
Commit: 5cae8dc
Date: 2022-07-30T07:58:05Z

go1.17.3 windows/amd64
```

### source
Bot only indexes messages during the join period, use this command if you want to index history messages. 

Support groups/channels, support very large `JSON` file.

Use **official client** to export history messages. Export options: **Uncheck all, format as `JSON`**.

- `-f`: exported history messages in `JSON` file
- `-d`: search engine driver
- `-o`: search engine options

```shell
# Using bleve to import messages
./searchx source -f YOUR_PATH/result.json -d bleve -o path="index" -o dict="config/dict.txt"
```

### query
Command line query

- `-d`: search engine driver
- `-o`: search engine options
- `-q`: keyword
- `--pn`: page number, start from 0, default is 0
- `--ps`: number of result per page, default is 10
- `--json`: output as `JSON` format

```shell
# Using bleve to query
./searchx query -d bleve -o path="index" -o dict="config/dict.txt" -q KEYWORD --pn 0 --ps 10 --json
```

## FAQ
**Q: I'm having problems using it?**

A: If you still can't solve the problem after searching, give feedback by [SUBMIT ISSUE](https://github.com/iyear/searchx/issues/new).

When submit an `ISSUE`, we recommend describing the problem in English and providing relevant screenshots and steps to reproduce it.

**Q: I want to add a feature?**

A: Same as above

**Q: Why do I need to disable `Group Privacy`? Does it cause security issues?**

A: When `Group Privacy` is enabled, it will cause `Bot` to not receive all messages in the group, resulting in missing indexes. Please refer to: https://core.telegram.org/bots#privacy-mode
At the same time, it does not create security issues. The project is self-deployed and the data is stored locally, so it will not lead to data leakage.

## LICENSE
Apache License 2.0
