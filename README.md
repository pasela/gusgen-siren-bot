# gusgen-siren-bot

gusgen-siren-bot is a Twitter bot that tweets Gusgen's siren and Vana'diel time on each Vana'diel day.

Twitter: [@GusgenSirenBot](https://twitter.com/gusgensirenbot)

## Configuration

Configuration file and environment variables are supported.

### Configuration file

Search path:

- `./conf/gusgen-siren-bot.yml`
- `./gusgen-siren-bot.yml`

Or specify the file path by `GUSGEN_SIREN_BOT_CONF` environment variable.

Contents of configuration file:

```yaml
notifier: "stdout" or "twitter"
twitter:
    consumer_key: CONSUMER KEY
    consumer_secret: CONSUMER SECRET
    access_token: ACCESS TOKEN
    access_secret: ACCESS SECRET
```

### Environment variables

| Name                    | Value                               |
| ----------------------- | ----------------------------------- |
| NOTIFIER                | stdout\|twitter                     |
| TWITTER_CONSUMER_KEY    | Consumer key for Twitter API        |
| TWITTER_CONSUMER_SECRET | Consumer secret for Twitter API     |
| TWITTER_ACCESS_TOKEN    | Access token for Twitter API        |
| TWITTER_ACCESS_SECRET   | Access token secret for Twitter API |

## Commandline options

Specify the notifier:

`--notifier=stdout|twitter`

## Docker

See `_example` directory.

## License

MIT

## Author

pasela
