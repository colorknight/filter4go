package metadata

import (
	"fmt"
	"strings"
)

const (
	INFINITE = 0
)

type OperatorType int

const (
	UNARY OperatorType = iota
	BINARY
)

type OperatableMetadata interface{}

type OperationMetadata struct {
	OperatorType OperatorType
	Operator     string
	LOperand     OperatableMetadata
	ROperand     OperatableMetadata
}

type RelationOperationMetadata struct {
	OperatorType OperatorType
	Operator     string
	LOperand     string
	ROperand     string
}

func ToString(meta OperatableMetadata) string {
	var bd strings.Builder
	buildString(meta, &bd)
	return bd.String()
}

func buildString(meta OperatableMetadata, bd *strings.Builder) error {
	om, ok := meta.(OperationMetadata)
	if ok {
		if om.OperatorType == BINARY {
			buildString(om.LOperand, bd)
			bd.WriteByte(' ')
		}
		bd.WriteString(om.Operator)
		bd.WriteByte(' ')
		buildString(om.ROperand, bd)
	} else {
		rom, ok := meta.(RelationOperationMetadata)
		if !ok {
			return fmt.Errorf("Unsupported OperatableMetadata!")
		}
		if rom.OperatorType == BINARY {
			bd.WriteString(rom.LOperand)
			bd.WriteByte(' ')
		}
		bd.WriteString(rom.Operator)
		bd.WriteByte(' ')
		bd.WriteString(rom.ROperand)
	}
	return nil
}
