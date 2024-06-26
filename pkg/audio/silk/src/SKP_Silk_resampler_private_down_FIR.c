 

/*																		*
 * File Name:	SKP_Silk_resampler_private_down_FIR.c                 *
 *																		*
 * Description: Hybrid IIR/FIR polyphase implementation of resampling	*
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 * All rights reserved.													*
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"
#if EMBEDDED_ARM<5
SKP_INLINE SKP_int16 *SKP_Silk_resampler_private_down_FIR_INTERPOL0(
	SKP_int16 *out, SKP_int32 *buf2, const SKP_int16 *FIR_Coefs, SKP_int32 max_index_Q16, SKP_int32 index_increment_Q16){
	
	SKP_int32 index_Q16, res_Q6;
	SKP_int32 *buf_ptr;
	for( index_Q16 = 0; index_Q16 < max_index_Q16; index_Q16 += index_increment_Q16 ) {
		/* Integer part gives pointer to buffered input */
		buf_ptr = buf2 + SKP_RSHIFT( index_Q16, 16 );

		/* Inner product */
		res_Q6 = SKP_SMULWB(         SKP_ADD32( buf_ptr[ 0 ], buf_ptr[ 11 ] ), FIR_Coefs[ 0 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, SKP_ADD32( buf_ptr[ 1 ], buf_ptr[ 10 ] ), FIR_Coefs[ 1 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, SKP_ADD32( buf_ptr[ 2 ], buf_ptr[  9 ] ), FIR_Coefs[ 2 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, SKP_ADD32( buf_ptr[ 3 ], buf_ptr[  8 ] ), FIR_Coefs[ 3 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, SKP_ADD32( buf_ptr[ 4 ], buf_ptr[  7 ] ), FIR_Coefs[ 4 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, SKP_ADD32( buf_ptr[ 5 ], buf_ptr[  6 ] ), FIR_Coefs[ 5 ] );

			    /* Scale down, saturate and store in output array */
		*out++ = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( res_Q6, 6 ) );
	}
	return out;
}

SKP_INLINE SKP_int16 *SKP_Silk_resampler_private_down_FIR_INTERPOL1(
	SKP_int16 *out, SKP_int32 *buf2, const SKP_int16 *FIR_Coefs, SKP_int32 max_index_Q16, SKP_int32 index_increment_Q16, SKP_int32 FIR_Fracs){
	
	SKP_int32 index_Q16, res_Q6;
	SKP_int32 *buf_ptr;
	SKP_int32 interpol_ind;
	const SKP_int16 *interpol_ptr;
	for( index_Q16 = 0; index_Q16 < max_index_Q16; index_Q16 += index_increment_Q16 ) {
		/* Integer part gives pointer to buffered input */
		buf_ptr = buf2 + SKP_RSHIFT( index_Q16, 16 );

		/* Fractional part gives interpolation coefficients */
		interpol_ind = SKP_SMULWB( index_Q16 & 0xFFFF, FIR_Fracs );

		/* Inner product */
		interpol_ptr = &FIR_Coefs[ RESAMPLER_DOWN_ORDER_FIR / 2 * interpol_ind ];
		res_Q6 = SKP_SMULWB(         buf_ptr[ 0 ], interpol_ptr[ 0 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 1 ], interpol_ptr[ 1 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 2 ], interpol_ptr[ 2 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 3 ], interpol_ptr[ 3 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 4 ], interpol_ptr[ 4 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 5 ], interpol_ptr[ 5 ] );
		interpol_ptr = &FIR_Coefs[ RESAMPLER_DOWN_ORDER_FIR / 2 * ( FIR_Fracs - 1 - interpol_ind ) ];
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 11 ], interpol_ptr[ 0 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[ 10 ], interpol_ptr[ 1 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[  9 ], interpol_ptr[ 2 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[  8 ], interpol_ptr[ 3 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[  7 ], interpol_ptr[ 4 ] );
		res_Q6 = SKP_SMLAWB( res_Q6, buf_ptr[  6 ], interpol_ptr[ 5 ] );

		/* Scale down, saturate and store in output array */
		*out++ = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( res_Q6, 6 ) );
	}
	return out;
}

#else
extern SKP_int16 *SKP_Silk_resampler_private_down_FIR_INTERPOL0(
	SKP_int16 *out, SKP_int32 *buf2, const SKP_int16 *FIR_Coefs, SKP_int32 max_index_Q16, SKP_int32 index_increment_Q16);
extern SKP_int16 *SKP_Silk_resampler_private_down_FIR_INTERPOL1(
	SKP_int16 *out, SKP_int32 *buf2, const SKP_int16 *FIR_Coefs, SKP_int32 max_index_Q16, SKP_int32 index_increment_Q16, SKP_int32 FIR_Fracs);	
#endif

/* Resample with a 2x downsampler (optional), a 2nd order AR filter followed by FIR interpolation */
void SKP_Silk_resampler_private_down_FIR(
	void	                        *SS,		    /* I/O: Resampler state 						*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
)
{
    SKP_Silk_resampler_state_struct *S = (SKP_Silk_resampler_state_struct *)SS;
	SKP_int32 nSamplesIn;
	SKP_int32 max_index_Q16, index_increment_Q16;
	SKP_int16 buf1[ RESAMPLER_MAX_BATCH_SIZE_IN / 2 ];
	SKP_int32 buf2[ RESAMPLER_MAX_BATCH_SIZE_IN + RESAMPLER_DOWN_ORDER_FIR ];
	const SKP_int16 *FIR_Coefs;

	/* Copy buffered samples to start of buffer */	
	SKP_memcpy( buf2, S->sFIR, RESAMPLER_DOWN_ORDER_FIR * sizeof( SKP_int32 ) );

    FIR_Coefs = &S->Coefs[ 2 ];

	/* Iterate over blocks of frameSizeIn input samples */
    index_increment_Q16 = S->invRatio_Q16;
	while( 1 ) {
		nSamplesIn = SKP_min( inLen, S->batchSize );

        if( S->input2x == 1 ) {
            /* Downsample 2x */
            SKP_Silk_resampler_down2( S->sDown2, buf1, in, nSamplesIn );

            nSamplesIn = SKP_RSHIFT32( nSamplesIn, 1 );

		    /* Second-order AR filter (output in Q8) */
		    SKP_Silk_resampler_private_AR2( S->sIIR, &buf2[ RESAMPLER_DOWN_ORDER_FIR ], buf1, S->Coefs, nSamplesIn );
        } else {
		    /* Second-order AR filter (output in Q8) */
		    SKP_Silk_resampler_private_AR2( S->sIIR, &buf2[ RESAMPLER_DOWN_ORDER_FIR ], in, S->Coefs, nSamplesIn );
        }

        max_index_Q16 = SKP_LSHIFT32( nSamplesIn, 16 );

		/* Interpolate filtered signal */
        if( S->FIR_Fracs == 1 ) {
    		out = SKP_Silk_resampler_private_down_FIR_INTERPOL0(out, buf2, FIR_Coefs, max_index_Q16, index_increment_Q16);
        } else {
    		out = SKP_Silk_resampler_private_down_FIR_INTERPOL1(out, buf2, FIR_Coefs, max_index_Q16, index_increment_Q16, S->FIR_Fracs);
        }
        
		in += nSamplesIn << S->input2x;
		inLen -= nSamplesIn << S->input2x;

		if( inLen > S->input2x ) {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			SKP_memcpy( buf2, &buf2[ nSamplesIn ], RESAMPLER_DOWN_ORDER_FIR * sizeof( SKP_int32 ) );
		} else {
			break;
		}
	}

	/* Copy last part of filtered signal to the state for the next call */
	SKP_memcpy( S->sFIR, &buf2[ nSamplesIn ], RESAMPLER_DOWN_ORDER_FIR * sizeof( SKP_int32 ) );
}

