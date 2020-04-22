cgo=0
arch=amd64
CD windows
GOOS=windows GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-apiserver/quark-apiserver.go
GOOS=windows GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-node/quark-node.go
GOOS=windows GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-worker/quark_worker.go
cd ..
# shellcheck disable=SC2164
cd unix
GOOS=linux GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-apiserver/quark-apiserver.go
GOOS=linux GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-node/quark-node.go
GOOS=linux GOARCH=$arch CGO_ENABLED=$cgo go build ../../cmd/quark-worker/quark_worker.go
CD ..

tar -cvf apiserver.tar.gz ./unix/quark-apiserver ./config/core.yaml ./client
tar -cvf node.tar.gz ./unix/quark-node ./unix/quark_worker ./config/node.yaml
tar -cvf node.tar.gz ./unix/quark_worker ./config/worker.yaml
