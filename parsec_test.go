package pjson

import (
	"fmt"
	"testing"
)

// Test_Parsec 测试 Parsec
func Test_Parse(t *testing.T) {
	code := `{"cc": 1.2, "bb": true}`
	fmt.Printf("test code %v\n", code)
	v, err := Parse(code)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(v)
}
