GO ?= CGO_ENABLED=0 go

.PHONY: all
all: lint test

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	rm \
		coverage.txt \
		coverge.xml \
		gl-code-quality-report \
		govulncheck.sarif \
		junit.xml \
		staticcheck.json \

.PHONY: setup
setup:
	$(GO) get github.com/boumenot/gocover-cobertura

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: tidy
tidy:
	test -z $(gofmt -l .)
	$(GO) vet
	staticcheck || $(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -version
	staticcheck

.PHONY: prof
prof:
	$(GO) test -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	$(GO) pprof cpu.profile

.PHONY: test
test:
	$(GO) test -run=Day -coverprofile=coverage.txt -covermode count gitlab.com/jhinrichsen/adventofcode2023
	$(GO) run github.com/boumenot/gocover-cobertura < coverage.txt > coverage.xml

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml

# Gitlab coverage report
coverage.xml:
	# which gocover-cobertura 2>/dev/null || $(GO) install github.com/boumenot/gocover-cobertura
	which gocover-cobertura || $(GO) install github.com/boumenot/gocover-cobertura
	gocover-cobertura < $< > $@

coverage.txt:
	$(GO) test -short -coverprofile=$@ -covermode count

# Gitlab code quality report
gl-code-quality-report.json: staticcheck.json
	# which golint-convert 2>/dev/null || $(GO) get github.com/banyansecurity/golint-convert
	which golint-convert || $(GO) get github.com/banyansecurity/golint-convert
	$(GO) run github.com/banyansecurity/golint-convert > $@

staticcheck.json:
	-staticcheck -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	# which govulncheck 2>/dev/null || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version
	govulncheck -format=sarif ./... > $@

# Gitlab test report
junit.xml:
	# which go-junit-report 2>/dev/null || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	which go-junit-report || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	go-junit-report -version
	# $(GO) test -short -v 2>&1 ./... | go-junit-report -set-exit-code > $@
	$(GO) test -short -v ./... | go-junit-report -set-exit-code > $@
