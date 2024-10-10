package domain

type validator struct {
	Errors map[string]string
}

func NewValidator() *validator {
	return &validator{Errors: make(map[string]string)}
}

func (v *validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *validator) addError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *validator) check(ok bool, key, message string) {
	if !ok {
		v.addError(key, message)
	}
}
