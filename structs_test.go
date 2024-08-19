package structs

import (
	"testing"
)

type Original struct {
	Field1 int
	Field2 string
}

type DifferentFieldTypes struct {
	Field1 string
	Field2 int
}

type DifferentFieldNames struct {
	Field3 int
	Field4 string
}

type MoreFields struct {
	Field1 int
	Field2 string
	Field3 string
}

type LessFields struct {
	Field1 int
}

func TestEquals(t *testing.T) {

}
