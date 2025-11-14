GO ?= CGO_ENABLED=0 go
BENCH_FILE ?= benches/$(shell go env GOOS)-$(shell go env GOARCH)-$(shell lscpu | grep "Model name:" | cut -d: -f2 | xargs | sed 's/ \(CPU\|@\|w\/\).*//' | sed 's/ /_/g').txt

.PHONY: all
all: tidy test

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	-rm \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		govulncheck.sarif \
		junit.xml \
		README.html \
		golangci-lint.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: tidy
tidy:
	test -z $(gofmt -l .)
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest --version
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

cpu.profile:
	$(GO) test -run=^$ -bench=Day10Part1$ -benchmem -memprofile mem.profile -cpuprofile $@

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

# Gitlab test report
junit.xml: test.log
	which go-junit-report || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	go-junit-report -version
	go-junit-report < $< > $@

# Gitlab coverage report
coverage.xml: coverage.txt
	which gocover-cobertura || $(GO) install github.com/boumenot/gocover-cobertura@latest
	gocover-cobertura < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: golangci-lint.json
	which golint-convert || $(GO) install github.com/banyansecurity/golint-convert@latest
	golint-convert < $< > $@

golangci-lint.json:
	-$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --out-format json > $@

# Gitlab dependency report
govulncheck.sarif:
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version
	govulncheck -format=sarif ./... > $@

$(BENCH_FILE):
	@echo "Running benchmarks and saving to $@..."
	@$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@

README.html: README.adoc
	asciidoc $^

.PHONY: total
total: $(BENCH_FILE)
	awk -f total.awk < $(BENCH_FILE)

