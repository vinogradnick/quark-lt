package ssh_agent

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	Total  uint32
	Used   uint32
	Free   uint32
	Caches uint32
}

func ParseMemory(data string) *MemoryInfo {

	parsor := strings.Split(data, " ")

	return &MemoryInfo{
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
func GetMemoryInfo(session *ssh.Session, output *bytes.Buffer) *MemoryInfo {
	err = session.Run("free")
	if err != nil {
		panic(err)
	}
	data := output.String()
	return ParseMemory(data)
}