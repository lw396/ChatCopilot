 

/*                                                                      *
 * SKP_Silk_MA.c                                                      *
 *                                                                      *
 * Variable order MA filter                                             *
 *                                                                      *
 * Copyright 2006 (c), Skype Limited                                    *
 * Date: 060221                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

#if EMBEDDED_ARM<5
/* Variable order MA prediction error filter */
void SKP_Silk_MA_Prediction(
    const SKP_int16      *in,            /* I:   Input signal                                */
    const SKP_int16      *B,             /* I:   MA prediction coefficients, Q12 [order]     */
    SKP_int32            *S,             /* I/O: State vector [order]                        */
    SKP_int16            *out,           /* O:   Output signal                               */
    const SKP_int32      len,            /* I:   Signal length                               */
    const SKP_int32      order           /* I:   Filter order                                */
)
{
    SKP_int   k, d, in16;
    SKP_int32 out32;

    for( k = 0; k < len; k++ ) {
        in16 = in[ k ];
        out32 = SKP_LSHIFT( in16, 12 ) - S[ 0 ];
        out32 = SKP_RSHIFT_ROUND( out32, 12 );
        
        for( d = 0; d < order - 1; d++ ) {
            S[ d ] = SKP_SMLABB_ovflw( S[ d + 1 ], in16, B[ d ] );
        }
        S[ order - 1 ] = SKP_SMULBB( in16, B[ order - 1 ] );

        /* Limit */
        out[ k ] = (SKP_int16)SKP_SAT16( out32 );
    }
}
#endif

#if EMBEDDED_ARM<5

void SKP_Silk_LPC_analysis_filter(
    const SKP_int16      *in,            /* I:   Input signal                                */
    const SKP_int16      *B,             /* I:   MA prediction coefficients, Q12 [order]     */
    SKP_int16            *S,             /* I/O: State vector [order]                        */
    SKP_int16            *out,           /* O:   Output signal                               */
    const SKP_int32      len,            /* I:   Signal length                               */
    const SKP_int32      Order           /* I:   Filter order                                */
)
{
    SKP_int   k, j, idx, Order_half = SKP_RSHIFT( Order, 1 );
    SKP_int32 out32_Q12, out32;
    SKP_int16 SA, SB;
    /* Order must be even */
    SKP_assert( 2 * Order_half == Order );

    /* S[] values are in Q0 */
    for( k = 0; k < len; k++ ) {
        SA = S[ 0 ];
        out32_Q12 = 0;
        for( j = 0; j < ( Order_half - 1 ); j++ ) {
            idx = SKP_SMULBB( 2, j ) + 1;
            /* Multiply-add two prediction coefficients for each loop */
            SB = S[ idx ];
            S[ idx ] = SA;
            out32_Q12 = SKP_SMLABB( out32_Q12, SA, B[ idx - 1 ] );
            out32_Q12 = SKP_SMLABB( out32_Q12, SB, B[ idx ] );
            SA = S[ idx + 1 ];
            S[ idx + 1 ] = SB;
        }

        /* Unrolled loop: epilog */
        SB = S[ Order - 1 ];
        S[ Order - 1 ] = SA;
        out32_Q12 = SKP_SMLABB( out32_Q12, SA, B[ Order - 2 ] );
        out32_Q12 = SKP_SMLABB( out32_Q12, SB, B[ Order - 1 ] );

        /* Subtract prediction */
        out32_Q12 = SKP_SUB_SAT32( SKP_LSHIFT( (SKP_int32)in[ k ], 12 ), out32_Q12 );

        /* Scale to Q0 */
        out32 = SKP_RSHIFT_ROUND( out32_Q12, 12 );

        /* Saturate output */
        out[ k ] = ( SKP_int16 )SKP_SAT16( out32 );

        /* Move input line */
        S[ 0 ] = in[ k ];
    }
}
#endif

