package validator

import (
	"errors"
	"log"
	"time"
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

func DurationConvertation(duration string) time.Duration {
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln("Duration convertation err", err)
	}
	return d
}
