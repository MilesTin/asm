#include "textflag.h"

TEXT ·Arccos(SB),$8-16

    FMOVD   x+0(FP), F0
    FMOVD   F0, F1
    FMULD   F0, F0
    FLD1

    FSUBRDP F0, F1
    FSQRT
    FLD1
    FSUBD   $1.0, F0
    FADDD   F2, F0
    FPATAN

    FMOVDP F0, ret+8(FP)
    RET
