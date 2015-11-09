package pjson

import (
	"fmt"

	p "github.com/Dwarfartisan/goparsec2"
)

// Bool 是内置的 bool 类型
type Bool bool

// BoolParser 解析 bool
var BoolParser = p.Choice(p.Str("true"), p.Str("false")).Bind(func(input interface{}) p.P {
	return func(st p.State) (interface{}, error) {
		switch input.(string) {
		case "true":
			return Bool(true), nil
		case "false":
			return Bool(false), nil
		default:
			return nil, fmt.Errorf("Unexpect bool token %v", input)
		}
	}
})

// var BoolParser = p.Do(func(st p.State) interface{} {
// 	f := p.Try(p.Bool).Exec(st)
// 	switch f.(string) {
// 	case "true":
// 		return Bool(true), nil
// 	case "false":
// 		return Bool(false), nil
// 	default:
// 		return nil, fmt.Errorf("Unexpect bool token %v", input)
// 	}
// })

// NullParser 解析 null
var NullParser = p.Str("null").Then(p.Return(nil))

// // Null 类型定义空值行为
// type Null struct {
// }
