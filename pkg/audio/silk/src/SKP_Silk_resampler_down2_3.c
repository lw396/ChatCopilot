 

/*                                                                      *
 * SKP_Silk_resampler_down2_3.c                                       *
 *                                                                      *
 * Downsample by a factor 2/3, low quality                              *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

#define ORDER_FIR                   4

/* Downsample by a factor 2/3, low quality */
void SKP_Silk_resampler_down2_3(
    SKP_int32                           *S,         /* I/O: State vector [ 6 ]                  */
    SKP_int16                           *out,       /* O:   Output signal [ floor(2*inLen/3) ]  */
    const SKP_int16                     *in,        /* I:   Input signal [ inLen ]              */
    SKP_int32                           inLen       /* I:   Number of input samples             */
)
{
	SKP_int32 nSamplesIn, counter, res_Q6;
	SKP_int32 buf[ RESAMPLER_MAX_BATCH_SIZE_IN + ORDER_FIR ];
	SKP_int32 *buf_ptr;

	/* Copy buffered samples to start of buffer */	
	SKP_memcpy( buf, S, ORDER_FIR * sizeof( SKP_int32 ) );

	/* Iterate over blocks of frameSizeIn input samples */
	while( 1 ) {
		nSamplesIn = SKP_min( inLen, RESAMPLER_MAX_BATCH_SIZE_IN );

	    /* Second-order AR filter (output in Q8) */
	    SKP_Silk_resampler_private_AR2( &S[ ORDER_FIR ], &buf[ ORDER_FIR ], in, 
            SKP_Silk_Resampler_2_3_COEFS_LQ, nSamplesIn );

		/* Interpolate filtered signal */
        buf_ptr = buf;
        counter = nSamplesIn;
        while( counter > 2 ) {
            /* Inner product */
		    res_Q6 = SKP_SMULWB(         buf_ptr[ 0 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 2 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 1 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 3 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 2 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 5 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 3 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 4 ] );

            /* Scale down, saturate and store in output array */
            *out++ = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( res_Q6, 6 ) );

		    res_Q6 = SKP_SMULWB(         buf_ptr[ 1 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 4 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 2 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 5 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 3 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 3 ] );
		    res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 4 ], SKP_Silk_Resampler_2_3_COEFS_LQ[ 2 ] );

            /* Scale down, saturate and store in output array */
            *out++ = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( res_Q6, 6 ) );

            buf_ptr += 3;
            counter -= 3;
        }

		in += nSamplesIn;
		inLen -= nSamplesIn;

		if( inLen > 0 ) {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			SKP_memcpy( buf, &buf[ nSamplesIn ], ORDER_FIR * sizeof( SKP_int32 ) );
		} else {
			break;
		}
	}

	/* Copy last part of filtered signal to the state for the next call */
	SKP_memcpy( S, &buf[ nSamplesIn ], ORDER_FIR * sizeof( SKP_int32 ) );
}
