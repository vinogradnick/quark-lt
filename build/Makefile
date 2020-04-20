all: build.apiserver build.node build.worker build.tube





build.apiserver:
	go build ../cmd/quark-apiserver/*.go


build.node:
	go build ../cmd/quark-node/*.go
build.worker:
	go build ../cmd/quark-worker/*.go

build.tube:
	go build ../cmd/quark-tube/quark-tube.go