# kowa-sandbox

Dedicated **test** repository for the Kowa project (`SnWalker/kowa`). It is used for the whole Kowa development and debugging period to exercise real `git` / `gh` behavior (clone, commit, push of work branches, pull requests, reviews, merge observation, Runner/Provider runs) without ever touching the Kowa source repository or any production repository.

It contains a tiny, fake Go package (`calc`). Everything here is disposable.

## Rules

- **Private only.** Never make this repository public.
- **No secrets and no real data.** No tokens, keys, credentials, personal data, customer data or proprietary code, ever. Use fake values only.
- **Not a Kowa dependency.** Kowa code, CI and docs must not depend on this repository being present or in a particular state.
- `main` stays green and minimal. Test work happens on branches named `task/<id>` or `exp/<topic>`; delete them when the experiment is recorded.
- `fixture/*` branches hold reusable starting points (for example `fixture/failing-test`, a branch with one deliberately failing test for "fix the failing test" tasks). Branch from them, never push to them from experiments.
- No GitHub App installations, webhooks, secrets, Actions workflows or deploy keys unless a Kowa stage explicitly requires it and records it in `doc/Kowa验收与运行/` in the Kowa repository.
- Do not commit generated or large files. Do not store Kowa evidence here; evidence belongs in the Kowa repository records.

## Reset

Delete every `task/*` and `exp/*` branch and close open PRs. `main` and `fixture/*` are the only long-lived refs.

- `main`: if it was changed by mistake, restore it to the commit tagged `baseline` (`d4943b0`).
- `fixture/failing-test`: if it was changed by mistake, restore it to `0a62e29` — **not** to `baseline`. The fixture is `baseline` plus one commit that carries the deliberate failure.
- The authoritative SHAs live in the Kowa repository at `doc/Kowa验收与运行/测试仓库约定.md` section 2.

## Check

```bash
go test ./...
```
