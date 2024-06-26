
#include "SKP_Silk_SigProc_FIX.h"

/* Second order ARMA filter */
/* Can handle slowly varying filter coefficients */
void SKP_Silk_biquad(
    const SKP_int16      *in,        /* I:    input signal               */
    const SKP_int16      *B,         /* I:    MA coefficients, Q13 [3]   */
    const SKP_int16      *A,         /* I:    AR coefficients, Q13 [2]   */
    SKP_int32            *S,         /* I/O:  state vector [2]           */
    SKP_int16            *out,       /* O:    output signal              */
    const SKP_int32      len         /* I:    signal length              */
)
{
    SKP_int   k, in16;
    SKP_int32 A0_neg, A1_neg, S0, S1, out32, tmp32;

    S0 = S[ 0 ];
    S1 = S[ 1 ];
    A0_neg = -A[ 0 ];
    A1_neg = -A[ 1 ];
    for( k = 0; k < len; k++ ) {
        /* S[ 0 ], S[ 1 ]: Q13 */
        in16  = in[ k ];
        out32 = SKP_SMLABB( S0, in16, B[ 0 ] );

        S0 = SKP_SMLABB( S1, in16, B[ 1 ] );
        S0 += SKP_LSHIFT( SKP_SMULWB( out32, A0_neg ), 3 );

        S1 = SKP_LSHIFT( SKP_SMULWB( out32, A1_neg ), 3 );
        S1 = SKP_SMLABB( S1, in16, B[ 2 ] );
        tmp32    = SKP_RSHIFT_ROUND( out32, 13 ) + 1;
        out[ k ] = (SKP_int16)SKP_SAT16( tmp32 );
    }
    S[ 0 ] = S0;
    S[ 1 ] = S1;
}
