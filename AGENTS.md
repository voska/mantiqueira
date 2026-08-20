# Mantiqueira CLI

Go CLI for ordering eggs from Mantiqueira em Casa (`www.mantiqueiraemcasa.com.br`) in Brazil.

All logic lives in the shared library **`github.com/voska/vtexkit`**, which also powers the
Frescatto and Zona Sul CLIs. This repo holds only the store descriptor.

## Project Structure

```
store.go             # The Mantiqueira descriptor
cmd/mantiqueira/     # Entry point, ~15 lines
```

Everything else — VTEX client, auth strategies, search, cart, checkout, subscriptions,
output modes, exit codes — is in vtexkit. Change behavior there, not here.

## The one thing that is not derivable

The VTEX account is **`grupomantiqueira`**, not `mantiqueiraemcasa`. Every other store so
far has an account matching its domain, so `store.AccountName` falls back to the host.
This store must set `Account` explicitly. Getting it wrong breaks the auth cookie name and
the Subscriptions API host, both quietly.

## Subscriptions

This store sells on recurring delivery, so `subs` is the primary surface. The VTEX
Subscriptions (RNS) API resolves the account from the request subdomain, so it is
addressed on `grupomantiqueira.vtexcommercestable.com.br`, not the storefront domain —
handled inside vtexkit by `store.AccountBaseURL`.

`cancel` is intentionally absent: `CANCELED` is terminal in RNS.

## Build & test

`make build` `make test` `make lint` `make vet` `make ci`

## Commits

Conventional Commits (`feat:`, `fix:`, `chore:`, `docs:`, `test:`, `refactor:`).
