 

#include "SKP_Silk_main.h"

#define OFFSET          ( ( MIN_QGAIN_DB * 128 ) / 6 + 16 * 128 )
#define SCALE_Q16       ( ( 65536 * ( N_LEVELS_QGAIN - 1 ) ) / ( ( ( MAX_QGAIN_DB - MIN_QGAIN_DB ) * 128 ) / 6 ) )
#define INV_SCALE_Q16   ( ( 65536 * ( ( ( MAX_QGAIN_DB - MIN_QGAIN_DB ) * 128 ) / 6 ) ) / ( N_LEVELS_QGAIN - 1 ) )

/* Gain scalar quantization with hysteresis, uniform on log scale */
void SKP_Silk_gains_quant(
    SKP_int                         ind[ NB_SUBFR ],        /* O    gain indices                            */
    SKP_int32                       gain_Q16[ NB_SUBFR ],   /* I/O  gains (quantized out)                   */
    SKP_int                         *prev_ind,              /* I/O  last index in previous frame            */
    const SKP_int                   conditional             /* I    first gain is delta coded if 1          */
)
{
    SKP_int k;

    for( k = 0; k < NB_SUBFR; k++ ) {
        /* Add half of previous quantization error, convert to log scale, scale, floor() */
        ind[ k ] = SKP_SMULWB( SCALE_Q16, SKP_Silk_lin2log( gain_Q16[ k ] ) - OFFSET );

        /* Round towards previous quantized gain (hysteresis) */
        if( ind[ k ] < *prev_ind ) {
            ind[ k ]++;
        }

        /* Compute delta indices and limit */
        if( k == 0 && conditional == 0 ) {
            /* Full index */
            ind[ k ] = SKP_LIMIT_int( ind[ k ], 0, N_LEVELS_QGAIN - 1 );
            ind[ k ] = SKP_max_int( ind[ k ], *prev_ind + MIN_DELTA_GAIN_QUANT );
            *prev_ind = ind[ k ];
        } else {
            /* Delta index */
            ind[ k ] = SKP_LIMIT_int( ind[ k ] - *prev_ind, MIN_DELTA_GAIN_QUANT, MAX_DELTA_GAIN_QUANT );
            /* Accumulate deltas */
            *prev_ind += ind[ k ];
            /* Shift to make non-negative */
            ind[ k ] -= MIN_DELTA_GAIN_QUANT;
        }

        /* Convert to linear scale and scale */
        gain_Q16[ k ] = SKP_Silk_log2lin( SKP_min_32( SKP_SMULWB( INV_SCALE_Q16, *prev_ind ) + OFFSET, 3967 ) ); /* 3968 = 31 in Q7 */
    }
}

/* Gains scalar dequantization, uniform on log scale */
void SKP_Silk_gains_dequant(
    SKP_int32                       gain_Q16[ NB_SUBFR ],   /* O    quantized gains                         */
    const SKP_int                   ind[ NB_SUBFR ],        /* I    gain indices                            */
    SKP_int                         *prev_ind,              /* I/O  last index in previous frame            */
    const SKP_int                   conditional             /* I    first gain is delta coded if 1          */
)
{
    SKP_int   k;

    for( k = 0; k < NB_SUBFR; k++ ) {
        if( k == 0 && conditional == 0 ) {
            *prev_ind = ind[ k ];
        } else {
            /* Delta index */
            *prev_ind += ind[ k ] + MIN_DELTA_GAIN_QUANT;
        }

        /* Convert to linear scale and scale */
        gain_Q16[ k ] = SKP_Silk_log2lin( SKP_min_32( SKP_SMULWB( INV_SCALE_Q16, *prev_ind ) + OFFSET, 3967 ) ); /* 3968 = 31 in Q7 */
    }
}
