
TEXT ·add(SB), $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    ADDQ AX, BX
    MOVQ BX, ret+16(FP)
    RET
