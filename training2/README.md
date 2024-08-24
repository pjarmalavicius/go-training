# Dependencies

```
go run cmd/training/main.go
go get rsc.io/quote
go run cmd/training/main.go
go mod tidy
go env GOPATH
ls -l $(go env GOPATH)/pkg/mod
rm -rf $(go env GOPATH)/pkg/mod
go run cmd/training/main.go
go mod download
```
