# AGENTS.md

## Project

Single-module Go CLI (`github.com/danmrichards/xkpassgo`, Go 1.26) that generates
XKCD-style memorable passwords.

**Entrypoint:** `cmd/xkpassgo/main.go`

## Build

- Word-list assets are embedded at compile time via `//go:embed`.
- Plain `go build` works — no external tool needed.
- Makefile cross-compiles for multiple platforms.

```makefile
make build           # builds linux + darwin + windows into ../../bin/
```

## Test & lint

CI runs in this order:
```
golangci-lint run -n ./...
go test -v -race ./...
```

- Randomness uses `crypto/rand` in production; tests use `math/rand` with fixed seed and golden-file assertions via a `config.Intner` interface.
- Linter config: `.golangci.yml`.

## Config

Three-layer precedence: **flags > JSON config file > defaults**.

- Default config path: `~/.xkpassgo.json`
- Flags defined with `spf13/pflag`, merged via `spf13/viper`.
- Config struct tags use `mapstructure`.

## Architecture

```
cmd/xkpassgo/main.go          # CLI entrypoint
pkg/config/                    # config loading + validation
pkg/assets/                    # embedded word list (//go:embed)
pkg/generator/                 # core password generation
pkg/transform/                 # case transformation (LOWER, UPPER, RANDOM, etc.)
pkg/separator/                 # separator insertion between words
pkg/padding/                   # digit + symbol padding (FIXED / ADAPTIVE)
```

The generation pipeline in `pkg/generator` is:
1. load embedded word list
2. pick `NumWords` random words (filtered by `WordLenMin`/`WordLenMax`)
3. apply case `transform`
4. insert `separator`
5. apply `padding` (digits then symbols)

## Notable quirks

- No external services, databases, or concurrency — purely a stdlib CLI with
  a few vendored dependencies (via `go.mod`).
- Randomness uses `crypto/rand` in production via `config.Intner` interface,
  seeded `math/rand` in tests.
