package validation

type StringField[T StringType] struct {
	FieldValidator[T]
}
type StringType interface {
	~*string
}

var _ Validatable = &StringField[*string]{}

func lengthBetweenRule[T StringType](min, max int) *rule[T] {
	var params = make(map[string]interface{})
	params["min"] = min
	params["max"] = max
	return &rule[T]{
		code: MsgKeyStringMaxLength,
		validatorFn: func(value T) (bool, error) {
			return len(*value) >= min && len(*value) <= max, nil
		},
		params: params,
	}
}

func (f *StringField[T]) LengthBetween(min, max int) *StringField[T] {
	f.rules = append(f.rules, lengthBetweenRule[T](min, max))
	return f
}
