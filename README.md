# Pocket Bot

Pocket - is a Telegram bot that allows you to save links in the [Pocket](https://getpocket.com/ru/?=) application. We can say that he is a small client for this service.

To work with the Pocket API, the SDK - [go-pocket-sdk](https://github.com/zhashkevych/go-pocket-sdk) is used.

[Bolt DB](https://github.com/boltdb/bolt) is used as storage.

To implement user authorization, an HTTP server is launched together with the bot on port 80, to which a redirect from Pocket occurs when the user is successfully authorized.

When the server accepts the request, it generates an Access Token via the Pocket API for the user and stores it in storage.

### Stack: 
  Go 1.19; 
  BoltDB; 
  Docker (for deployment)
