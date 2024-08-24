# Upgrading major versions

```
# won't show major versions
go list -m -versions rsc.io/quote

go get rsc.io/quote/v3
```

# GOPROXY and GOPRIVATE

```
go get -x rsc.io/quote
go env GOPROXY
go env -w GOPROXY=direct
go get -x rsc.io/quote
curl 'https://rsc.io/quote?go-get=1'

# private modules
go get -x github.com/vinted/go-vinted-slog
go env GOPRIVATE
go env -w GOPRIVATE=github.com/vinted
go get github.com/vinted/go-vinted-slog
go mod tidy
go env -w GOPRIVATE=
go env -w GOPROXY=https://proxy.golang.org,direct

# private proxy
go env -w GOPROXY=https://artefact-storage.svc.vinted.com/repository/go-proxy-repos-group/
go env -w GONOSUMDB=github.com/vinted
go get -x github.com/vinted/go-vinted-slog

# reset go env
go env -w GONOSUMDB=
go env -w GOPROXY=https://proxy.golang.org,direct
```

# go fmt

```
go fmt ./...
```
