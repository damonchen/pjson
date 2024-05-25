package pjson

import (
	p "github.com/Dwarfartisan/goparsec2"
)

// Parse 分析JSON源码
func Parse(code string) (interface{}, error) {
	st := p.BasicStateFromText(code)
	// yes, i understand the monad
	parser := p.Try(Skip).Then(ValueParser())
	return parser.Parse(&st)
}
