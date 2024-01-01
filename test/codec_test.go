package test

import (
	"bytes"
	"encoding/gob"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestGob(t *testing.T) {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(Person{Name: "yyh", Age: 15})
	if err != nil {
		t.Fatal(err)
	}

	dec := gob.NewDecoder(&buffer)
	var decodePerson Person
	err = dec.Decode(&decodePerson)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decodePerson)
}
