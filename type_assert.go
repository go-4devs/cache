package cache

import (
	"reflect"
	"time"
)

// TypeAssert assert source to target.
func TypeAssert(source, target interface{}) (err error) {
	if source == nil {
		return nil
	}

	if directTypeAssert(source, target) {
		return nil
	}

	v := reflect.ValueOf(target)
	if !v.IsValid() {
		return ErrTargetNil
	}

	if v.Kind() != reflect.Ptr {
		return NewErrorTarget(target)
	}

	v = v.Elem()

	if !v.IsValid() {
		return NewErrorTarget(target)
	}

	s := reflect.ValueOf(source)
	if !s.IsValid() {
		return ErrSourceNotValid
	}

	s = deReference(s)
	v.Set(s)

	return nil
}

// nolint: cyclop,gocyclo
func directTypeAssert(source, target interface{}) bool {
	var ok bool
	switch v := target.(type) {
	case *string:
		*v, ok = source.(string)
	case *[]byte:
		*v, ok = source.([]byte)
	case *int:
		*v, ok = source.(int)
	case *int8:
		*v, ok = source.(int8)
	case *int16:
		*v, ok = source.(int16)
	case *int32:
		*v, ok = source.(int32)
	case *int64:
		*v, ok = source.(int64)
	case *uint:
		*v, ok = source.(uint)
	case *uint8:
		*v, ok = source.(uint8)
	case *uint16:
		*v, ok = source.(uint16)
	case *uint32:
		*v, ok = source.(uint32)
	case *uint64:
		*v, ok = source.(uint64)
	case *bool:
		*v, ok = source.(bool)
	case *float32:
		*v, ok = source.(float32)
	case *float64:
		*v, ok = source.(float64)
	case *time.Duration:
		*v, ok = source.(time.Duration)
	case *time.Time:
		*v, ok = source.(time.Time)
	case *[]string:
		*v, ok = source.([]string)
	case *map[string]string:
		*v, ok = source.(map[string]string)
	case *map[string]interface{}:
		*v, ok = source.(map[string]interface{})
	}

	return ok
}

func deReference(v reflect.Value) reflect.Value {
	if (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && !v.IsNil() {
		return v.Elem()
	}

	return v
}
