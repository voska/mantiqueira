---
name: mantiqueira
description: >-
  Order eggs from Mantiqueira em Casa (www.mantiqueiraemcasa.com.br) using the
  `mantiqueira` CLI, and manage the recurring subscriptions the store is built
  around. Use for subscriptions, searching products, carts, delivery, and orders.
allowed-tools: Bash, Read
---

# mantiqueira

Order from Mantiqueira em Casa using the `mantiqueira` CLI.

## Always start here

```bash
mantiqueira doctor
```

Exit 0 means ordering will work. Any other exit code means it will not. Each
failed line prints the exact fix. Do that fix, or report it to the user. Do not
retry the command that failed.

## Subscriptions come first

This store sells mostly on recurring delivery, so most questions about it
("when do my eggs arrive?", "pause it while I travel") are subscription
questions, not order questions.

```bash
mantiqueira subs                  # status, frequency, next delivery
mantiqueira subs <id>             # one subscription's schedule and items
mantiqueira subs pause <id>       # pause indefinitely
mantiqueira subs resume <id>
mantiqueira subs skip <id>        # skip the next delivery only
mantiqueira subs unskip <id>
```

**Prefer `skip` to `pause`.** Skipping drops one delivery and the schedule
carries on by itself; pausing stops everything until someone remembers to
resume. "I'm away next week" means `skip`.

**There is no `cancel`.** VTEX has no transition out of `CANCELED`, so the CLI
does not offer it. If the user wants to cancel, tell them to do it on the
website — do not try to reach it another way.

Confirm the ID with `mantiqueira subs` before any pause, resume, or skip. These
change a real recurring grocery order.

## Ordering one-off

```bash
mantiqueira search ovos --limit 5      # 1. find a SKU
mantiqueira cart add 1 --qty 2         # 2. add it
mantiqueira delivery windows           # 3. pick a window number
mantiqueira checkout --window 0        # 4. preview — places nothing
mantiqueira checkout --window 0 --confirm   # 5. order
```

Between steps 4 and 5: **show the preview to the user and get explicit
approval.** `--confirm` spends real money.

**Search** takes Portuguese terms. The first column of the output is the SKU.

Every result carries both `sku` and `productId`. Commands take the `sku` —
the two are separate sequences and the same number routinely appears in
both, naming two unrelated products.

**Cart** never needs a `--seller`; it is looked up automatically.

```bash
mantiqueira cart show
mantiqueira cart update 0 --qty 3    # index from 'cart show'
mantiqueira cart remove 0
mantiqueira cart clear
```

**Payment** defaults to pix. `mantiqueira checkout payments` lists what the
store accepts and any card saved on the account. For a card, add `--cvv 123`.

## Output for scripts and agents

```bash
mantiqueira subs --json
mantiqueira search ovos --json --select sku,name,price
mantiqueira search ovos --plain          # tab-separated
mantiqueira search ovos --quiet --select sku   # bare values
```

Prices in `--json` are integer centavos: `6590` is R$65,90. On a subscription
item the field is `priceAtSubscriptionDate` — the price locked in when the
subscription started, which is not today's shelf price. Do not present it as
the current price.

Data goes to stdout; progress and errors go to stderr.

## Exit codes

| Code | Meaning | What to do |
|---|---|---|
| 0 | success | continue |
| 2 | bad arguments | fix the command; do not retry it unchanged |
| 3 | empty result | tell the user nothing matched |
| 4 | login required | `mantiqueira auth login --email <email>` |
| 5 | not found | the SKU or subscription ID is wrong; list again |
| 7, 8 | temporary | wait, then retry once |
| 9 | store rule refused | read the message; it names the rule |
| 10 | not set up | `mantiqueira doctor` and follow the fixes |

Full table: `mantiqueira exit-codes --json`

## Rules

- Never run `--confirm` without the user approving that exact cart and total.
- Never pause or skip a subscription the user did not name.
- If a command fails twice the same way, stop and report it. Do not loop.
- `mantiqueira doctor` diagnoses anything unexpected; its output names the fix.
