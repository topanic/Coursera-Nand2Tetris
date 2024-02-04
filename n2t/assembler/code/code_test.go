package code

import (
	"assembler/ast"
	"testing"
)

func TestBinary(t *testing.T) {
	testCases := []struct {
		command   ast.Command
		binaryStr string
	}{
		{&ast.ACommand{Value: 100}, "0000000001100100"},
		{&ast.CCommand{Comp: "A", Dest: "D"}, "1110110000010000"},
		// {&ast.CCommand{Comp: "M", Dest: "D"}, "1110001100001000"},
		{&ast.CCommand{Comp: "D|A", Dest: "AM", Jump: "JMP"}, "1110010101101111"},
		{&ast.CCommand{
			Comp: "D+M",
			Dest: "AMD",
			Jump: "JLE",
		}, "1111000010111110"},
		{
			&ast.CCommand {
				Comp: "M",
				Dest: "D",
			},
			"1111110000010000",
		},
	}
	for _, tt := range testCases {
		binaryStr := Binary(tt.command)
		if tt.binaryStr != binaryStr {
			t.Fatalf("binaryStr should be %s, got %s ", tt.binaryStr, binaryStr)
		}
	}
}