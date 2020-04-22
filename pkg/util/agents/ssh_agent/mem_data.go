package ssh_agent

import (
	"bytes"
	"github.com/vinogradnick/quark-lt/pkg/metrics"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
)

func ParseMemory(data string) *metrics.MemoryInfo {

	parsor := strings.Split(data, " ")
	//for i := 0; i < len(parsor); i++ {
	//	fmt.Println("[%d]-", i, parsor[i])
	//}

	return &metrics.MemoryInfo{
		Total:  parse(parsor[48]),
		Used:   parse(parsor[53]),
		Free:   parse(parsor[58]),
		Caches: parse(parsor[64]),
	}

}
func parse(data string) uint32 {

	arrData, err := strconv.ParseUint(data, 10, 32)
	if err != nil {

		return 0
	}
	return uint32(arrData)
}
func GetMemoryInfo(session *ssh.Session, output *bytes.Buffer) (*metrics.MemoryInfo, float64) {
	err := session.Run(`free && echo ';' && grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage ""}'  `)
	if err != nil {
		panic(err)
	}
	data := strings.Split(output.String(), ";")

	return ParseMemory(data[0]), GetCpuAvInfo(data[1])
}
