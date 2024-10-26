assembly:
	go build -gcflags="-S" ./pkg/math/statistics > statistics.S 2>&1

test:
	go test ./... --parallel 4 --cover -coverprofile=reports/coverage

build:
	go build ./...

coverage-report: test
	go tool cover -html=reports/coverage

benchmark:
	go test -benchmem -run=^$$ -bench . ./docs/benchmark --cpuprofile ./data/pgo/benchmark-cpu.pprof --memprofile ./data/pgo/benchmark-mem.pprof

pprof:
	go tool pprof --http=:8080 .\data\pgo\benchmark.pprof