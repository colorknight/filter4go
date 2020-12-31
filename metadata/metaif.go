package metadata

type ColumnDefinition interface {
	GetName() string
	GetValue() string
}

type SelectorDefinition interface {}