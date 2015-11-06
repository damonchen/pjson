package pjson

import p "github.com/Dwarfartisan/goparsec2"

// Parse 分析JSON源码
func Parse(code string) (interface{}, error) {
	st := p.BasicStateFromText(code)
	var v interface{}
	var e error

	for {
		Skip(&st)
		_, err := p.Try(p.EOF)(&st)
		if err == nil {
			break
		}

		value, err := ValueParser()(&st)
		if err != nil {
			return nil, err
		}

		switch vv := value.(type) {
		default:
			v = vv
			e = nil
		}
	}
	return v, e
}
