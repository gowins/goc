DEFAULT_EXCEPT_PKGS := e2e
VERSION:=v0.1.3

all:
	go install ./...

test:
	go test -cover -p 1 `go list ./... | grep -v -E ${DEFAULT_EXCEPT_PKGS}`

fmt:
	go fmt ./...

docker:
	docker build --rm -t "reg.weipaitang.com/micro/all-in-one:go1.14-${VERSION}" .

bm:
	goc build -v -mod '' .

build:
	 go build -ldflags "-X 'github.com/qiniu/goc/cmd.center=http://127.0.0.1:7777'" .

govet-check:
	go vet ./...

statik:
	go get -u -v github.com/rakyll/statik
	go install -v github.com/rakyll/statik

gen:
	statik -src=./templates

clean:
	find tests/ -type f -name '*.bak' -delete 
	find tests/ -type f -name '*.cov' -delete 
	find tests/ -type f -name 'simple-project' -delete 
	find tests/ -type f -name '*_profile_listen_addr' -delete 
	find tests/ -type f -name 'simple_gopath_project' -delete 
	