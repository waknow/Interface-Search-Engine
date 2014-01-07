//basic data type

package lib

type Value struct {
	Name string
	Type string
}

func (v *Value) IsEqual(value Value) bool {
	return v.Type == value.Type
}

type Values []Value

func (v *Values) IsEqual(value Values) (res bool) {
	res = len(v) == len(value)
	if !res {
		return
	}

	for index, val := range v {
		res &= val.IsEqual(val[index])
	}
	return
}
