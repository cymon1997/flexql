package flexql

type Result struct {
	raw map[string]interface{}
}

func NewResult(raw map[string]interface{}) *Result {
	return &Result{
		raw: raw,
	}
}

func (r *Result) Get(field *Field) interface{} {
	switch field.Type {
	case TypeInt:
		return *(r.raw[field.As()].(*int))
	default:
		return *(r.raw[field.As()].(*string))
	}
}
