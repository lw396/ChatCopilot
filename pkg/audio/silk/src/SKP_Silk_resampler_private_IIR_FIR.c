 

/*																		*
 * File Name:	SKP_Silk_resampler_private_IIR_FIR.c  			    *
 *																		*
 * Description: Hybrid IIR/FIR polyphase implementation of resampling	*
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 * All rights reserved.													*
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"
#if EMBEDDED_ARM<5
SKP_INLINE SKP_int16 *SKP_Silk_resampler_private_IIR_FIR_INTERPOL( 
			SKP_int16 * out, SKP_int16 * buf, SKP_int32 max_index_Q16 , SKP_int32 index_increment_Q16 ){
	SKP_int32 index_Q16, res_Q15;
	SKP_int16 *buf_ptr;
	SKP_int32 table_index;
	/* Interpolate upsampled signal and store in output array */
	for( index_Q16 = 0; index_Q16 < max_index_Q16; index_Q16 += index_increment_Q16 ) {
        table_index = SKP_SMULWB( index_Q16 & 0xFFFF, 144 );
        buf_ptr = &buf[ index_Q16 >> 16 ];
            
        res_Q15 = SKP_SMULBB(          buf_ptr[ 0 ], SKP_Silk_resampler_frac_FIR_144[       table_index ][ 0 ] );
        res_Q15 = SKP_SMLABB( res_Q15, buf_ptr[ 1 ], SKP_Silk_resampler_frac_FIR_144[       table_index ][ 1 ] );
        res_Q15 = SKP_SMLABB( res_Q15, buf_ptr[ 2 ], SKP_Silk_resampler_frac_FIR_144[       table_index ][ 2 ] );
        res_Q15 = SKP_SMLABB( res_Q15, buf_ptr[ 3 ], SKP_Silk_resampler_frac_FIR_144[ 143 - table_index ][ 2 ] );
        res_Q15 = SKP_SMLABB( res_Q15, buf_ptr[ 4 ], SKP_Silk_resampler_frac_FIR_144[ 143 - table_index ][ 1 ] );
        res_Q15 = SKP_SMLABB( res_Q15, buf_ptr[ 5 ], SKP_Silk_resampler_frac_FIR_144[ 143 - table_index ][ 0 ] );          
		*out++ = (SKP_int16)SKP_SAT16( SKP_RSHIFT_ROUND( res_Q15, 15 ) );
	}
	return out;	
}
#else
extern SKP_int16 *SKP_Silk_resampler_private_IIR_FIR_INTERPOL( 
			SKP_int16 * out, SKP_int16 * buf, SKP_int32 max_index_Q16 , SKP_int32 index_increment_Q16 );
#endif
/* Upsample using a combination of allpass-based 2x upsampling and FIR interpolation */
void SKP_Silk_resampler_private_IIR_FIR(
	void	                        *SS,		    /* I/O: Resampler state 						*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
)
{
    SKP_Silk_resampler_state_struct *S = (SKP_Silk_resampler_state_struct *)SS;
	SKP_int32 nSamplesIn;
	SKP_int32 max_index_Q16, index_increment_Q16;
	SKP_int16 buf[ 2 * RESAMPLER_MAX_BATCH_SIZE_IN + RESAMPLER_ORDER_FIR_144 ];
    

	/* Copy buffered samples to start of buffer */	
	SKP_memcpy( buf, S->sFIR, RESAMPLER_ORDER_FIR_144 * sizeof( SKP_int32 ) );

	/* Iterate over blocks of frameSizeIn input samples */
    index_increment_Q16 = S->invRatio_Q16;
	while( 1 ) {
		nSamplesIn = SKP_min( inLen, S->batchSize );

        if( S->input2x == 1 ) {
		    /* Upsample 2x */
            S->up2_function( S->sIIR, &buf[ RESAMPLER_ORDER_FIR_144 ], in, nSamplesIn );
        } else {
		    /* Fourth-order ARMA filter */
            SKP_Silk_resampler_private_ARMA4( S->sIIR, &buf[ RESAMPLER_ORDER_FIR_144 ], in, S->Coefs, nSamplesIn );
        }

        max_index_Q16 = SKP_LSHIFT32( nSamplesIn, 16 + S->input2x );         /* +1 if 2x upsampling */
		out = SKP_Silk_resampler_private_IIR_FIR_INTERPOL(out, buf, max_index_Q16, index_increment_Q16);    
		in += nSamplesIn;
		inLen -= nSamplesIn;

		if( inLen > 0 ) {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			SKP_memcpy( buf, &buf[ nSamplesIn << S->input2x ], RESAMPLER_ORDER_FIR_144 * sizeof( SKP_int32 ) );
		} else {
			break;
		}
	}

	/* Copy last part of filtered signal to the state for the next call */
	SKP_memcpy( S->sFIR, &buf[nSamplesIn << S->input2x ], RESAMPLER_ORDER_FIR_144 * sizeof( SKP_int32 ) );
}

