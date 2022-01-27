#include "textflag.h"

// 交换int64
TEXT ·SwapInt64(SB), NOSPLIT,$0-24
    MOVQ    addr+0(FP), BP
    MOVQ    new+8(FP), AX
    XCHGQ   AX, 0(BP)
    MOVQ    AX, old+16(FP)
    RET
