SOURCES=./
BINS=bin

dep:
	dep init

.PHONY: clean
clean:
	rm -rf ${BINS}/logger

simple:
	go run ${SOURCES}/examples/simple/main.go

gcp:
	go run ${SOURCES}/examples/gcp/main.go

gcp_simple:
	go run ${SOURCES}/examples/gcp_simple/main.go

.DEFAULT_GOAL := test
test: simple gcp gcp_simple

