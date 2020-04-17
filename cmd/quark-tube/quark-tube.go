package main

import (
	node_exec "github.com/quark_lt/pkg/node-exec"
	"os"
)

func main() {
	args := os.Args[1:]
	node_exec.ExecPart("./quark_worker", "-db "+args[0])
}
