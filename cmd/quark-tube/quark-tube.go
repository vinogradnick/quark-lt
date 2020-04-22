package main

import (
	"fmt"
	node_exec "github.com/vinogradnick/quark-lt/pkg/node-exec"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args[0])
	node_exec.ExecPart("./quark_worker", "-db "+args[0])
}
