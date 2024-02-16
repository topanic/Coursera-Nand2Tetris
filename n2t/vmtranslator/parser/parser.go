package parser

import (
	"strings"
	"translator/ast"
	"translator/token"
)


type Parser struct {
	CurrentCommandIdx int
	CurrentTokenIdx int
	CommandStrArr []string	// ["push local 1", ...]
	CurrentCommandTokenArr []string  // ["push","local","1"]
	input string
}

func New(input string) *Parser {
	commandStrArr := strings.Split(input, token.NEW_LINE)
	initialCurrentCommandTokenArr := strings.Split(commandStrArr[0], token.SPACE)
	return &Parser{
		CurrentCommandIdx: 0,
		CurrentTokenIdx: 0,
		CommandStrArr: commandStrArr,
		CurrentCommandTokenArr: initialCurrentCommandTokenArr,
		input: input,
	}
}

func (p *Parser) HasMoreCommand() bool {
	return len(p.CommandStrArr) > p.CurrentCommandIdx
}

func (p *Parser) CommandType() ast.CommandType {
	if p.CommandStrArr[p.CurrentCommandIdx] == "" {
		return ast.C_EMPTY
	}
	currentCommandPrefix := ast.CommandSymbol(p.CurrentCommandTokenArr[0])
	switch currentCommandPrefix {
	case ast.PUSH:
		return ast.C_PUSH
	case ast.POP:
		return ast.C_POP
	case ast.ADD, ast.SUB, ast.NEG, ast.EQ, ast.GT, ast.LT, ast.AND, ast.OR, ast.NOT:
		return ast.C_ARITHMETIC
	case ast.LABEL:
		return ast.C_LABEL
	case ast.GOTO:
		return ast.C_GOTO
	case ast.IF_GOTO:
		return ast.C_IF
	case ast.CALL:
		return ast.C_CALL
	case ast.FUNCTION:
		return ast.C_FUNCTION
	case ast.RETURN:
		return ast.C_RETURN
	default:
		return ast.C_EMPTY
	
	}
}

func (p *Parser) Advance() {
	for {
		p.CurrentCommandIdx++
		if !p.HasMoreCommand() {
			break
		}
		p.CurrentTokenIdx = 0
		p.CurrentCommandTokenArr = strings.Split(p.CommandStrArr[p.CurrentCommandIdx], token.SPACE)
		if p.CommandType() != ast.C_EMPTY {
			break
		}
	}
}
