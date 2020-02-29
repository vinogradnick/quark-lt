package ssh_agent

import (
	"bytes"
	"github.com/quark_lt/pkg/metrics"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
)

func ParseMemory(data string) *metrics.MemoryInfo {

	parsor := strings.Split(data, " ")

	return &metrics.MemoryInfo{
		Total:  parse(parsor[37]),
		Used:   parse(parsor[42]),
		Free:   parse(parsor[47]),
		Caches: parse(parsor[54]),
	}

}
func parse(data string) uint32 {
	arrData, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(arrData)
}
func GetMemoryInfo(session *ssh.Session, output *bytes.Buffer) *metrics.MemoryInfo {
	err := session.Run("free")
	if err != nil {
		panic(err)
	}
	data := output.String()
	return ParseMemory(data)
}
