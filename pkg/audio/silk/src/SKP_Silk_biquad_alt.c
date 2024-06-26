
#include "SKP_Silk_SigProc_FIX.h"


/* Second order ARMA filter, alternative implementation */
void SKP_Silk_biquad_alt(
    const SKP_int16      *in,            /* I:    Input signal                   */
    const SKP_int32      *B_Q28,         /* I:    MA coefficients [3]            */
    const SKP_int32      *A_Q28,         /* I:    AR coefficients [2]            */
    SKP_int32            *S,             /* I/O: State vector [2]                */
    SKP_int16            *out,           /* O:    Output signal                  */
    const SKP_int32      len             /* I:    Signal length (must be even)   */
)
{
    /* DIRECT FORM II TRANSPOSED (uses 2 element state vector) */
    SKP_int   k;
    SKP_int32 inval, A0_U_Q28, A0_L_Q28, A1_U_Q28, A1_L_Q28, out32_Q14;

    /* Negate A_Q28 values and split in two parts */
    A0_L_Q28 = ( -A_Q28[ 0 ] ) & 0x00003FFF;        /* lower part */
    A0_U_Q28 = SKP_RSHIFT( -A_Q28[ 0 ], 14 );       /* upper part */
    A1_L_Q28 = ( -A_Q28[ 1 ] ) & 0x00003FFF;        /* lower part */
    A1_U_Q28 = SKP_RSHIFT( -A_Q28[ 1 ], 14 );       /* upper part */
    
    for( k = 0; k < len; k++ ) {
        /* S[ 0 ], S[ 1 ]: Q12 */
        inval = in[ k ];
        out32_Q14 = SKP_LSHIFT( SKP_SMLAWB( S[ 0 ], B_Q28[ 0 ], inval ), 2 );

        S[ 0 ] = S[1] + SKP_RSHIFT_ROUND( SKP_SMULWB( out32_Q14, A0_L_Q28 ), 14 );
        S[ 0 ] = SKP_SMLAWB( S[ 0 ], out32_Q14, A0_U_Q28 );
        S[ 0 ] = SKP_SMLAWB( S[ 0 ], B_Q28[ 1 ], inval);

        S[ 1 ] = SKP_RSHIFT_ROUND( SKP_SMULWB( out32_Q14, A1_L_Q28 ), 14 );
        S[ 1 ] = SKP_SMLAWB( S[ 1 ], out32_Q14, A1_U_Q28 );
        S[ 1 ] = SKP_SMLAWB( S[ 1 ], B_Q28[ 2 ], inval );

        /* Scale back to Q0 and saturate */
        out[ k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT( out32_Q14 + (1<<14) - 1, 14 ) );
    }
}
