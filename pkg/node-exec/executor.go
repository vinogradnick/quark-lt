package node_exec

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecPart(program string) *exec.Cmd {
	path, err := os.Getwd()

	cmd := exec.Command(program)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	logrus.Println("Running Worker command and waiting for it to finish...")

	err = cmd.Run()
	fmt.Println(cmd.Process.Pid)

	logrus.Warnf("Worker-[%s] exit with err <==%v", stderr.String(), err)
	return cmd
}
