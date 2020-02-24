package config

import "errors"

type Validator interface {
	validate() bool
}
func PositiveValidate(value int32) error {
	if value <= 0 {
		return errors.New("value aren't postive")
	}
	return nil
}