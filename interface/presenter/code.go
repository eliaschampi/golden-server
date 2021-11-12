package presenter

import "github.com/google/uuid"

//CODE entity ID
type CODE = uuid.UUID

//NewID create a new entity ID
func NewID() CODE {
	return CODE(uuid.New())
}

//StringToCode convert a string to an entity ID
func StringToCode(s string) (CODE, error) {
	code, err := uuid.Parse(s)
	return CODE(code), err
}

func CodeToString(code CODE) string {
	return CODE.String(code)
}


