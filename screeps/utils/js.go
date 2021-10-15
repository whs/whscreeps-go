package utils

import "syscall/js"

var object = js.Global().Get("Object")

func ObjectKeys(obj js.Value) []string {
	keys := object.Call("keys", obj)
	length := keys.Length()

	out := make([]string, length)
	for i := 0; i < length; i++ {
		out[i] = keys.Index(i).String()
	}
	return out
}

func NewObject() *js.Value {
	val := object.New()
	return &val
}
