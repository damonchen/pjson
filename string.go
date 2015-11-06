package pjson

import p "github.com/Dwarfartisan/goparsec2"

// Rune 是 rune 的简单封装
type Rune rune

// EscapeChars 用于string
var EscapeChars = p.Do(func(st p.State) interface{} {
	p.Chr('\\').Exec(st)

	r := p.RuneOf("nrt\"\\").Exec(st)
	ru := r.(rune)
	switch ru {
	case 'r':
		return '\r'
	case 'n':
		return '\n'
	case '"':
		return '"'
	case '\\':
		return '\\'
	case 't':
		return '\t'
	default:
		panic(st.Trap("Unknown escape sequence \\%c", r))
	}
})

//用于rune
var EscapeCharr = p.Do(func(st p.State) interface{} {
	p.Chr('\\').Exec(st)
	r := p.RuneOf("nrt'\\").Exec(st)
	ru := r.(rune)
	switch ru {
	case 'r':
		return '\r'
	case 'n':
		return '\n'
	case '\'':
		return '\''
	case '\\':
		return '\\'
	case 't':
		return '\t'
	default:
		panic(st.Trap("Unknown escape sequence \\%c", r))
	}
})

// RuneParser 实现 rune 的解析
var RuneParser = p.Do(func(state p.State) interface{} {
	p.Chr('\'').Exec(state)
	c := p.Choice(p.Try(EscapeCharr), p.NChr('\'')).Exec(state)
	p.Chr('\'').Exec(state)
	return Rune(c.(rune))
})

// StringParser 实现字符串解析
var StringParser = p.Between(p.Chr('"'), p.Chr('"'),
	p.Many(p.Choice(p.Try(EscapeChars), p.NChr('"')))).Bind(p.ReturnString)
