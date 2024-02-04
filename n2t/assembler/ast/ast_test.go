package ast

import (
	"assembler/token"
	"testing"
)




func TestACommandValue(t *testing.T) {
	val := 10000
	aCommand := ACommand{Value: val, StringValue: "10000"}
	if aCommand.String() != "@10000" + token.NEW_LINE {
		t.Fatalf("aCommand.String() should be %s, got %s ", "@10000" + token.NEW_LINE, aCommand.String())
	}
}


func TestCCommand(t *testing.T) {
	testCases := []struct {
		command CCommand
		commandstr string
	} {
		{CCommand{Comp: "-1", Dest: "M"}, "M=-1" + token.NEW_LINE},
		{CCommand{Comp: "D", Jump: "JMP"}, "D;JMP" + token.NEW_LINE},
		{CCommand{Comp: "D|A", Dest: "AM", Jump: "JMP"}, "AM=D|A;JMP" + token.NEW_LINE},
	}
	for _, tt := range testCases {
		if tt.command.String() != tt.commandstr {
			t.Fatalf("CCommand.String() should be %s, got %s ", tt.commandstr, tt.command.String())
		}
	}
	
}