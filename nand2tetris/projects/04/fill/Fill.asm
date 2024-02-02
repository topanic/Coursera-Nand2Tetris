// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen
// by writing 'black' in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen by writing
// 'white' in every pixel;
// the screen should remain fully clear as long as no key is pressed.

//// Replace this comment with your code.

// 有输入就涂黑，没有输入就涂白

(LOOP)
    @KBD
    D = M
    @BLACK  
    D; JNE  // goto black if KBD
    @WHITE
    D; JEQ  // goto white if not KBD


(BLACK)
    // sc = SCREEN address(16384)
    @SCREEN
    D = A
    @sc
    M = D

    // n = KBD address(24576) - SCREEN address(16384)
    @KBD
    D = A
    @SCREEN
    D = D - A
    @n
    M = D

    // initial i = 0
    @i
    M = 0

    (LOOPB)
        // if (i == n) goto LOOP
        @i
        D = M
        @n
        D = D - M
        @LOOP
        D; JEQ

        // RAM[sc + i] = -1
        @sc
        D = M
        @i
        A = D + M
        M = -1

        // i++
        @i
        M = M + 1
        
        @LOOPB
        0; JMP


(WHITE)
    // sc = SCREEN address(16384)
    @SCREEN
    D = A
    @sc
    M = D

    // n = KBD address(24576) - SCREEN address(16384)
    @KBD
    D = A
    @SCREEN
    D = D - A
    @n
    M = D

    // initial i = 0
    @i
    M = 0

    (LOOPW)
        // if (i == n) goto LOOP
        @i
        D = M
        @n
        D = D - M
        @LOOP
        D; JEQ

        // RAM[sc + i] = -1
        @sc
        D = M
        @i
        A = D + M
        M = 0

        // i++
        @i
        M = M + 1
        
        @LOOPW
        0; JMP