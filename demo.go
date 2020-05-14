//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 9:37

package main

import (
	"fmt"
	"reflect"
)

type Login struct {
	User     string `form:"user" validate:"required,min=5,max=20"`
	Password string `form:"password" validate:"required,min=5,max=20"`
}


func main() {
	value := reflect.ValueOf(Login{User: "lee",Password:"123456"})

	fmt.Println(value)
	v := value

	if value.Kind() ==  reflect.Ptr {
		v =value.Elem()
	}
	fmt.Println(v.Type().Name())
	Name, _ := v.Type().FieldByName("User")
	fmt.Println(Name.Tag.Get("form"))
	for i := 0; i < v.Type().NumField(); i++ {
		typeField := v.Type().Field(i)
		valueField := v.Field(i)
		value := valueField.Interface()
		fmt.Printf("fieldName :[%s], fieldValue :[%s] tag :[%s]\n", typeField.Name, value, typeField.Tag)

		switch vv := value.(type) {
		case int:
			fmt.Printf("this is a int : %d\n", vv)
		case string:
			fmt.Println("this is a string : " + vv)
		default:
			fmt.Println("other value")
		}
	}
}
