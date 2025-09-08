# sandbox-go-webserver

Sandbox for learning how to use Go to build a webserver.

## Initialization

1. Prerequisite: Have [Go](https://go.dev/dl/), [`golangci-lint`](https://golangci-lint.run/docs/welcome/install/#local-installation) and [Make](https://www.gnu.org/software/make/) installed.
1. Run the [`./init.sh`](./init.sh) script from the root of the repo to initialize the codebase.

## Validation

- The `.githooks/` scripts validate the codebase before every commit and push.
- The validation commands can also be run manually from the root of the repo:

```sh
make validate
```
