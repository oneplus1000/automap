package automap

import (
	"errors"
	//"log"
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

	if srckind == reflect.Slice {

		if srckind != destkind {
			return ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH
		}
		i := 0
		max := srcval.Len()
		destTypeOfItemInSlice := destval.Type().Elem()
		newarr := reflect.MakeSlice(destval.Type(), max, max)
		for i < max {
			destValOfItemInSlice := reflect.New(destTypeOfItemInSlice)
			err := me.Auto(srcval.Index(i).Interface(), destValOfItemInSlice.Interface())
			if err != nil {
				return err
			}
			newarr.Index(i).Set(destValOfItemInSlice.Elem())
			i++
		}
		destval.Set(newarr)
		return nil

	} else if srckind == reflect.Struct {

		if srckind != destkind {
			return ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH
		}

		i := 0
		srctype := srcval.Type()
		desttype := destval.Type()
		srcfieldsize := srcval.NumField()
		for i < srcfieldsize {
			srcfield := srctype.Field(i)
			if srcfield.PkgPath != "" { //ไม่ใช้ export field ข้ามไป( PkgPath is empty for exported fields )
				i++
				continue
			}

			name := srcfield.Name
			_, ok := desttype.FieldByName(name)
			if !ok { //dest ไม่มี field นี้ ข้ามไป
				i++
				continue
			}

			srcfieldval := srcval.FieldByName(name)
			destfieldval := destval.FieldByName(name)

			newval := reflect.New(destfieldval.Type())
			me.Auto(srcfieldval.Interface(), newval.Interface())
			destfieldval.Set(newval.Elem())
			i++
		}
		return nil

	} else {

		if srckind != destkind {
			return ERROR_MAPPER_KIND_OF_SRC_AND_DEST_NOT_MATCH
		}
		//log.Printf("--------> srckind=%s destkind=%s\n\n", srckind, destkind)
		destval.Set(srcval)
		return nil

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
