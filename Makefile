GO ?= CGO_ENABLED=0 go

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
		staticcheck.json \
		test.log

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: tidy
tidy:
	test -z $(gofmt -l .)
	$(GO) vet
	which staticcheck || $(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -version
	staticcheck

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
gl-code-quality-report.json: staticcheck.json
	which golint-convert || $(GO) install github.com/banyansecurity/golint-convert@latest
	golint-convert > $@

staticcheck.json:
	-staticcheck -f json > $@

# Gitlab dependency report
govulncheck.sarif:
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version
	govulncheck -format=sarif ./... > $@

.PHONY: bench-file
bench-file:
	@echo "Running benchmarks..."
	@TMPFILE=$$(mktemp); \
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem > "$$TMPFILE"; \
	CPU_NAME=$$(grep '^cpu:' "$$TMPFILE" | head -1 | sed 's/^cpu: //' | sed 's/ CPU.*//' | sed 's/[()@]//g' | sed 's/ /_/g' | sed 's/__*/_/g' | sed 's/_$$//' ); \
	BENCH_FILE="benches/$$(go env GOOS)-$$(go env GOARCH)-$$CPU_NAME.txt"; \
	echo "Saving to $$BENCH_FILE..."; \
	cp "$$TMPFILE" "$$BENCH_FILE"; \
	rm -f "$$TMPFILE"; \
	cat "$$BENCH_FILE"

README.html: README.adoc
	asciidoc $^

.PHONY: total
total: bench-file
	@BENCH_FILE=$$(ls -t benches/*.txt | head -1); \
	echo "Calculating total from $$BENCH_FILE..."; \
	awk -f total.awk < "$$BENCH_FILE"

