package ast

import (
	"assembler/token"
	"fmt"
	"strings"
)

// Define command types
type CommandType string

const (
	A_COMMAND CommandType = "A_COMMAND"
	C_COMMAND CommandType = "C_COMMAND"
	L_COMMAND CommandType = "L_COMMAND"
)

// Define command interface, implementing String() method
type Command interface {
	String() string
}

// Define ACommand struct implementing Command interface
type ACommand struct {
	Value int
	StringValue string
}

func (ac *ACommand) String() string {
	return fmt.Sprintf("@%s", ac.StringValue) + token.NEW_LINE
}

// Define CCommand struct implementing Command interface
type CCommand struct {
	Comp string
	Dest string
	Jump string
}

func (cc *CCommand) String() string {
	commandStr := fmt.Sprintf("%s=%s;%s", cc.Dest, cc.Comp, cc.Jump) + token.NEW_LINE
	if cc.Jump == "" {
		commandStr = strings.Replace(commandStr, ";", "", 1)
	} else if cc.Dest == "" {
		commandStr = strings.Replace(commandStr, "=", "", 1)
	}
	return commandStr
}

// Define LCommand struct implementing Command interface
type LCommand struct {
	Symbol string
}

func (lc *LCommand) String() string {
	return fmt.Sprintf("(%s)", lc.Symbol)
}