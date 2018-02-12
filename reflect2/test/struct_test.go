package test

import (
	"testing"
	"github.com/v2pro/plz/reflect2"
	"github.com/v2pro/plz/test"
	"github.com/v2pro/plz/countlog"
	"unsafe"
	"github.com/v2pro/plz/test/must"
)

func Test_struct(t *testing.T) {
	type TestObject struct {
		Field1 int
		Field2 int
	}
	t.Run("New", testOp(func(api reflect2.API) interface{} {
		valType := api.TypeOf(TestObject{})
		obj := valType.New()
		obj.(*TestObject).Field1 = 20
		obj.(*TestObject).Field2 = 100
		return obj
	}))
	t.Run("Set", testOp(func(api reflect2.API) interface{} {
		valType := api.TypeOf(TestObject{}).(reflect2.StructType)
		field1 := valType.FieldByName("Field1")
		obj := TestObject{}
		field1.Set(&obj, 100)
		return obj
	}))
	t.Run("UnsafeSet", test.Case(func(ctx *countlog.Context) {
		valType := reflect2.TypeOf(TestObject{}).(reflect2.StructType)
		field1 := valType.FieldByName("Field1")
		obj := TestObject{}
		field1.UnsafeSet(unsafe.Pointer(&obj), reflect2.PtrOf(100))
		must.Equal(100, obj.Field1)
	}))
	t.Run("Get", testOp(func(api reflect2.API) interface{} {
		obj := TestObject{Field1: 100}
		valType := api.TypeOf(obj).(reflect2.StructType)
		field1 := valType.FieldByName("Field1")
		return field1.Get(&obj)
	}))
	t.Run("UnsafeGet", test.Case(func(ctx *countlog.Context) {
		obj := TestObject{Field1: 100}
		valType := reflect2.TypeOf(obj).(reflect2.StructType)
		field1 := valType.FieldByName("Field1")
		value := field1.UnsafeGet(unsafe.Pointer(&obj))
		must.Equal(100, *(*int)(value))
	}))
}
