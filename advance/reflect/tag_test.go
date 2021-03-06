package reflecter

import (
	"fmt"
	"reflect"
)

// Elem() // 获取指针所指对象 v = v.Elem()

// Get Struct Tag
func Test_StructTag() {
	type S struct {
		S1 struct {
			S1F string `name:"test" color:"blue"`
		} `id:"1"`

		S2 struct {
			S2F string `name:"test" color:"green"`
		} `id:"2"`
	}

	s := S{}
	sVal := reflect.ValueOf(&s).Elem()
	sType := sVal.Type()

	for i := 0; i < sVal.NumField(); i++ {
		vf := sVal.Field(i)
		vfT := vf.Type()

		tf := sType.Field(i)
		fmt.Println("id", tf.Tag.Get("id"))

		for j := 0; j < vf.NumField(); j++ {
			fmt.Println("--color", vfT.Field(j).Tag.Get("color"))
		}
	}
}

// ValueOf
func Test_ValueOf() {
	type T struct {
		A int `name:"czx" color:"A"`
		//B string `name:"czx" color:"B"`
	}
	//t := T{12, "skidoo"}
	t := T{12}

	v := reflect.ValueOf(&t).Elem()
	vType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		canSetValue := reflect.ValueOf(13)
		vf.Set(canSetValue) // 赋值
		fmt.Println(vf, vf.Type())

		tf := vType.Field(i)
		tf.Tag = `name:"czx11" color:"B"`
		fmt.Println(tf.Name, tf.Tag.Get("name"))
		fmt.Println("--------------")
	}
}

// TypeOf
func Test_TypeOf() {
	type S struct {
		F string `name:"test" color:"blue"`
	}
	s := S{}

	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		stf := st.Field(i)
		fmt.Println(stf.Name, stf.Type, stf.Tag.Get("name"))
	}
}
