package flexql

import (
	"fmt"
)

type Field struct {
	Name  string
	Type  Type
	alias string
}

func (f *Field) String() string {
	return fmt.Sprintf("%s.%s", f.alias, f.Name)
}

func (f *Field) StringAs() string {
	return fmt.Sprintf("%s.%s as %s_%s", f.alias, f.Name, f.alias, f.Name)
}

func (f *Field) As() string {
	return fmt.Sprintf("%s_%s", f.alias, f.Name)
}
