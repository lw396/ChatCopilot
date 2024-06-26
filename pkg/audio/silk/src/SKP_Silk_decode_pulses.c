 

#include "SKP_Silk_main.h"

/*********************************************/
/* Decode quantization indices of excitation */
/*********************************************/
void SKP_Silk_decode_pulses(
    SKP_Silk_range_coder_state      *psRC,              /* I/O  Range coder state                           */
    SKP_Silk_decoder_control        *psDecCtrl,         /* I/O  Decoder control                             */
    SKP_int                         q[],                /* O    Excitation signal                           */
    const SKP_int                   frame_length        /* I    Frame length (preliminary)                  */
)
{
    SKP_int   i, j, k, iter, abs_q, nLS, bit;
    SKP_int   sum_pulses[ MAX_NB_SHELL_BLOCKS ], nLshifts[ MAX_NB_SHELL_BLOCKS ];
    SKP_int   *pulses_ptr;
    const SKP_uint16 *cdf_ptr;
    
    /*********************/
    /* Decode rate level */
    /*********************/
    SKP_Silk_range_decoder( &psDecCtrl->RateLevelIndex, psRC, 
            SKP_Silk_rate_levels_CDF[ psDecCtrl->sigtype ], SKP_Silk_rate_levels_CDF_offset );

    /* Calculate number of shell blocks */
    iter = frame_length / SHELL_CODEC_FRAME_LENGTH;
    
    /***************************************************/
    /* Sum-Weighted-Pulses Decoding                    */
    /***************************************************/
    cdf_ptr = SKP_Silk_pulses_per_block_CDF[ psDecCtrl->RateLevelIndex ];
    for( i = 0; i < iter; i++ ) {
        nLshifts[ i ] = 0;
        SKP_Silk_range_decoder( &sum_pulses[ i ], psRC, cdf_ptr, SKP_Silk_pulses_per_block_CDF_offset );

        /* LSB indication */
        while( sum_pulses[ i ] == ( MAX_PULSES + 1 ) ) {
            nLshifts[ i ]++;
            SKP_Silk_range_decoder( &sum_pulses[ i ], psRC, 
                    SKP_Silk_pulses_per_block_CDF[ N_RATE_LEVELS - 1 ], SKP_Silk_pulses_per_block_CDF_offset );
        }
    }
    
    /***************************************************/
    /* Shell decoding                                  */
    /***************************************************/
    for( i = 0; i < iter; i++ ) {
        if( sum_pulses[ i ] > 0 ) {
            SKP_Silk_shell_decoder( &q[ SKP_SMULBB( i, SHELL_CODEC_FRAME_LENGTH ) ], psRC, sum_pulses[ i ] );
        } else {
            SKP_memset( &q[ SKP_SMULBB( i, SHELL_CODEC_FRAME_LENGTH ) ], 0, SHELL_CODEC_FRAME_LENGTH * sizeof( SKP_int ) );
        }
    }

    /***************************************************/
    /* LSB Decoding                                    */
    /***************************************************/
    for( i = 0; i < iter; i++ ) {
        if( nLshifts[ i ] > 0 ) {
            nLS = nLshifts[ i ];
            pulses_ptr = &q[ SKP_SMULBB( i, SHELL_CODEC_FRAME_LENGTH ) ];
            for( k = 0; k < SHELL_CODEC_FRAME_LENGTH; k++ ) {
                abs_q = pulses_ptr[ k ];
                for( j = 0; j < nLS; j++ ) {
                    abs_q = SKP_LSHIFT( abs_q, 1 ); 
                    SKP_Silk_range_decoder( &bit, psRC, SKP_Silk_lsb_CDF, 1 );
                    abs_q += bit;
                }
                pulses_ptr[ k ] = abs_q;
            }
        }
    }

    /****************************************/
    /* Decode and add signs to pulse signal */
    /****************************************/
    SKP_Silk_decode_signs( psRC, q, frame_length, psDecCtrl->sigtype, 
        psDecCtrl->QuantOffsetType, psDecCtrl->RateLevelIndex);
}
