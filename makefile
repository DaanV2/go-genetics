assembly:
	go build -gcflags="-S" ./pkg/nucleic_acid_codes > nucleic_acid_codes.S 2>&1

test:
	go test ./... --cover -coverprofile=reports/coverage

build:
	go build ./...

coverage-report: test
	go tool cover -html=reports/coverage

benchmark:
	go test -benchmem -run=^$$ -bench . ./docs/benchmark --cpuprofile ./data/pgo/benchmark-cpu.pprof --memprofile ./data/pgo/benchmark-mem.pprof

pprof:
	go tool pprof --http=:8080 .\data\pgo\benchmark.pprof