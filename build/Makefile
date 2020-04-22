all: build.core build.node build.worker




build.core:
	go build ../cmd/quark-apiserver/*.go


build.node:
	go build ../cmd/quark-node/*.go
build.worker:
	go build ../cmd/quark-worker/*.go