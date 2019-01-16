package pungi

import "reflect"

const (
	pungi    = "pungi"
	nullable = "nullable"
)

// IsValid returns nil if v doesn't contains zero values (except "nullable" fields); ErrInvalid otherwise.
func IsValid(v interface{}) error {
	var name string
	if typ := reflect.TypeOf(v); typ.Kind() == reflect.Ptr {
		name = typ.Elem().Name()
	} else {
		name = typ.Name()
	}
	value := reflect.ValueOf(v)
	return isValidValue(name, value, false)
}

// isValidValue returns true if the specified value is valid or optional.
// If the value is a pointer or structure, then recursively checks nested values.
func isValidValue(name string, value reflect.Value, nullable bool) error {
	valid := true

	switch value.Kind() {
	case reflect.Bool:
		valid = value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = value.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = value.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = value.Float() != 0
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		valid = value.Len() != 0
	case reflect.Invalid:
		valid = false
	case reflect.Interface, reflect.Ptr:
		if value.IsNil() {
			valid = false
		} else {
			return isValidValue(name, value.Elem(), nullable)
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			if err := isValidValue(name+"."+field.Name, value.Field(i), isStructFieldNullable(field)); err != nil {
				return err
			}
		}
	}

	if !valid && !nullable {
		return &ErrInvalid{Name: name}
	}

	return nil
}

// isStructFieldNullable returns true if the structure field is nullable.
func isStructFieldNullable(field reflect.StructField) bool {
	value, ok := field.Tag.Lookup(pungi)
	return ok && value == nullable
}
