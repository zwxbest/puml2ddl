package ast

var curTable Table

type Column struct {
	Name string
	Type string

	Comment string
}

type KEY struct {
	Exist   bool
	KeyName string
	Keys    string
}

type Table struct {
	Name    string
	Comment string
	Columns []Column
	Pk      KEY
	Uk      KEY
	Key     KEY
}

type Program struct {
	Tables []Table
}

