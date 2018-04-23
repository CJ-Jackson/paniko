package uri

import "fmt"

type Reverse map[string]string

func (r Reverse) Print(name string) string {
	return r[name]
}

func (r Reverse) PrintF(name string, a ...interface{}) string {
	return fmt.Sprintf(r[name], a...)
}
