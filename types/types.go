package types

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	ErrInvalidNumber = errors.New("parsing failed - invalid number")
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Parse[T Number](s string) (T, error) {
	var z T
	rt := reflect.TypeOf(z)
	switch rt.Kind() {
	case reflect.Float32, reflect.Float64:
		t, err := strconv.ParseFloat(s, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t, err := strconv.ParseInt(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		t, err := strconv.ParseUint(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	default:
		return z, ErrInvalidNumber
	}
}

func ToString[T Number](number T) string {
	return fmt.Sprintf("%v", number)
}
