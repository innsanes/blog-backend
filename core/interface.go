package core

type Confer interface {
	RegisterConfWithName(name string, confStruct interface{})
}
