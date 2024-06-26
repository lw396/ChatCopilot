 

/*                                                                      *
 * SKP_Silk_resampler_private_up2_HQ.c                                *
 *                                                                      *
 * Upsample by a factor 2, high quality                                 *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

/* Upsample by a factor 2, high quality */
/* Uses 2nd order allpass filters for the 2x upsampling, followed by a      */
/* notch filter just above Nyquist.                                         */
#if (EMBEDDED_ARM<5) 
void SKP_Silk_resampler_private_up2_HQ(
	SKP_int32	                    *S,			    /* I/O: Resampler state [ 6 ]					*/
    SKP_int16                       *out,           /* O:   Output signal [ 2 * len ]               */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       len             /* I:   Number of INPUT samples                 */
)
{
    SKP_int32 k;
    SKP_int32 in32, out32_1, out32_2, Y, X;

    SKP_assert( SKP_Silk_resampler_up2_hq_0[ 0 ] > 0 );
    SKP_assert( SKP_Silk_resampler_up2_hq_0[ 1 ] < 0 );
    SKP_assert( SKP_Silk_resampler_up2_hq_1[ 0 ] > 0 );
    SKP_assert( SKP_Silk_resampler_up2_hq_1[ 1 ] < 0 );
    
    /* Internal variables and state are in Q10 format */
    for( k = 0; k < len; k++ ) {
        /* Convert to Q10 */
        in32 = SKP_LSHIFT( (SKP_int32)in[ k ], 10 );

        /* First all-pass section for even output sample */
        Y       = SKP_SUB32( in32, S[ 0 ] );
        X       = SKP_SMULWB( Y, SKP_Silk_resampler_up2_hq_0[ 0 ] );
        out32_1 = SKP_ADD32( S[ 0 ], X );
        S[ 0 ]  = SKP_ADD32( in32, X );

        /* Second all-pass section for even output sample */
        Y       = SKP_SUB32( out32_1, S[ 1 ] );
        X       = SKP_SMLAWB( Y, Y, SKP_Silk_resampler_up2_hq_0[ 1 ] );
        out32_2 = SKP_ADD32( S[ 1 ], X );
        S[ 1 ]  = SKP_ADD32( out32_1, X );

        /* Biquad notch filter */
        out32_2 = SKP_SMLAWB( out32_2, S[ 5 ], SKP_Silk_resampler_up2_hq_notch[ 2 ] );
        out32_2 = SKP_SMLAWB( out32_2, S[ 4 ], SKP_Silk_resampler_up2_hq_notch[ 1 ] );
        out32_1 = SKP_SMLAWB( out32_2, S[ 4 ], SKP_Silk_resampler_up2_hq_notch[ 0 ] );
        S[ 5 ]  = SKP_SUB32(  out32_2, S[ 5 ] );
        
        /* Apply gain in Q15, convert back to int16 and store to output */
        out[ 2 * k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT32( 
            SKP_SMLAWB( 256, out32_1, SKP_Silk_resampler_up2_hq_notch[ 3 ] ), 9 ) );

        /* First all-pass section for odd output sample */
        Y       = SKP_SUB32( in32, S[ 2 ] );
        X       = SKP_SMULWB( Y, SKP_Silk_resampler_up2_hq_1[ 0 ] );
        out32_1 = SKP_ADD32( S[ 2 ], X );
        S[ 2 ]  = SKP_ADD32( in32, X );

        /* Second all-pass section for odd output sample */
        Y       = SKP_SUB32( out32_1, S[ 3 ] );
        X       = SKP_SMLAWB( Y, Y, SKP_Silk_resampler_up2_hq_1[ 1 ] );
        out32_2 = SKP_ADD32( S[ 3 ], X );
        S[ 3 ]  = SKP_ADD32( out32_1, X );

        /* Biquad notch filter */
        out32_2 = SKP_SMLAWB( out32_2, S[ 4 ], SKP_Silk_resampler_up2_hq_notch[ 2 ] );
        out32_2 = SKP_SMLAWB( out32_2, S[ 5 ], SKP_Silk_resampler_up2_hq_notch[ 1 ] );
        out32_1 = SKP_SMLAWB( out32_2, S[ 5 ], SKP_Silk_resampler_up2_hq_notch[ 0 ] );
        S[ 4 ]  = SKP_SUB32(  out32_2, S[ 4 ] );
        
        /* Apply gain in Q15, convert back to int16 and store to output */
        out[ 2 * k + 1 ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT32( 
            SKP_SMLAWB( 256, out32_1, SKP_Silk_resampler_up2_hq_notch[ 3 ] ), 9 ) );
    }
}
#endif


void SKP_Silk_resampler_private_up2_HQ_wrapper(
	void	                        *SS,		    /* I/O: Resampler state (unused)				*/
    SKP_int16                       *out,           /* O:   Output signal [ 2 * len ]               */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       len             /* I:   Number of input samples                 */
)
{
    SKP_Silk_resampler_state_struct *S = (SKP_Silk_resampler_state_struct *)SS;
    SKP_Silk_resampler_private_up2_HQ( S->sIIR, out, in, len );
}
