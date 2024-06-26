 

/*                                                                      *
 * SKP_Silk_schur.c                                                   *
 *                                                                      *
 * Calculates the reflection coefficients from the correlation sequence *
 *                                                                      *
 * Copyright 2008 (c), Skype Limited                                    *
 * Date: 080103                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

/* Faster than schur64(), but much less accurate.                       */
/* uses SMLAWB(), requiring armv5E and higher.                          */ 
SKP_int32 SKP_Silk_schur(                     /* O:    Returns residual energy                     */
    SKP_int16            *rc_Q15,               /* O:    reflection coefficients [order] Q15         */
    const SKP_int32      *c,                    /* I:    correlations [order+1]                      */
    const SKP_int32      order                  /* I:    prediction order                            */
)
{
    SKP_int        k, n, lz;
    SKP_int32    C[ SKP_Silk_MAX_ORDER_LPC + 1 ][ 2 ];
    SKP_int32    Ctmp1, Ctmp2, rc_tmp_Q15;

    /* Get number of leading zeros */
    lz = SKP_Silk_CLZ32( c[ 0 ] );

    /* Copy correlations and adjust level to Q30 */
    if( lz < 2 ) {
        /* lz must be 1, so shift one to the right */
        for( k = 0; k < order + 1; k++ ) {
            C[ k ][ 0 ] = C[ k ][ 1 ] = SKP_RSHIFT( c[ k ], 1 );
        }
    } else if( lz > 2 ) {
        /* Shift to the left */
        lz -= 2; 
        for( k = 0; k < order + 1; k++ ) {
            C[ k ][ 0 ] = C[ k ][ 1 ] = SKP_LSHIFT( c[k], lz );
        }
    } else {
        /* No need to shift */
        for( k = 0; k < order + 1; k++ ) {
            C[ k ][ 0 ] = C[ k ][ 1 ] = c[ k ];
        }
    }

    for( k = 0; k < order; k++ ) {
        
        /* Get reflection coefficient */
        rc_tmp_Q15 = -SKP_DIV32_16( C[ k + 1 ][ 0 ], SKP_max_32( SKP_RSHIFT( C[ 0 ][ 1 ], 15 ), 1 ) );

        /* Clip (shouldn't happen for properly conditioned inputs) */
        rc_tmp_Q15 = SKP_SAT16( rc_tmp_Q15 );

        /* Store */
        rc_Q15[ k ] = (SKP_int16)rc_tmp_Q15;

        /* Update correlations */
        for( n = 0; n < order - k; n++ ) {
            Ctmp1 = C[ n + k + 1 ][ 0 ];
            Ctmp2 = C[ n ][ 1 ];
            C[ n + k + 1 ][ 0 ] = SKP_SMLAWB( Ctmp1, SKP_LSHIFT( Ctmp2, 1 ), rc_tmp_Q15 );
            C[ n ][ 1 ]         = SKP_SMLAWB( Ctmp2, SKP_LSHIFT( Ctmp1, 1 ), rc_tmp_Q15 );
        }
    }

    /* return residual energy */
    return C[0][1];
}
