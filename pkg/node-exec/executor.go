package node_exec

import (
	"log"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecPart(program string, args string) {

	cmd := exec.Command(program, args)

	logrus.Println("Running Worker command and waiting for it to finish...")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Print("Command finished with error: %v", err)

}
