// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

// for (int i = 0; i <= R1; i++) {
//     R2 += R1;   
// }
// 没有乘法，需要加法循环

@R2
M = 0
@i
M = 0


(LOOP)
    @i
    D = M   //D = i
    @R1
    D = M - D   
    @END
    D; JEQ  // if (R1 - i == 0)
    @R0
    D = M
    @R2
    M = M + D   // R2 += R1
    @i
    M = M + 1   // i += 1
    @LOOP
    0; JMP


(END)
    @END
    0; JMP