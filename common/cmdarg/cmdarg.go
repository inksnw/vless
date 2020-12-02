package cmdarg

import "strings"

type Arg []string

func (receiver *Arg) String() string {
	return strings.Join([]string(*receiver), " ")
}

func (receiver *Arg) Set(value string) error {
	*receiver = append(*receiver, value)
	return nil

}
