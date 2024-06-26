 

/*                                                                      *
 * SKP_Silk_lin2log.c                                                 *
 *                                                                      *
 * Convert input to a log scale                                         *
 * Approximation of 128 * log2()                                        *
 *                                                                      *
 * Copyright 2006 (c), Skype Limited                                    *
 * Date: 060221                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"
#if EMBEDDED_ARM<4
/* Approximation of 128 * log2() (very close inverse of approx 2^() below) */
/* Convert input to a log scale    */ 
SKP_int32 SKP_Silk_lin2log( const SKP_int32 inLin )    /* I:    Input in linear scale */
{
    SKP_int32 lz, frac_Q7;

    SKP_Silk_CLZ_FRAC( inLin, &lz, &frac_Q7 );

    /* Piece-wise parabolic approximation */
    return( SKP_LSHIFT( 31 - lz, 7 ) + SKP_SMLAWB( frac_Q7, SKP_MUL( frac_Q7, 128 - frac_Q7 ), 179 ) );
}
#endif

