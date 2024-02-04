package symboltable

import "fmt"

/*
	when we meet @symbol ast, we need to store the symbol and its address in the symbol table, so this is why we need a symbol table.

	Add initial symbols to the table, and add new symbols as we encounter them in the assembly code.

	Define a SymbolTable struct with a map field to store the symbols and their addresses. Implement some methods to add and check symbols in the table.
*/

// Define SymbolTable struct
type SymbolTable struct {
	table map[string]int
}

// NewSymbolTable returns a new SymbolTable instance
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table: map[string]int{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"R0":     0,
			"R1":     1,
			"R2":     2,
			"R3":     3,
			"R4":     4,
			"R5":     5,
			"R6":     6,
			"R7":     7,
			"R8":     8,
			"R9":     9,
			"R10":    10,
			"R11":    11,
			"R12":    12,
			"R13":    13,
			"R14":    14,
			"R15":    15,
			"SCREEN": 16384,
			"KBD":    24576,
		},
	}
}

// if contians symbol in the table
func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}

// AddEntry adds a new symbol to the table
func (st *SymbolTable) AddEntry(symbol string, address int) error {
	st.table[symbol] = address
	return nil
}

// GetAddress returns the address of a symbol in the table
func (st *SymbolTable) GetAddress(symbol string) (int, error) {
	if !st.Contains(symbol) {
		return -1, fmt.Errorf("symbol %s not found in table", symbol)
	}
	return st.table[symbol], nil
}