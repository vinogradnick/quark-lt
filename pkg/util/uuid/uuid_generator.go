package uuid

import flex "github.com/satori/go.uuid"

func GenerateUuid() string {
	u2, err := flex.NewV4()
	if err != nil {
		panic(err)
	}
	return u2.String()
}
