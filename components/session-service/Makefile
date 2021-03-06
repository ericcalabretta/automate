include ../../Makefile.common_go

PACKAGE_PATH = github.com/chef/automate/components/session-service

BINS = ${PACKAGE_PATH}/cmd/session-service

VERSION = $$(git describe --always)
GO_LDFLAGS = --ldflags "-X ${PACKAGE_PATH}/version.Version=${VERSION}"

packages:=${PACKAGE_PATH}/...

all: lint build test

static: lint
unit: build test

${BINS}: bin
	@echo "GO $@"
	@cd bin; go build ${GO_LDFLAGS} $@

bin:
	mkdir -p bin

build: ${BINS}

test:
	@go test -v -count=1 $(packages) -cover

# this command lists all the changes since master, and looks for modifications
# to the migration files -- if there's any of (M)odify, (R)ename, or (D)elete,
# it will trigger an error
migrations-append-only-check:
	@git diff --name-status master... | awk '/[RMD][0-9]*\tcomponents\/session-service\/migration\/sql\//{ print "modified/deleted/renamed: " $$2; ec=1 } END{ exit ec }'

# Regenerate all *.pb.go files
proto:
	cd ../../ && hab studio run 'source .studiorc; compile_go_protobuf_component session-service'

.PHONY: all fmt test vet static unit
