 

/*                                                                      *
 * SKP_Silk_log2lin.c                                                 *
 *                                                                      *
 * Convert input to a linear scale                                      *
 *                                                                      *
 * Copyright 2006 (c), Skype Limited                                    *
 * Date: 060221                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

/* Approximation of 2^() (very close inverse of SKP_Silk_lin2log()) */
/* Convert input to a linear scale    */ 
SKP_int32 SKP_Silk_log2lin( const SKP_int32 inLog_Q7 )    /* I:    Input on log scale */ 
{
    SKP_int32 out, frac_Q7;

    if( inLog_Q7 < 0 ) {
        return( 0 );
    } else if( inLog_Q7 >= ( 31 << 7 ) ) {
        /* Saturate, and prevent wrap-around */
        return( SKP_int32_MAX );
    }

    out = SKP_LSHIFT( 1, SKP_RSHIFT( inLog_Q7, 7 ) );
    frac_Q7 = inLog_Q7 & 0x7F;
    if( inLog_Q7 < 2048 ) {
        /* Piece-wise parabolic approximation */
        out = SKP_ADD_RSHIFT( out, SKP_MUL( out, SKP_SMLAWB( frac_Q7, SKP_MUL( frac_Q7, 128 - frac_Q7 ), -174 ) ), 7 );
    } else {
        /* Piece-wise parabolic approximation */
        out = SKP_MLA( out, SKP_RSHIFT( out, 7 ), SKP_SMLAWB( frac_Q7, SKP_MUL( frac_Q7, 128 - frac_Q7 ), -174 ) );
    }
    return out;
}
