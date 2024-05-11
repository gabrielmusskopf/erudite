package main

import (
	"fmt"
	"strings"
)

type Flags []string

func (i *Flags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *Flags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Flag[T any] struct {
	Usage     string
	ShortName string
	LongName  string
	Value     T
}

func (f Flag[T]) Names() (short, long string) {
	return f.ShortName, f.LongName
}

func (f Flag[T]) FullUsage() string {
	s := "  "
	short, long := f.Names()
	if short != "" {
		s += "-" + short + ", "
	 }
	if long != "" {
		s += "--" + long
	}
	s += ":" + strings.Repeat(" ", 25-len(s)) + f.Usage
	return s
}
