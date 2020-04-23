package main

import (
	node_exec "github.com/vinogradnick/quark-lt/pkg/node-exec"
)

func main() {

	node_exec.ExecPart("./quark_worker", "w")
}
