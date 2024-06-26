/*                                                                          *
 * SKP_ana_filt_bank_1.c                                                    *
 *                                                                          *
 * Split signal into two decimated bands using first-order allpass filters  *
 *                                                                          *
 * Copyright 2006 (c), Skype Limited                                        *
 * Date: 060221                                                             *
 *                                                                          */
#include "SKP_Silk_SigProc_FIX.h"

#if EMBEDDED_ARM<5
/* Coefficients for 2-band filter bank based on first-order allpass filters */
// old
static SKP_int16 A_fb1_20[ 1 ] = {  5394 << 1 };
static SKP_int16 A_fb1_21[ 1 ] = {  (SKP_int16) (20623 << 1) };        /* wrap-around to negative number is intentional */

/* Split signal into two decimated bands using first-order allpass filters */
void SKP_Silk_ana_filt_bank_1(
    const SKP_int16      *in,        /* I:   Input signal [N]        */
    SKP_int32            *S,         /* I/O: State vector [2]        */
    SKP_int16            *outL,      /* O:   Low band [N/2]          */
    SKP_int16            *outH,      /* O:   High band [N/2]         */
    SKP_int32            *scratch,   /* I:   Scratch memory [3*N/2]  */   // todo: remove - no longer used
    const SKP_int32      N           /* I:   Number of input samples */
)
{
    SKP_int      k, N2 = SKP_RSHIFT( N, 1 );
    SKP_int32    in32, X, Y, out_1, out_2;

    /* Internal variables and state are in Q10 format */
    for( k = 0; k < N2; k++ ) {
        /* Convert to Q10 */
        in32 = SKP_LSHIFT( (SKP_int32)in[ 2 * k ], 10 );

        /* All-pass section for even input sample */
        Y      = SKP_SUB32( in32, S[ 0 ] );
        X      = SKP_SMLAWB( Y, Y, A_fb1_21[ 0 ] );
        out_1  = SKP_ADD32( S[ 0 ], X );
        S[ 0 ] = SKP_ADD32( in32, X );

        /* Convert to Q10 */
        in32 = SKP_LSHIFT( (SKP_int32)in[ 2 * k + 1 ], 10 );

        /* All-pass section for odd input sample */
        Y      = SKP_SUB32( in32, S[ 1 ] );
        X      = SKP_SMULWB( Y, A_fb1_20[ 0 ] );
        out_2  = SKP_ADD32( S[ 1 ], X );
        S[ 1 ] = SKP_ADD32( in32, X );

        /* Add/subtract, convert back to int16 and store to output */
        outL[ k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( SKP_ADD32( out_2, out_1 ), 11 ) );
        outH[ k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( SKP_SUB32( out_2, out_1 ), 11 ) );
    }
}
#endif
