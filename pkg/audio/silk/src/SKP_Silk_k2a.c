 

/*                                                                      *
 * SKP_Silk_k2a.c                                                     *
 *                                                                      *
 * Step up function, converts reflection coefficients to prediction     *
 * coefficients                                                         *
 *                                                                      *
 * Copyright 2008 (c), Skype Limited                                    *
 * Date: 080103                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

/* Step up function, converts reflection coefficients to prediction coefficients */
void SKP_Silk_k2a(
    SKP_int32            *A_Q24,                 /* O:    Prediction coefficients [order] Q24         */
    const SKP_int16      *rc_Q15,                /* I:    Reflection coefficients [order] Q15         */
    const SKP_int32      order                   /* I:    Prediction order                            */
)
{
    SKP_int   k, n;
    SKP_int32 Atmp[ SKP_Silk_MAX_ORDER_LPC ];

    for( k = 0; k < order; k++ ) {
        for( n = 0; n < k; n++ ) {
            Atmp[ n ] = A_Q24[ n ];
        }
        for( n = 0; n < k; n++ ) {
            A_Q24[ n ] = SKP_SMLAWB( A_Q24[ n ], SKP_LSHIFT( Atmp[ k - n - 1 ], 1 ), rc_Q15[ k ] );
        }
        A_Q24[ k ] = -SKP_LSHIFT( (SKP_int32)rc_Q15[ k ], 9 );
    }
}
