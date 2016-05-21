// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"fmt"
	"reflect"
)

// -----------------------------------------------------------------------------

// Type represents the underlying type of an interface.
type Type int

const (
	// TypeInvalid is an illegal type.
	TypeInvalid Type = iota
	// TypeNumeric is any of bool/int/uint/float/complex and their variants.
	TypeNumeric Type = iota
	// TypeArray is an array of any underlying type.
	TypeArray Type = iota
	// TypeStruct is any struct.
	TypeStruct Type = iota
	// TypeUnsafePointer is any pointer from the unsafe package.
	TypeUnsafePointer Type = iota
)

// TypeOf returns the underlying type of an interface.
func TypeOf(v interface{}) (Type, error) {
	return typeOf(reflect.ValueOf(v))
}

func typeOf(v reflect.Value) (Type, error) {
	if !v.IsValid() {
		return TypeInvalid, Error(fmt.Sprintf("unsuppported type: %#v", v))
	}

	k := v.Type().Kind()

	switch k {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32,
		reflect.Float64, reflect.Complex64, reflect.Complex128:
		return TypeNumeric, nil
	case reflect.Array:
		return TypeArray, nil
	case reflect.Struct:
		return TypeStruct, nil
	case reflect.UnsafePointer:
		return TypeUnsafePointer, nil
	default:
		return TypeInvalid, Error(fmt.Sprintf("unsuppported type: %#v", k.String()))
	}
}
