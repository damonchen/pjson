package pjson

import (
	"fmt"
	"strings"
)

// List 实现基本的 List 类型
type List []interface{}

func (l List) String() string {
	frags := []string{}
	for _, item := range l {
		frags = append(frags, fmt.Sprintf("%v", item))
	}

	body := strings.Join(frags, " ")
	return fmt.Sprintf("(%s)", body)
}

// L 构造一个List
func L(data ...interface{}) List {
	l := make(List, len(data))
	for idx, item := range data {
		l[idx] = item
	}
	return l
}
