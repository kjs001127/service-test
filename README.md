
# Channel App store Server 🛍️

## Getting started to develop 🛠️

### Install Golang

- Install Go at [here](https://go.dev/doc/install).
- Check installed Go version via `go version`

### Install tools & dependencies for develop

- Simply use `make init` to install [sqlboiler](https://github.com/volatiletech/sqlboiler), [mockery](https://github.com/vektra/mockery), [swaggo](https://github.com/swaggo/swag) & dependencies.

```bash
make init
```

- Simply user `make database-init` to create user & database.

```bash
make database-init
```

### Generate database schema & mock files

- Simply use `make generate`

```bash
make generate
```

### Run Local Server
```bash
make dev
```
- http://localhost:3020
- default port is 3020

### Run on other environments
```bash
GO_ENV={environment} make run
```
- development (default)
- exp
- ci
- production

### Check API document
- Create doc file with `make docs`
```bash
make docs
```
- Run local server via `make dev`
- Then you can access swaggo on http://localhost:3020/swagger/general/index.html
