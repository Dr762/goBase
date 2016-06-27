#=============#
# TheQuestion #
# ============#
all: get install


fmt:
	@. ./gopath.sh && go fmt ./...

get:
	@. ./gopath.sh && go get -t -v ./...

test:
	@. ./gopath.sh && go test -p 1 ./...

test-verbose:
	@. ./gopath.sh && go test -p 1 -v ./...

install:
	@. ./gopath.sh && go install ./...
