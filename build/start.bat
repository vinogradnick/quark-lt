set cgo=0
set arch=amd64
CD windows
env GOOS=windows GOARCH=%arch% CGO_ENABLED=%cgo% go build ..\..\cmd\quark-apiserver\quark-apiserver.go
env GOOS=windows GOARCH=%arch% CGO_ENABLED=%cgo%  go build ..\..\cmd\quark-node\quark-node.go
env GOOS=windows GOARCH=%arch% CGO_ENABLED=%cgo%  go build ..\..\cmd\quark-worker\quark_worker.go
CD ..
CD unix
env GOOS=linux GOARCH=%arch% CGO_ENABLED=%cgo%  go build ..\..\cmd\quark-apiserver\quark-apiserver.go
env GOOS=linux GOARCH=%arch% CGO_ENABLED=%cgo%  go build ..\..\cmd\quark-node\quark-node.go
env GOOS=linux GOARCH=%arch% CGO_ENABLED=%cgo%  go build ..\..\cmd\quark-worker\quark_worker.go
CD ..