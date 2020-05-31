package flexql

type Table struct {
	Name   string
	Fields []Field
}

func NewTable(name string) *Table {
	return &Table{
		Name: name,
	}
}

func (t *Table) AddField(field Field) {
	field.alias = t.Name
	t.Fields = append(t.Fields, field)
}

func (t *Table) Get(fieldName string) *Field {
	for _, f := range t.Fields {
		if f.Name == fieldName {
			return &f
		}
	}
	return nil
}
