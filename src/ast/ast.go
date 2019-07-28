package ast

var curTable Table

type Column struct {
	Name     string
	key        bool
	uniqueKey  bool
	primaryKey bool
	comment    string
}

type Table struct {
	Name string
	Comment   string
	Columns   []Column
}

type Program struct {
	Tables []Table
}
