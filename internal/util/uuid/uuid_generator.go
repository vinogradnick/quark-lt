package uuid

import uuid "github.com/satori/go.uuid"

func GenerateUuid() string {
	u2, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u2
}
