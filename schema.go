package validation


type ValidatableFunc func() (Error, error)

type Validatable interface {
	Validate() (Error, error)
}

func (v ValidatableFunc) Validate() (Error, error) {
	return v()
}


type FieldValidator[T any] struct {
	rules      []ValidationRule[T]
	fieldName  string
	fieldLabel string
	value      T
}

func (f *FieldValidator[T]) Validate() (Error, error) {
	var fe *fieldError
	for _, rule := range f.rules {
		valid, err := rule.IsValid(f.value)
		if err != nil {
			return nil, err
		}

		if !valid {
			if fe == nil {
				fe = newFieldError(f.fieldName, f.value, rule.MsgKey(), rule.Params())
			} else {
				fe.violations = append(fe.violations, Violation{MsgKey: rule.MsgKey(), Params: rule.Params()})
			}
		}
	}
	return fe, nil
}
