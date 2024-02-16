package ast

import "fmt"

type CommandType string

const (
	C_PUSH       CommandType = "C_PUSH"
	C_POP        CommandType = "C_POP"
	C_ARITHMETIC CommandType = "C_ARITHMETIC"
	C_LABEL      CommandType = "C_LABEL"
	C_GOTO       CommandType = "C_GOTO"
	C_IF         CommandType = "C_IF"
	C_FUNCTION   CommandType = "C_FUNCTION"
	C_RETURN     CommandType = "C_RETURN"
	C_CALL       CommandType = "C_CALL"
	C_EMPTY      CommandType = "C_EMPTY"
)

type CommandSymbol string

const (
	PUSH     CommandSymbol = "push"
	POP      CommandSymbol = "pop"
	LABEL    CommandSymbol = "label"
	GOTO     CommandSymbol = "goto"
	IF_GOTO  CommandSymbol = "if-goto"
	CALL     CommandSymbol = "call"
	FUNCTION CommandSymbol = "function"
	RETURN   CommandSymbol = "return"
	ADD      CommandSymbol = "add"
	SUB      CommandSymbol = "sub"
	NEG      CommandSymbol = "neg"
	EQ       CommandSymbol = "eq"
	GT       CommandSymbol = "gt"
	LT       CommandSymbol = "lt"
	AND      CommandSymbol = "and"
	OR       CommandSymbol = "or"
	NOT      CommandSymbol = "not"
)

type SegmentType string

const (
	ARGUMENT SegmentType = "argument"
	LOCAL    SegmentType = "local"
	STATIC   SegmentType = "static"
	CONSTANT SegmentType = "constant"
	THIS     SegmentType = "this"
	THAT     SegmentType = "that"
	POINTER  SegmentType = "pointer"
	TEMP     SegmentType = "temp"
)


type VmCommand interface {
	String() string
}


type ArithmeticCommand struct {
	Command CommandType		// C_ARITHMETIC
	Symbol  CommandSymbol	// ADD, SUB, NEG, EQ, GT, LT, AND, OR, NOT
}

func (ac *ArithmeticCommand) String() string {
	return string(ac.Symbol)
}


// Memory Access Command, "push" and "pop"
type MemoryAccessCommand interface {
	VmCommand
}

type PushCommand struct {
	Command CommandType	// C_PUSH
	Symbol CommandSymbol	// push
	Segment SegmentType
	Index int
}

func (pc *PushCommand) String() string {
	return fmt.Sprintf("%s %s %d", pc.Symbol, pc.Segment, pc.Index)
}

type PopCommand struct {
	Command CommandType // C_POP
	Symbol CommandSymbol // pop
	Segment SegmentType
	Index int 
}

func (pc *PopCommand) String() string {
	return fmt.Sprintf("%s %s %d", pc.Symbol, pc.Segment, pc.Index)
}

type LabelCommand struct {
	Command CommandType
	Symbol CommandSymbol
	LabelName string
}

func (lc *LabelCommand) String() string {
	return fmt.Sprintf("%s %s", lc.Symbol, lc.LabelName)
}

type GotoCommand struct {
	Command CommandType
	Symbol CommandSymbol
	LabelName string
}

func (gc *GotoCommand) String() string {
	return fmt.Sprintf("%s %s", gc.Symbol, gc.LabelName)
}

type IfCommand struct {
	Command CommandType
	Symbol CommandSymbol
	LabelName string
}

func (ic *IfCommand) String() string {
	return fmt.Sprintf("%s %s", ic.Symbol, ic.LabelName) 
}

type CallCommand struct {
	Command CommandType
	Symbol CommandSymbol
	FunctionName string
	NumArgs int
}

func (cc *CallCommand) String() string {
	return fmt.Sprintf("%s %s %d", cc.Symbol, cc.FunctionName, cc.NumArgs)
}

type FunctionCommand struct {
	Command CommandType
	Symbol CommandSymbol
	FunctionName string
	NumLocals int
}

func (fc *FunctionCommand) String() string {
	return fmt.Sprintf("%s %s %d", fc.Symbol, fc.FunctionName, fc.NumLocals) 
}

type ReturnCommand struct {
	Command CommandType
	Symbol CommandSymbol
}

func (rc *ReturnCommand) String() string {
	return string(rc.Symbol)
}
