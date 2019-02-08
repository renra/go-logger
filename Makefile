SOURCES=./
BINS=bin

dep:
	dep init

.PHONY: clean
clean:
	rm -rf ${BINS}/logger

dev:
	go run ${SOURCES}/examples/main.go

build: clean
	CGO_ENABLED=0 GOOS=linux packr2 build -a -installsuffix cgo -o ${BINS}/logger ${SOURCES}/examples/main.go

run:
	${BINS}/logger

build_and_run: build run

test:
	@echo "Testing"

