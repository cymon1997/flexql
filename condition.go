package flexql

import "fmt"

type Condition struct {
	A, B     interface{}
	Operator string
}

func (c *Condition) String() string {
	return fmt.Sprintf("%s %s %s",
		c.Parse(c.A), c.Operator, c.Parse(c.B))
}

func (c *Condition) Parse(arg interface{}) string {
	if f, ok := arg.(*Field); ok {
		return f.String()
	} else {
		return fmt.Sprintf("'%v'", arg)
	}
}
