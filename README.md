# traefik-acme

A small Go CLI tool that converts a Traefik `acme.json` file into standalone certificate files (`<domain>.crt` and `<domain>.key`). It is useful when migrating from Let's Encrypt certificates managed by Traefik to static ones.

## How it works
1. Reads the ACME store — a JSON map of resolvers, each with an account and its certificates.
2. Base64-decodes each certificate and private key.
3. Writes them into an output directory as `<domain>.crt` / `<domain>.key`.

> ⚠️ `acme.json` contains private keys — treat it like a secret. Generated `.key` files are written with owner-only permissions (`0600`).

## Requirements
- Go 1.26+ (or Docker)

## Build
```bash
go mod download
go build -o build/traefik-acme .
```

## Usage

### List domains and certificates
```bash
./build/traefik-acme list
# with a custom acme.json:
./build/traefik-acme list -i /path/to/acme.json
```

### Export certificates
```bash
./build/traefik-acme get
# output to a specific directory (created if it does not exist):
./build/traefik-acme get -d ./certs -i ./acme/acme.json
```

#### Flags
| Flag | Default | Description |
| --- | --- | --- |
| `-i, --input` | `acme/acme.json` | Path to the Traefik `acme.json` file (global flag) |
| `-d, --dir` | `certs` | Output directory for `.crt` / `.key` files (`get` only) |

## Docker
```bash
docker build -t traefik-acme:v0.0.1 .
# mount acme.json at /acme/acme.json and an output dir at /certs:
docker run --rm -v ./acme:/acme -v ./certs:/certs traefik-acme get
```

## Project layout
| Path | Description |
| --- | --- |
| `cmd/` | CLI commands (`root`, `list`, `get`) built on cobra |
| `internal/resolver/` | Parsing of `acme.json`, base64 decoding, writing certificate files |
| `Dockerfile` | Multi-stage build (golang → scratch) |

## Tests
```bash
go test ./...
```
