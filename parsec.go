package pjson

import (
	"fmt"
	"reflect"

	p "github.com/Dwarfartisan/goparsec2"
)

// Skip 忽略匹配指定算子的内容
var Skip = p.Skip(p.Space)

// Comma 逗号分隔符
var Comma = p.Skip(p.Str(","))

// Colon 冒号分隔符
var Colon = p.Skip(p.Str(":"))

func listBodyParser(st p.State) (interface{}, error) {
	value, err := p.SepBy1(ValueParser(), Comma)(st)
	fmt.Printf("type :%v, value :%v, err :%v\n", reflect.TypeOf(value), value, err)
	return value, err
}

// ListParser 实现数组解析器
func ListParser() p.P {
	return func(st p.State) (interface{}, error) {
		left := p.Chr('[').Then(Skip)
		right := Skip.Then(p.Chr(']'))
		empty := p.Between(left, right, Skip)

		list, err := p.Between(left, right, p.Many(p.Choice(p.NChr(']'), listBodyParser)))(st)
		fmt.Printf("list type :%v, value :%v, err: %v\n", reflect.TypeOf(list), list, err)
		if err != nil {
			_, e := empty(st)
			if e != nil {
				return nil, err
			}
			return List{}, nil
		}

		switch l := list.(type) {
		case List:
			return L(l), nil
		case []interface{}:
			return list.([]interface{}), nil
		default:
			return nil, fmt.Errorf("List Parser Error: %v type is unexpected: %v", list, reflect.TypeOf(list))
		}
	}
}

func objectBodyParser(st p.State) (interface{}, error) {
	return p.SepBy1(ValueParser(), p.Choice(Colon))(st)
}

// ObjectParser 实现数组解析器
func ObjectParser() p.P {
	return func(st p.State) (interface{}, error) {
		left := p.Chr('{').Then(Skip)
		right := Skip.Then(p.Chr('}'))
		empty := p.Between(left, right, Skip)

		object, err := p.Between(left, right, p.Many(p.Choice(objectBodyParser, p.NChr('}'))))(st)
		if err != nil {
			_, e := empty(st)
			if e != nil {
				return nil, err
			}
			return Object{}, nil
		}

		switch o := object.(type) {
		case Object:
			return O(o), nil
		case map[string]interface{}:
			return object.(map[string]interface{}), nil
		default:
			return nil, fmt.Errorf("Object Parser Error: %v type is unexpected: %v", object, reflect.TypeOf(object))
		}

	}
}

// ValueParser 实现简单的值解释器
func ValueParser() p.P {
	return func(st p.State) (interface{}, error) {
		value, err := p.Choice(p.Try(StringParser),
			p.Try(FloatParser),
			p.Try(ObjectParser()),
			p.Try(ListParser()),
			p.Try(BoolParser),
			p.Try(NullParser),
		)(st)

		return value, err
	}
}
