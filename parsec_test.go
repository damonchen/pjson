package pjson

import (
	"fmt"
	"testing"
)

// Test_Parsec 测试 Parsec
func Test_Parse(t *testing.T) {
	code := `{"abc": "bcd", "damon" : true, "oo": {"inner": null, "other": false}, "list"	:
	 			["abc", "txt", "value", true, null]}`
	v, err := Parse(code)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(v)
}
