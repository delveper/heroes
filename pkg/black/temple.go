package black

type Queer interface {
	Name() (string, error)
	Fields() ([]StructValue, error)
}

type Type struct{}

func (t *Type) Name() (string, error) {
	return GetStructName(t)
}

func (t *Type) Fields() ([]StructValue, error) {
	return GetStructFieldValues(t)
}
