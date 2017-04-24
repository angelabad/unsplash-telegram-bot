# unsplash-telegram-bot

[![Build Status](https://travis-ci.org/angelabad/unsplash-telegram-bot.svg?branch=master)](https://travis-ci.org/angelabad/unsplash-telegram-bot)

[Telegram](https://telegram.org) bot for [Unsplash](https://unsplash.com) service in [aws api gateway](https://aws.amazon.com/api-gateway) and [aws lambda](https://aws.amazon.com/lambda).

This bot is built using [Go Programming Language](https://golang.org/) on top of [Sparta](http://gosparta.io/) library, you can see documentation about how to provision or test bot [here](http://gosparta.io/docs/overview/)

To run this bot you should have an AWS account, and it can incurs in costs. You also need a telegram bot created and its token and an unsplash api token.

You can use [Glide](https://glide.sh) to install vendor deps:

```bash
$ curl https://glide.sh/get | sh
$ glide install
```

(you could use only go get, but in this case its possible to have error with vendor versions)

You can define environment variables in your system (replace \*\*\* with your tokens):

```bash
BOT_TELEGRAMID=***
BOT_UNSPLASHID=***
```

Or if you prefer you can create a file called *config.json* with this content (replace \*\*\* with your tokens):

```json
{
  "unsplashID": "***",
  "telegramID": "***"
}
```

Now you can explore the service with:

```bash
$ go run !(*_test).go explore
```

You can provision bot, you need to have your aws credentials configured:

```bash
$ export AWS_REGION=us-west-2 (or another)
$ go run !(*_test).go provision --help
```

Provision with existing bucket

```bash
$ export AWS_REGION=us-west-2 (or another)
$ go run !(*_test).go provision -s YOURBUCKET
```
