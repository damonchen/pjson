package pjson

import (
	"fmt"
	"reflect"

	p "github.com/Dwarfartisan/goparsec2"
)

// Skip 忽略匹配指定算子的内容
var Skip = p.Skip(p.Space)

// Comma 逗号分隔符
var Comma = p.Str(",")

// Colon 冒号分隔符
var Colon = p.Str(":")

func listBodyParser(st p.State) (interface{}, error) {
	value, err := p.SepBy(ValueParser(), Skip.Then(Comma).Then(Skip))(st)
	return value, err
}

// ListParser 实现数组解析器
func ListParser() p.P {
	return func(st p.State) (interface{}, error) {
		left := p.Chr('[').Then(Skip)
		right := Skip.Then(p.Chr(']'))
		empty := p.Between(left, right, Skip)

		//		list, err := p.Between(left, right, p.UnionAll(listBodyParser))(st)
		list, err := p.Between(left, right, listBodyParser)(st)

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

// key只能是string，value可以为任意值，中间以冒号分隔
func objectKeyValueParser() p.P {
	return func(st p.State) (interface{}, error) {
		pair := Pair{}
		s, err := Skip.Then(StringParser)(st)
		if err != nil {
			return nil, err
		}

		if str, ok := s.(string); ok {
			pair.Key = str
		} else {
			return nil, fmt.Errorf("%v is not a string", s)
		}

		_, err = Skip.Then(Colon).Then(Skip)(st)
		if err != nil {
			return nil, err
		}

		v, err := ValueParser()(st)
		if err != nil {
			return nil, err
		}

		pair.Value = v

		return pair, nil
	}
}

// key,value通过逗号进行分隔，类似list处理
func objectBodyParser(st p.State) (interface{}, error) {
	value, err := p.SepBy(objectKeyValueParser(), Skip.Then(Comma).Then(Skip))(st)
	if err != nil {
		return nil, err
	}

	o := Object{}
	if vlist, ok := value.([]interface{}); ok {
		for _, p := range vlist {
			if pair, ok := p.(Pair); ok {
				o[pair.Key] = pair.Value
			}
		}
	}

	return o, nil
}

// ObjectParser 实现数组解析器
func ObjectParser() p.P {
	return func(st p.State) (interface{}, error) {
		left := p.Chr('{').Then(Skip)
		right := Skip.Then(p.Chr('}'))
		empty := p.Between(left, right, Skip)

		object, err := p.Between(left, right, objectBodyParser)(st)
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
