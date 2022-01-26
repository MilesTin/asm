#include "textflag.h"

TEXT ·Arccos(SB),NOSPLIT,$0-16

    FMOVD x+0(FP), F0
    FMOVD x+0(FP), F1
    FMULD F0, F0
    FLD1

    FSUBRDP F0, F1

    FSQRT

    FMOVD F1, F2
    FMOVD F0, F1
    FMOVD F2, F0

    FPATAN

    FMOVDP F0, ret+8(FP)

    RET