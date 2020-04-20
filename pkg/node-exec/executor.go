package node_exec

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecPart(program string, args string) *exec.Cmd {
	path, err := os.Getwd()
	fmt.Println(path)
	cmd := exec.Command(program, args)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	logrus.Println("Running Worker command and waiting for it to finish...")

	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(cmd.Process.Pid)

	logrus.Warnf("Worker-[%s] exit with err <==%v", stderr.String(), err)
	return cmd
}
