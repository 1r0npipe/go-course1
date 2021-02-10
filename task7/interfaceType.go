package main

import (
	"fmt"	
	"reflect"
)

type simpleInt32 int32
type simplefloat64 float64

func (si simpleInt32) printType() string {
	return  reflect.TypeOf(si)
}

func (sf simplefloat64) printType() string {
	return string(reflect.TypeOf(sf))
}

type objectType interface {
	printType() string
}

func printType() string {
	fmt.Println(reflect.TypeOf())
}