# Profiling

```
brew install graphviz

go run cmd/training/main.go -cpu 2 -duration 2m

http://localhost:6060/debug/pprof

curl -o ./profiles/heap1.pprof "http://localhost:6060/debug/pprof/heap?seconds=5"
curl -o ./profiles/cpu1.pprof "http://localhost:6060/debug/pprof/profile?seconds=5"
curl -o ./profiles/heap2.pprof "http://localhost:6060/debug/pprof/heap?seconds=5"
curl -o ./profiles/cpu2.pprof "http://localhost:6060/debug/pprof/profile?seconds=5"

go tool pprof -http=:38080 ./profiles/cpu1.pprof
go tool pprof -http=:38081 ./profiles/cpu2.pprof
go tool pprof -http=:38080 ./profiles/heap1.pprof
go tool pprof -http=:38081 ./profiles/heap2.pprof
```


# Continuous Profiling

```
docker-compose -f ./pyroscope/docker-compose.yml up

http://localhost:3300/a/grafana-pyroscope-app/explore
```
