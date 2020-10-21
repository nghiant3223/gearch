package main

import "flag"

type Result struct {
	Path        string
	Description string
	URL         string
}

type Flag struct {
	Name        string
	Description string
}

type IntFlag struct {
	Flag
	DefaultValue int
	valuePointer *int
}

func (f *IntFlag) Read() {
	f.valuePointer = flag.Int(f.Name, f.DefaultValue, f.Description)
}

func (f *IntFlag) Value() int {
	return *f.valuePointer
}
