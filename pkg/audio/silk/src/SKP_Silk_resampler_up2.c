 

/*                                                                      *
 * SKP_Silk_resampler_up2.c                                           *
 *                                                                      *
 * Upsample by a factor 2, low quality                                  *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_rom.h"

/* Upsample by a factor 2, low quality */
#if EMBEDDED_ARM<5
void SKP_Silk_resampler_up2(
    SKP_int32                           *S,         /* I/O: State vector [ 2 ]                  */
    SKP_int16                           *out,       /* O:   Output signal [ 2 * len ]           */
    const SKP_int16                     *in,        /* I:   Input signal [ len ]                */
    SKP_int32                           len         /* I:   Number of input samples             */
)
{
    SKP_int32 k;
    SKP_int32 in32, out32, Y, X;

    SKP_assert( SKP_Silk_resampler_up2_lq_0 > 0 );
    SKP_assert( SKP_Silk_resampler_up2_lq_1 < 0 );
    /* Internal variables and state are in Q10 format */
    for( k = 0; k < len; k++ ) {
        /* Convert to Q10 */
        in32 = SKP_LSHIFT( (SKP_int32)in[ k ], 10 );

        /* All-pass section for even output sample */
        Y      = SKP_SUB32( in32, S[ 0 ] );
        X      = SKP_SMULWB( Y, SKP_Silk_resampler_up2_lq_0 );
        out32  = SKP_ADD32( S[ 0 ], X );
        S[ 0 ] = SKP_ADD32( in32, X );

        /* Convert back to int16 and store to output */
        out[ 2 * k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( out32, 10 ) );

        /* All-pass section for odd output sample */
        Y      = SKP_SUB32( in32, S[ 1 ] );
        X      = SKP_SMLAWB( Y, Y, SKP_Silk_resampler_up2_lq_1 );
        out32  = SKP_ADD32( S[ 1 ], X );
        S[ 1 ] = SKP_ADD32( in32, X );

        /* Convert back to int16 and store to output */
        out[ 2 * k + 1 ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( out32, 10 ) );
    }
}
#endif
