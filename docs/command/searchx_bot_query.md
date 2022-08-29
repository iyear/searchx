## searchx bot query

Query messages

```
searchx bot query [flags]
```

### Examples

```
searchx bot query -q hello --pn 0 --ps 15 --json
```

### Options

```
  -h, --help           help for query
      --json           json format output
      --pn int         page number, starting from 0
      --ps int         page size (default 10)
  -q, --query string   query keyword or statement
```

### Options inherited from parent commands

```
  -c, --config string   the path to the config file (default "config/bot/config.min.yaml")
```

### SEE ALSO

* [searchx bot](searchx_bot.md)	 - Official Telegram Bot for group/channel owner

