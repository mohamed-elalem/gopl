package json

import (
	"bytes"
	"fmt"
	"reflect"
)

func Encode(buf *bytes.Buffer, v interface{}) error {
	return encode(buf, reflect.ValueOf(v), 0)
}

func encode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, `"%f"`, v.Complex())
	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Array, reflect.Slice:
		if v.Kind() == reflect.Slice && v.IsNil() {
			return nil
		}
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			encode(buf, v.Index(i), depth+1)
		}
		buf.WriteByte(']')
	case reflect.Map:
		if v.IsNil() {
			return nil
		}
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			encode(buf, key, depth+1)
			buf.WriteByte(':')
			encode(buf, v.MapIndex(key), depth+1)
		}
		buf.WriteByte('}')
	case reflect.Struct:
		if depth > 1 {
			buf.WriteByte('{')
			fmt.Fprintf(buf, "%q", v.Type().Name())
			buf.WriteByte(':')
		}

		buf.WriteByte('{')
		curIndex := 0
		for i := 0; i < v.NumField(); i++ {
			if !isNil(v.Field(i)) {
				if curIndex > 0 {
					buf.WriteByte(',')
				}
				fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
				encode(buf, v.Field(i), depth+1)
				curIndex++
			}
		}
		buf.WriteByte('}')
		if depth > 0 {
			buf.WriteByte('}')
		}
	case reflect.Ptr:
		if !isNil(v) {
			encode(buf, v.Elem(), depth)
		}
	}

	return nil
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Interface:
		return v.IsNil()
	}
	return false
}
