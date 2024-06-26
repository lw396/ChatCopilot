 

/*                                                                      *
 * SKP_Silk_schur64.c                                                 *
 *                                                                      *
 * Calculates the reflection coefficients from the correlation sequence *
 * using extra precision                                                *
 *                                                                      *
 * Copyright 2008 (c), Skype Limited                                    *
 * Date: 080103                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

/* Slower than schur(), but more accurate.                              */
/* Uses SMULL(), available on armv4                                     */ 
#if EMBEDDED_ARM<6
SKP_int32 SKP_Silk_schur64(                    /* O:    Returns residual energy                     */
    SKP_int32            rc_Q16[],               /* O:    Reflection coefficients [order] Q16         */
    const SKP_int32      c[],                    /* I:    Correlations [order+1]                      */
    SKP_int32            order                   /* I:    Prediction order                            */
)
{
    SKP_int   k, n;
    SKP_int32 C[ SKP_Silk_MAX_ORDER_LPC + 1 ][ 2 ];
    SKP_int32 Ctmp1_Q30, Ctmp2_Q30, rc_tmp_Q31;

    /* Check for invalid input */
    if( c[ 0 ] <= 0 ) {
        SKP_memset( rc_Q16, 0, order * sizeof( SKP_int32 ) );
        return 0;
    }
    
    for( k = 0; k < order + 1; k++ ) {
        C[ k ][ 0 ] = C[ k ][ 1 ] = c[ k ];
    }

    for( k = 0; k < order; k++ ) {
        /* Get reflection coefficient: divide two Q30 values and get result in Q31 */
        rc_tmp_Q31 = SKP_DIV32_varQ( -C[ k + 1 ][ 0 ], C[ 0 ][ 1 ], 31 );

        /* Save the output */
        rc_Q16[ k ] = SKP_RSHIFT_ROUND( rc_tmp_Q31, 15 );

        /* Update correlations */
        for( n = 0; n < order - k; n++ ) {
            Ctmp1_Q30 = C[ n + k + 1 ][ 0 ];
            Ctmp2_Q30 = C[ n ][ 1 ];
            
            /* Multiply and add the highest int32 */
            C[ n + k + 1 ][ 0 ] = Ctmp1_Q30 + SKP_SMMUL( SKP_LSHIFT( Ctmp2_Q30, 1 ), rc_tmp_Q31 );
            C[ n ][ 1 ]         = Ctmp2_Q30 + SKP_SMMUL( SKP_LSHIFT( Ctmp1_Q30, 1 ), rc_tmp_Q31 );
        }
    }

    return C[ 0 ][ 1 ];
}
#endif
