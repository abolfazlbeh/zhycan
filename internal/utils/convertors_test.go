package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertor_GenerateStructType(t *testing.T) {
	source := `package src

type User struct {
	Name string
}

func Add() *User {
	return nil
}
`

	type User2 struct {
		Name string
	}

	types := ExtractStructs(source)
	if types == nil {
		t.Errorf("Extracting struct from source code expecting Not be nil, but got %v", types)
	}
	userType := types[0]

	p := GenerateStructType(userType)
	if p == nil {
		t.Errorf("converting ast type to reflect type, but got %v", p)
	}

	intPtr := reflect.New(p)

	expected := reflect.ValueOf(User2{}).Interface()
	got := reflect.ValueOf(intPtr).Type()

	fmt.Println(expected, got)

	if got != expected {
		t.Errorf("Expecting the type %v, but got %v", expected, got)
	}
}
