#include "textflag.h"

TEXT ·Arcsin(SB), NOSPLIT,$0-16

    FMOVD   x+0(FP), F0
    FMOVD   F0, F1
    FMULD   F0, F0
    FLD1

    FSUBRDP F0, F1
    FSQRT
    FPATAN

    FMOVDP F0, ret+8(FP)
    RET
