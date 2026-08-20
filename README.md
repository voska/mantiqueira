# mantiqueira

CLI for [Mantiqueira em Casa](https://www.mantiqueiraemcasa.com.br), the direct-to-consumer
egg brand of Grupo Mantiqueira. Built for humans and AI agents: data to stdout, hints to
stderr, stable exit codes, structured output.

Built on [vtexkit](https://github.com/voska/vtexkit), a shared library for Brazilian VTEX
storefronts.

## Install

```sh
make build   # -> bin/mantiqueira
```

## Use

Start with `doctor`. It answers "can I order right now?" and prints the exact fix for
anything in the way.

```sh
mantiqueira doctor
mantiqueira auth login --email you@example.com
mantiqueira search ovos
```

### Subscriptions

Mantiqueira sells mostly on recurring delivery, so `subs` is the command that matters.

```sh
mantiqueira subs                  # every subscription: status, frequency, next delivery
mantiqueira subs <id>             # one subscription's schedule and items
mantiqueira subs pause <id>       # pause indefinitely
mantiqueira subs resume <id>
mantiqueira subs skip <id>        # skip the next delivery only
mantiqueira subs unskip <id>
```

```
$ mantiqueira subs
7A1C4E9B2F6D48A3B5C7E1F09D2B6A84   ACTIVE   every 14 days    next 2026-09-03   1 item(s)

$ mantiqueira subs 7A1C4E9B2F6D48A3B5C7E1F09D2B6A84
id         7A1C4E9B2F6D48A3B5C7E1F09D2B6A84
status     ACTIVE
frequency  every 14 days
next       2026-09-03
delivery   Receba às terças-feiras
cycles     25
  1        Kit Orgânico                                 x1      R$65,90
```

There is deliberately no `cancel`. VTEX has no transition out of `CANCELED`, so a mistyped
ID would destroy a subscription with no way back — cancel on the website instead.

Add `--json` to any command for agent-readable output, and see `mantiqueira exit-codes`
for the exit code contract.

## What lives here

Only the store descriptor. Everything else — VTEX client, auth, search, cart, checkout,
subscriptions, output modes, exit codes — is in vtexkit. Change behavior there, not here.

```
store.go             # The Mantiqueira descriptor — 4 fields
cmd/mantiqueira/     # Entry point
```

## License

MIT
