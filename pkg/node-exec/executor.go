package node_exec

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecPart() {
	path, err := os.Getwd()

	cmd := exec.Command("./quark_worker", "w")
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	logrus.Println("Running Worker command and waiting for it to finish...")

	err = cmd.Run()
	logrus.Println("Worker-[%s] exit with err <==%v", stderr.String(), err.Error())
}
func ExecSsh(program string, conf string) {

	path, err := os.Getwd()

	var stderr bytes.Buffer
	cmd := exec.Command(program, conf, "/bin/bash")
	cmd.Dir = path
	cmd.Stderr = &stderr

	logrus.Println("Running Agent command and waiting for it to finish...")

	err = cmd.Run()
	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	logrus.Warnf("SSH_Agent %v:%s", err, stderr.String())
	logrus.Warnf("Command finished with error: %v", err)

}
