package automap

import (
	"errors"
	"log"
	"reflect"
)

var ERROR_MAPPER_VAL_IS_NIL = errors.New("mapper val is nil")
var ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH = errors.New("mapper kind of src and dest not match")

type AutoMapper struct {
}

func (me *AutoMapper) Auto(src interface{}, dest interface{}) error {

	srckind := me.GetKind(src)
	destkind := me.GetKind(dest)
	srcval := me.ReflectValue(src)
	destval := me.ReflectValue(dest)

	if me.IsNormalType(srckind) { //normal type
		if srckind != destkind {
			return ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH
		}

		destval.Set(srcval)
		return nil
	} else if srckind == reflect.Slice { //array
		if srckind != destkind {
			return ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH
		}
		i := 0
		max := srcval.Len()
		destelemtype := reflect.TypeOf(destval)
		for i < max {

			//me.Auto(srcval.Index(i).Interface())
			log.Printf("%#v", destelemtype)
			i++
		}
	}

	return nil
}

func (me *AutoMapper) ReflectValue(obj interface{}) reflect.Value {

	var val reflect.Value
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		val = reflect.ValueOf(obj).Elem()
	} else {
		val = reflect.ValueOf(obj)
	}
	return val
}

//NormalType -> int, float64, string, etc...
func (me *AutoMapper) IsNormalType(kind reflect.Kind) bool {
	if kind != reflect.Struct && kind != reflect.Slice {
		return true
	}
	return false
}

func (me *AutoMapper) GetKind(obj interface{}) reflect.Kind {
	kind := reflect.TypeOf(obj).Kind()
	if kind == reflect.Ptr {
		kind = reflect.TypeOf(obj).Elem().Kind()
	}
	return kind
}

/*
func (me *AutoMapper) ReflectValue(obj interface{}) reflect.Value {

	var val reflect.Value
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		val = reflect.ValueOf(obj).Elem()
	} else {
		val = reflect.ValueOf(obj)
	}

	return val
}*/
