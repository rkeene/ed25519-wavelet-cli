GOFLAGS      := -mod=vendor
GO111MODULE  := on
GOPATH       := $(shell pwd)/.go
export GOFLAGS G111MODULE GOPATH

PATH         := /opt/appfs/rkeene.org/go/platform/latest/bin:${PATH}
export PATH

all: verify

verify: verify.go vendor
	go build -o verify .

vendor: go.mod go.sum
	go mod vendor

clean:
	rm -f verify

distclean: clean
	chmod -R u+w .go
	rm -rf .go
	rm -rf vendor

.PHONY: all clean distclean
