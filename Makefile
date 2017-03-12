.PHONY: go build test
.SILENT: run

SVC_ROOT	= `pwd`
BINARY		= malinche
PACKAGE		= github.com/raeoks/malinche

all: go

clean:
	if [ -f build/${BINARY} ] ; then rm build/${BINARY} ; fi

init:
	mkdir -p build
	glide install

setup: init

build: setup
	go build -v -o ${SVC_ROOT}/build/${BINARY} ${SVC_ROOT}/command/main.go

test-init:
	glide install

test-only:
	go test

test: test-init test-only

tests: test

run:
	go run ${SVC_ROOT}/command/main.go

go: run
