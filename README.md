# Parol API

The HTTP API for Parol – the speech robot for Esperanto.

## Building

Run `make build` to create a development build of the API.

## Running

Run `dist/api` to bring up the API server.

```
→ ./dist/api
Parol API

Usage:
  api --conf=<config-file-path>
```

## Config

The API requires a TOML configuration file via the `--conf` flag.
The config file must contain the following values:

```toml
# address:port combination to bind to
bind = ":9000"

# AWS access key
aws-key = "..."

# AWS access secret
aws-secret = "..."

# specifies whether the API should run in local or production mode
development = true
```
