package rest

import "reflect"

// PayloadMember returns the payload field member of i if there is one, or nil.
func PayloadMember(i interface{}) interface{} {
	field := PayloadFieldName(i)
	if field == "" {
		return nil
	}

	iVal := reflect.ValueOf(i)
	if member := iVal.FieldByName(field); member.IsValid() {
		return member.Interface()
	}

	return nil
}

func PayloadFieldName(i interface{}) string {
	if i == nil {
		return ""
	}

	field := ""
	if t, ok := i.(interface {
		PayloadField() string
	}); ok {
		field = t.PayloadField()
	}

	return field
}

// PayloadType returns the type of a payload field member of i if there is one, or "".
func PayloadType(i interface{}) string {
	field := PayloadFieldName(i)
	if field == "" {
		return ""
	}

	iType := reflect.TypeOf(i).Elem()
	if member, ok := iType.FieldByName(field); ok {
		return member.Tag.Get("type")
	}

	return ""
}
