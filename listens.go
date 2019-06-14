package main

import (
	"fmt"
)

type Listens []string

func (l *Listens) String() string {
	return fmt.Sprintf("%s", *l)
}

func (l *Listens) Set(value string) error {
	*l = append(*l, value)
	return nil
}
