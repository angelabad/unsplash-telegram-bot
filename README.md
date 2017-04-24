# unsplash-telegram-bot

[![Build Status](https://travis-ci.org/angelabad/unsplash-telegram-bot.svg?branch=master)](https://travis-ci.org/angelabad/unsplash-telegram-bot)

[Telegram](https://telegram.org) bot for [Unsplash](https://unsplash.com) service in [aws api gateway](https://aws.amazon.com/api-gateway) and [aws lambda](https://aws.amazon.com/lambda).

To run this bot you should have an AWS account, and it can incurs in costs. You also need a telegram bot created and its token and an unsplash api token.

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
