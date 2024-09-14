package main

import (
	"fmt"
	publicstruct "lecture1/22_private_demo"
	"reflect"
	"unsafe"
)

func main() {
	p := new(publicstruct.PublicStruct)
	fmt.Println(p)

	setPrivateField(p, "privateField", "newData")
	fmt.Println(p)
}

func setPrivateField(
	s *publicstruct.PublicStruct,
	fieldName string,
	value any,
) {
	objElem := reflect.ValueOf(s).Elem()
	field := objElem.FieldByName(fieldName)

	p := unsafe.Pointer(field.UnsafeAddr())
	internalField := reflect.NewAt(field.Type(), p)

	internalField.Elem().Set(reflect.ValueOf(value))
}
