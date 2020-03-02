package ssh_agent

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCpuAvInfo(data string) float64 {

	val, err := strconv.ParseFloat(strings.Trim(data, "\n"), 64)
	if err != nil {
		fmt.Println("val")
		panic(err)
	}
	return val
}
