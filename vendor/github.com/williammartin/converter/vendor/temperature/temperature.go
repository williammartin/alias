package temperature

import "fmt"

type Temperature struct {
	Value int
	Unit  string
}

func (t *Temperature) String() string {
	return fmt.Sprintf("%d %s", t.Value, t.Unit)
}
