package main

import (
"fmt"
"reflect"
)

func main() {
	type User struct {
		Name string
		Age  int
	}

	re := reflect.New(reflect.TypeOf(User{})).Elem()

	rt := re.Type()
	p(re, rt)
	for i := 0; i < rt.NumField(); i++ {

		f := rt.Field(i)
		p(f.Name)
		p(f.Type)
		p(f.Tag)
	}

	if f, ok := rt.FieldByName("Name"); ok {
		p(f.Name, f.Type) //=> Name string
	}

	re.Field(0).SetString("user1")
	p(re.Field(0))
	
	// 関数の実行
	fn := func(s string, i int) string {
		out := ""
		for j := 0; j < i; j++ {
			out += s
		}
		return out
	}

	fnt := reflect.TypeOf(fn)
	for i := 0; i < fnt.NumIn(); i++ {
		// 引数の型の取得
		p(fnt.In(i))
	}

	for i := 0; i < fnt.NumOut(); i++ {
		// 返り値の型の取得
		p(fnt.Out(i))
	}

	fnv := reflect.ValueOf(fn)
	out := fnv.Call([]reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf(2)})
	if s, ok := out[0].Interface().(string); ok {
		p(s)
	}
}

func p(a ...interface{}) {
	fmt.Println(a...)
}
