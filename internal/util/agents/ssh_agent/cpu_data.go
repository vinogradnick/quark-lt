package ssh_agent

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"strconv"
)

func GetCpuAvInfo(session *ssh.Session, output *bytes.Buffer) float64 {
	err = session.Run(`grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage ""}'`)
	if err != nil {
		panic(err)
	}
	data := output.String()
	val, err := strconv.ParseFloat(data, 64)
	if err != nil {
		panic(err)
	}
	return val
}
