all: build.linux



build.linux:
	go build cmd/quark-apiserver/quark-apiserver.go
	go build cmd/quark-node/*.go
	go build cmd/quark-worker/*.go
