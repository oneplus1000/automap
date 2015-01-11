package basemapper

import (
	"errors"
	"github.com/oneplus1000/reflections"
	"log"
	"reflect"
)

var ERROR_MAPPER_VAL_IS_NIL = errors.New("mapper val is nil")

type BaseMapper struct {
}

func (me *BaseMapper) Auto(src interface{}, dest interface{}) error {

	var err error
	var fieldnames []string
	fieldnames, err = reflections.Fields(dest)
	if err != nil {
		return err
	}

	for _, fieldname := range fieldnames {
		if has, _ := reflections.HasField(src, fieldname); !has {
			continue
		}

		destkind, err := reflections.GetFieldKind(dest, fieldname)
		if err != nil {
			return err
		}
		srckind, err := reflections.GetFieldKind(src, fieldname)
		if err != nil {
			return err
		}

		if destkind != srckind {
			continue
		}

		val, err := reflections.GetField(src, fieldname)
		if err != nil {
			return err
		}

		if destkind == reflect.Slice {

			destfield, err := reflections.GetField(dest, fieldname)
			if err != nil {
				return err
			}
			typOfDestfield := reflect.TypeOf(destfield)
			obj := reflect.New(typOfDestfield.Elem())

			slice := reflect.ValueOf(val)
			i := 0
			max := slice.Len()
			for i < max {
				var c int
				inf := obj.Interface()
				_ = inf
				me.Auto(3, c)
				log.Printf("%#v  %#v", c, slice.Index(i).Interface())
				i++
			}
			continue
		}

		err = reflections.SetField(dest, fieldname, val)
		if err != nil {
			//Debug("fieldname %s srckind = %s", fieldname, srckind)
			return err
		}
	}

	return err
}
