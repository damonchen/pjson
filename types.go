package pjson

import "reflect"

var (
	// BOOL 类型
	BOOL = reflect.TypeOf((*bool)(nil)).Elem()
	// STRING 字符串类型
	STRING = reflect.TypeOf((*string)(nil)).Elem()
	// FLOAT 浮点数
	FLOAT = reflect.TypeOf((*float64)(nil)).Elem()
	// LIST 类型
	LIST = reflect.TypeOf((*List)(nil)).Elem()
	// OBJECT 类型
	OBJECT = reflect.TypeOf((*Object)(nil)).Elem()
)
