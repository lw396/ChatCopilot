 

/*                                                                      *
 * SKP_Silk_resampler_private_down4.c                                 *
 *                                                                      *
 * Downsample by a factor 4                                             *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

/* Downsample by a factor 4. Note: very low quality, only use with input sampling rates above 96 kHz. */
void SKP_Silk_resampler_private_down4(
    SKP_int32                       *S,             /* I/O: State vector [ 2 ]                      */
    SKP_int16                       *out,           /* O:   Output signal [ floor(len/2) ]          */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       inLen           /* I:   Number of input samples                 */
)
{
    SKP_int32 k, len4 = SKP_RSHIFT32( inLen, 2 );
    SKP_int32 in32, out32, Y, X;

    SKP_assert( SKP_Silk_resampler_down2_0 > 0 );
    SKP_assert( SKP_Silk_resampler_down2_1 < 0 );

    /* Internal variables and state are in Q10 format */
    for( k = 0; k < len4; k++ ) {
        /* Add two input samples and convert to Q10 */
        in32 = SKP_LSHIFT( SKP_ADD32( (SKP_int32)in[ 4 * k ], (SKP_int32)in[ 4 * k + 1 ] ), 9 );

        /* All-pass section for even input sample */
        Y      = SKP_SUB32( in32, S[ 0 ] );
        X      = SKP_SMLAWB( Y, Y, SKP_Silk_resampler_down2_1 );
        out32  = SKP_ADD32( S[ 0 ], X );
        S[ 0 ] = SKP_ADD32( in32, X );

        /* Add two input samples and convert to Q10 */
        in32 = SKP_LSHIFT( SKP_ADD32( (SKP_int32)in[ 4 * k + 2 ], (SKP_int32)in[ 4 * k + 3 ] ), 9 );

        /* All-pass section for odd input sample */
        Y      = SKP_SUB32( in32, S[ 1 ] );
        X      = SKP_SMULWB( Y, SKP_Silk_resampler_down2_0 );
        out32  = SKP_ADD32( out32, S[ 1 ] );
        out32  = SKP_ADD32( out32, X );
        S[ 1 ] = SKP_ADD32( in32, X );

        /* Add, convert back to int16 and store to output */
        out[ k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( out32, 11 ) );
    }
}
