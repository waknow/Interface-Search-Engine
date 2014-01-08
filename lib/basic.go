package lib

type Value struct {
	Name string
	Type string
}

func (v *Value) IsEqual(value Value) bool {
	return v.Type == value.Type
}

type Values []Value

func (v *Values) IsEqual(values Values) (equal bool) {
	equal = len(*v) == len(values)
	if !equal {
		return
	}

	for index, val := range *v {
		equal = equal && val.IsEqual(values[index])
	}
	return
}
