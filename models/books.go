package models

type Book struct {
	Bar string   `minLength:"4" maxLength:"16" example:"random string"`
	Baz int      `minimum:"10" maximum:"20" default:"15"`
	Qux []string `enums:"foo,bar,baz"`
}
