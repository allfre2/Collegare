# Collegare

A first approach at go trying to get familiar with WebRTC.
Signaling server to enable WebRTC P2P calls.

## Dependencies

- [Gin Web Framework](https://gin-gonic.com/)
- [Gorilla WebSockets](https://github.com/gorilla/websocket)
- [envconfig](https://github.com/kelseyhightower/envconfig)
- ... In progress

## Build and Debug in VSCode

Copy .env file:

```bash
cp ./config/env/local.env .env
```

```bash
docker compose up
```

Attach to delve in container by running VSCode debug launch task.