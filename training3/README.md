# Upgrading dependencies

```
go get golang.org/x/text
go get rsc.io/sampler
go test ./...
go list -m -versions rsc.io/sampler
go get rsc.io/sampler@v1.3.1
go test ./...
```
