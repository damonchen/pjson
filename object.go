package pjson

// Object 实现Object的基本类型
type Object map[string]interface{}

// O 构造一个Object对象
func O(data map[string]interface{}) Object {
	o := make(Object)
	for key, value := range data {
		o[key] = value
	}
	return o
}
