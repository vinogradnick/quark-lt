package validator

import (
	"errors"
	"strconv"
	"strings"
)

type Validator interface {
	Validate() bool
}

func PositiveValidate(value int32) error {
	if value <= 0 {
		return errors.New("value aren't postive")
	}
	return nil
}

func DurationConvertation(duration string) int64 {
	var str string
	var data int64
	var err error
	if strings.Contains(duration, "min") {
		str = duration[0:strings.Index(duration, "m")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data * 60000
	}

	if strings.Contains(duration, "sec") {
		str = duration[0:strings.Index(duration, "s")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data * 1000
	}
	if strings.Contains(duration, "ms") {
		str = duration[0:strings.Index(duration, "m")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data
	}
	return 0
}
